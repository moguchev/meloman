package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/moguchev/meloman/db"
	"github.com/moguchev/meloman/internal/access"
	"github.com/moguchev/meloman/internal/auth"
	"github.com/moguchev/meloman/internal/service"
	"github.com/moguchev/meloman/pkg/api/meloman"
	gwmeloman "github.com/moguchev/meloman/pkg/gw/meloman"
)

const (
	ServerAdressGRPC = ":8090"
	ServerAdressHTTP = ":8080"
	SwaggerDir       = "./swaggerui"
	SecretKey        = "aQd23nsoEd"
	TokenDuration    = 30 * time.Minute
)

func serveSwagger(mux *http.ServeMux) {
	fileServer := http.FileServer(http.Dir(SwaggerDir))
	prefix := "/swaggerui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

func main() {
	url := os.Getenv("DATABASE_URL")

	// Create a server. Recovery handlers should typically be last in the chain so that other middleware
	// (e.g. logging) can operate on the recovered state instead of being directly affected by any panic
	_ = grpc.NewServer()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("create logger")
	}
	defer logger.Sync()

	// DB
	database, err := db.Initialize(ctx, url, logger)
	if err != nil {
		logger.Fatal("init database", zap.Error(err))
	}
	defer database.Close()

	if err = db.Migrate(url); err != nil {
		logger.Fatal("migrate database", zap.Error(err))
	}

	// auth
	jwtManager := auth.NewJWTManager(SecretKey, TokenDuration)
	authManager := auth.NewManager(jwtManager, access.AccessibleRoles(), logger)

	// Define customfunc to handle panic
	customFunc := func(p interface{}) (err error) {
		logger.Error("panic", zap.Any("panic", p))
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}
	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []recovery.Option{
		recovery.WithRecoveryHandler(customFunc),
	}

	// Create a gRPC server object
	grpcs := grpc.NewServer(
		grpc.ConnectionTimeout(5*time.Second),
		grpc_middleware.WithUnaryServerChain(
			recovery.UnaryServerInterceptor(opts...),
		),
		grpc_middleware.WithStreamServerChain(
			recovery.StreamServerInterceptor(opts...),
		),
		grpc.UnaryInterceptor(authManager.Unary()),
		grpc.StreamInterceptor(authManager.Stream()),
	)
	// Create Service
	srv := service.NewService(logger, database, authManager)

	// Attach the Meloman service to the server
	meloman.RegisterMelomanServer(grpcs, srv)

	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ServerAdressGRPC)
	if err != nil {
		logger.Fatal("Failed to listen", zap.Error(err))
	}

	var group errgroup.Group

	logger.Sugar().Infof("Server gRPC started on %s", ServerAdressGRPC)
	// Serve gRPC server
	group.Go(func() error {
		return grpcs.Serve(lis)
	})

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(ctx, ServerAdressGRPC,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		logger.Fatal("Failed to dial server", zap.Error(err))
	}

	mux := http.NewServeMux()

	// Create a gRPC Gateway mux
	gwmux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}))

	mux.Handle("/", gwmux)
	serveSwagger(mux)

	// Register Meloman
	if err = gwmeloman.RegisterMelomanHandler(ctx, gwmux, conn); err != nil {
		logger.Fatal("Failed to register gateway", zap.Error(err))
	}

	// Create a gRPC Gateway server
	gwServer := &http.Server{
		Addr:    ServerAdressHTTP,
		Handler: mux,
	}

	logger.Sugar().Infof("Server gRPC-Gateway started on %s", ServerAdressHTTP)
	group.Go(func() error {
		return gwServer.ListenAndServe()
	})

	logger.Fatal("serve", zap.Error(group.Wait()))
}
