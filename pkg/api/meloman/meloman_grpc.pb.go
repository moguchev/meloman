// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package meloman

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MelomanClient is the client API for Meloman service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MelomanClient interface {
	// проверка работоспособности сервиса
	//
	// access: *
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// авторизация
	//
	// access: *
	Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	// создание пользователя/регистрация
	//
	// access: *
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	// получение пользователя по id
	//
	// access: user, admin
	GetUserByID(ctx context.Context, in *GetUserByIDRequest, opts ...grpc.CallOption) (*GetUserByIDResponse, error)
	// получение всех пользователей
	//
	// access: user, admin
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
	// изменение роли пользователя
	//
	// access: admin
	UpdateUserRole(ctx context.Context, in *UpdateUserRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// создание исполнителя
	//
	// access: admin
	CreateArtist(ctx context.Context, in *CreateArtistRequest, opts ...grpc.CallOption) (*CreateArtistsResponse, error)
	// получить список всех исполнитей
	//
	// access: user, admin
	GetArtists(ctx context.Context, in *GetArtistsRequest, opts ...grpc.CallOption) (*GetArtistsResponse, error)
	// получение исполнителя по id
	//
	// access: user, admin
	GetArtistByID(ctx context.Context, in *GetArtistByIDRequest, opts ...grpc.CallOption) (*GetArtistByIDResponse, error)
	// получение списка альбомов исполнителя
	//
	// access: user, admin
	GetArtistAlbums(ctx context.Context, in *GetArtistAlbumsRequest, opts ...grpc.CallOption) (*GetArtistAlbumsResponse, error)
	// получить список всех форматов
	//
	// access: user, admin
	GetFormats(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetFormatsResponse, error)
	// получить список всех издателей
	//
	// access: user, admin
	GetLabels(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetLabelsResponse, error)
	// создание альбома
	//
	// access: admin
	CreateAlbum(ctx context.Context, in *CreateAlbumRequest, opts ...grpc.CallOption) (*CreateAlbumResponse, error)
	// создание/добавление композиции в альбом
	//
	// access: admin
	CreateTrack(ctx context.Context, in *CreateTrackRequest, opts ...grpc.CallOption) (*CreateTrackResponse, error)
	// получение композиций альбома
	//
	// access: user, admin
	GetAlbumTracks(ctx context.Context, in *GetAlbumTracksRequest, opts ...grpc.CallOption) (*GetAlbumTracksResponse, error)
	// поиск альбома по фильтру
	//
	// access: user, admin
	GetAlbumsByFilter(ctx context.Context, in *GetAlbumsByFilterRequest, opts ...grpc.CallOption) (*GetAlbumsByFilterResponse, error)
	// добавление альбома в коллекцию пользователя
	//
	// access: owner
	AddAlbum(ctx context.Context, in *AddAlbumRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// удаление альбома из коллекции пользователя
	//
	// access: owner
	RemoveAlbum(ctx context.Context, in *RemoveAlbumRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// получить коллекцию пользователя
	//
	// access: user, admin
	GetUserCollection(ctx context.Context, in *GetUserCollectionRequest, opts ...grpc.CallOption) (*GetUserCollectionResponse, error)
	// получить топ популярных исполнителей
	//
	// access: user, admin
	GetTopPopularArtists(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTopPopularArtistsResponse, error)
}

type melomanClient struct {
	cc grpc.ClientConnInterface
}

func NewMelomanClient(cc grpc.ClientConnInterface) MelomanClient {
	return &melomanClient{cc}
}

func (c *melomanClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/github.moguchev.meloman.Meloman/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *melomanClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/github.moguchev.meloman.Meloman/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *melomanClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/github.moguchev.meloman.Meloman/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *melomanClient) GetUserByID(ctx context.Context, in *GetUserByIDRequest, opts ...grpc.CallOption) (*GetUserByIDResponse, error) {
	out := new(GetUserByIDResponse)
	err := c.cc.Invoke(ctx, "/github.moguchev.meloman.Meloman/GetUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *melomanClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, "/github.moguchev.meloman.Meloman/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *melomanClient) UpdateUserRole(ctx context.Context, in *UpdateUserRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/github.moguchev.meloman.Meloman/UpdateUserRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *melomanClient) CreateArtist(ctx context.Context, in *CreateArtistRequest, opts ...grpc.CallOption) (*CreateArtistsResponse, error) {
	out := new(CreateArtistsResponse)
	err := c.cc.Invoke(ctx, "/github.moguchev.meloman.Meloman/CreateArtist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *melomanClient) GetArtists(ctx context.Context, in *GetArtistsRequest, opts ...grpc.CallOption) (*GetArtistsResponse, error) {
	out := new(GetArtistsResponse)
	err := c.cc.Invoke(ctx, "/github.moguchev.meloman.Meloman/GetArtists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *melomanClient) GetArtistByID(ctx context.Context, in *GetArtistByIDRequest, opts ...grpc.CallOption) (*GetArtistByIDResponse, error) {
	out := new(GetArtistByIDResponse)
	err := c.cc.Invoke(ctx, "/github.moguchev.meloman.Meloman/GetArtistByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *melomanClient) GetArtistAlbums(ctx context.Context, in *GetArtistAlbumsRequest, opts ...grpc.CallOption) (*GetArtistAlbumsResponse, error) {
	out := new(GetArtistAlbumsResponse)
	err := c.cc.Invoke(ctx, "/github.moguchev.meloman.Meloman/GetArtistAlbums", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *melomanClient) GetFormats(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetFormatsResponse, error) {
	out := new(GetFormatsResponse)
	err := c.cc.Invoke(ctx, "/github.moguchev.meloman.Meloman/GetFormats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *melomanClient) GetLabels(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetLabelsResponse, error) {
	out := new(GetLabelsResponse)
	err := c.cc.Invoke(ctx, "/github.moguchev.meloman.Meloman/GetLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *melomanClient) CreateAlbum(ctx context.Context, in *CreateAlbumRequest, opts ...grpc.CallOption) (*CreateAlbumResponse, error) {
	out := new(CreateAlbumResponse)
	err := c.cc.Invoke(ctx, "/github.moguchev.meloman.Meloman/CreateAlbum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *melomanClient) CreateTrack(ctx context.Context, in *CreateTrackRequest, opts ...grpc.CallOption) (*CreateTrackResponse, error) {
	out := new(CreateTrackResponse)
	err := c.cc.Invoke(ctx, "/github.moguchev.meloman.Meloman/CreateTrack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *melomanClient) GetAlbumTracks(ctx context.Context, in *GetAlbumTracksRequest, opts ...grpc.CallOption) (*GetAlbumTracksResponse, error) {
	out := new(GetAlbumTracksResponse)
	err := c.cc.Invoke(ctx, "/github.moguchev.meloman.Meloman/GetAlbumTracks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *melomanClient) GetAlbumsByFilter(ctx context.Context, in *GetAlbumsByFilterRequest, opts ...grpc.CallOption) (*GetAlbumsByFilterResponse, error) {
	out := new(GetAlbumsByFilterResponse)
	err := c.cc.Invoke(ctx, "/github.moguchev.meloman.Meloman/GetAlbumsByFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *melomanClient) AddAlbum(ctx context.Context, in *AddAlbumRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/github.moguchev.meloman.Meloman/AddAlbum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *melomanClient) RemoveAlbum(ctx context.Context, in *RemoveAlbumRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/github.moguchev.meloman.Meloman/RemoveAlbum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *melomanClient) GetUserCollection(ctx context.Context, in *GetUserCollectionRequest, opts ...grpc.CallOption) (*GetUserCollectionResponse, error) {
	out := new(GetUserCollectionResponse)
	err := c.cc.Invoke(ctx, "/github.moguchev.meloman.Meloman/GetUserCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *melomanClient) GetTopPopularArtists(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTopPopularArtistsResponse, error) {
	out := new(GetTopPopularArtistsResponse)
	err := c.cc.Invoke(ctx, "/github.moguchev.meloman.Meloman/GetTopPopularArtists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MelomanServer is the server API for Meloman service.
// All implementations must embed UnimplementedMelomanServer
// for forward compatibility
type MelomanServer interface {
	// проверка работоспособности сервиса
	//
	// access: *
	Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// авторизация
	//
	// access: *
	Auth(context.Context, *AuthRequest) (*AuthResponse, error)
	// создание пользователя/регистрация
	//
	// access: *
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	// получение пользователя по id
	//
	// access: user, admin
	GetUserByID(context.Context, *GetUserByIDRequest) (*GetUserByIDResponse, error)
	// получение всех пользователей
	//
	// access: user, admin
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	// изменение роли пользователя
	//
	// access: admin
	UpdateUserRole(context.Context, *UpdateUserRoleRequest) (*emptypb.Empty, error)
	// создание исполнителя
	//
	// access: admin
	CreateArtist(context.Context, *CreateArtistRequest) (*CreateArtistsResponse, error)
	// получить список всех исполнитей
	//
	// access: user, admin
	GetArtists(context.Context, *GetArtistsRequest) (*GetArtistsResponse, error)
	// получение исполнителя по id
	//
	// access: user, admin
	GetArtistByID(context.Context, *GetArtistByIDRequest) (*GetArtistByIDResponse, error)
	// получение списка альбомов исполнителя
	//
	// access: user, admin
	GetArtistAlbums(context.Context, *GetArtistAlbumsRequest) (*GetArtistAlbumsResponse, error)
	// получить список всех форматов
	//
	// access: user, admin
	GetFormats(context.Context, *emptypb.Empty) (*GetFormatsResponse, error)
	// получить список всех издателей
	//
	// access: user, admin
	GetLabels(context.Context, *emptypb.Empty) (*GetLabelsResponse, error)
	// создание альбома
	//
	// access: admin
	CreateAlbum(context.Context, *CreateAlbumRequest) (*CreateAlbumResponse, error)
	// создание/добавление композиции в альбом
	//
	// access: admin
	CreateTrack(context.Context, *CreateTrackRequest) (*CreateTrackResponse, error)
	// получение композиций альбома
	//
	// access: user, admin
	GetAlbumTracks(context.Context, *GetAlbumTracksRequest) (*GetAlbumTracksResponse, error)
	// поиск альбома по фильтру
	//
	// access: user, admin
	GetAlbumsByFilter(context.Context, *GetAlbumsByFilterRequest) (*GetAlbumsByFilterResponse, error)
	// добавление альбома в коллекцию пользователя
	//
	// access: owner
	AddAlbum(context.Context, *AddAlbumRequest) (*emptypb.Empty, error)
	// удаление альбома из коллекции пользователя
	//
	// access: owner
	RemoveAlbum(context.Context, *RemoveAlbumRequest) (*emptypb.Empty, error)
	// получить коллекцию пользователя
	//
	// access: user, admin
	GetUserCollection(context.Context, *GetUserCollectionRequest) (*GetUserCollectionResponse, error)
	// получить топ популярных исполнителей
	//
	// access: user, admin
	GetTopPopularArtists(context.Context, *emptypb.Empty) (*GetTopPopularArtistsResponse, error)
	mustEmbedUnimplementedMelomanServer()
}

// UnimplementedMelomanServer must be embedded to have forward compatible implementations.
type UnimplementedMelomanServer struct {
}

func (UnimplementedMelomanServer) Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedMelomanServer) Auth(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedMelomanServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedMelomanServer) GetUserByID(context.Context, *GetUserByIDRequest) (*GetUserByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByID not implemented")
}
func (UnimplementedMelomanServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedMelomanServer) UpdateUserRole(context.Context, *UpdateUserRoleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserRole not implemented")
}
func (UnimplementedMelomanServer) CreateArtist(context.Context, *CreateArtistRequest) (*CreateArtistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArtist not implemented")
}
func (UnimplementedMelomanServer) GetArtists(context.Context, *GetArtistsRequest) (*GetArtistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArtists not implemented")
}
func (UnimplementedMelomanServer) GetArtistByID(context.Context, *GetArtistByIDRequest) (*GetArtistByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArtistByID not implemented")
}
func (UnimplementedMelomanServer) GetArtistAlbums(context.Context, *GetArtistAlbumsRequest) (*GetArtistAlbumsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArtistAlbums not implemented")
}
func (UnimplementedMelomanServer) GetFormats(context.Context, *emptypb.Empty) (*GetFormatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFormats not implemented")
}
func (UnimplementedMelomanServer) GetLabels(context.Context, *emptypb.Empty) (*GetLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabels not implemented")
}
func (UnimplementedMelomanServer) CreateAlbum(context.Context, *CreateAlbumRequest) (*CreateAlbumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAlbum not implemented")
}
func (UnimplementedMelomanServer) CreateTrack(context.Context, *CreateTrackRequest) (*CreateTrackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrack not implemented")
}
func (UnimplementedMelomanServer) GetAlbumTracks(context.Context, *GetAlbumTracksRequest) (*GetAlbumTracksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlbumTracks not implemented")
}
func (UnimplementedMelomanServer) GetAlbumsByFilter(context.Context, *GetAlbumsByFilterRequest) (*GetAlbumsByFilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlbumsByFilter not implemented")
}
func (UnimplementedMelomanServer) AddAlbum(context.Context, *AddAlbumRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAlbum not implemented")
}
func (UnimplementedMelomanServer) RemoveAlbum(context.Context, *RemoveAlbumRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAlbum not implemented")
}
func (UnimplementedMelomanServer) GetUserCollection(context.Context, *GetUserCollectionRequest) (*GetUserCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCollection not implemented")
}
func (UnimplementedMelomanServer) GetTopPopularArtists(context.Context, *emptypb.Empty) (*GetTopPopularArtistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopPopularArtists not implemented")
}
func (UnimplementedMelomanServer) mustEmbedUnimplementedMelomanServer() {}

// UnsafeMelomanServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MelomanServer will
// result in compilation errors.
type UnsafeMelomanServer interface {
	mustEmbedUnimplementedMelomanServer()
}

func RegisterMelomanServer(s grpc.ServiceRegistrar, srv MelomanServer) {
	s.RegisterService(&Meloman_ServiceDesc, srv)
}

func _Meloman_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MelomanServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.meloman.Meloman/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MelomanServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meloman_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MelomanServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.meloman.Meloman/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MelomanServer).Auth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meloman_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MelomanServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.meloman.Meloman/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MelomanServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meloman_GetUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MelomanServer).GetUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.meloman.Meloman/GetUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MelomanServer).GetUserByID(ctx, req.(*GetUserByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meloman_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MelomanServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.meloman.Meloman/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MelomanServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meloman_UpdateUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MelomanServer).UpdateUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.meloman.Meloman/UpdateUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MelomanServer).UpdateUserRole(ctx, req.(*UpdateUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meloman_CreateArtist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArtistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MelomanServer).CreateArtist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.meloman.Meloman/CreateArtist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MelomanServer).CreateArtist(ctx, req.(*CreateArtistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meloman_GetArtists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArtistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MelomanServer).GetArtists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.meloman.Meloman/GetArtists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MelomanServer).GetArtists(ctx, req.(*GetArtistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meloman_GetArtistByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArtistByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MelomanServer).GetArtistByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.meloman.Meloman/GetArtistByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MelomanServer).GetArtistByID(ctx, req.(*GetArtistByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meloman_GetArtistAlbums_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArtistAlbumsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MelomanServer).GetArtistAlbums(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.meloman.Meloman/GetArtistAlbums",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MelomanServer).GetArtistAlbums(ctx, req.(*GetArtistAlbumsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meloman_GetFormats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MelomanServer).GetFormats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.meloman.Meloman/GetFormats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MelomanServer).GetFormats(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meloman_GetLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MelomanServer).GetLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.meloman.Meloman/GetLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MelomanServer).GetLabels(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meloman_CreateAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAlbumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MelomanServer).CreateAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.meloman.Meloman/CreateAlbum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MelomanServer).CreateAlbum(ctx, req.(*CreateAlbumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meloman_CreateTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTrackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MelomanServer).CreateTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.meloman.Meloman/CreateTrack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MelomanServer).CreateTrack(ctx, req.(*CreateTrackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meloman_GetAlbumTracks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlbumTracksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MelomanServer).GetAlbumTracks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.meloman.Meloman/GetAlbumTracks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MelomanServer).GetAlbumTracks(ctx, req.(*GetAlbumTracksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meloman_GetAlbumsByFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlbumsByFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MelomanServer).GetAlbumsByFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.meloman.Meloman/GetAlbumsByFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MelomanServer).GetAlbumsByFilter(ctx, req.(*GetAlbumsByFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meloman_AddAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAlbumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MelomanServer).AddAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.meloman.Meloman/AddAlbum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MelomanServer).AddAlbum(ctx, req.(*AddAlbumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meloman_RemoveAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveAlbumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MelomanServer).RemoveAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.meloman.Meloman/RemoveAlbum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MelomanServer).RemoveAlbum(ctx, req.(*RemoveAlbumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meloman_GetUserCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MelomanServer).GetUserCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.meloman.Meloman/GetUserCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MelomanServer).GetUserCollection(ctx, req.(*GetUserCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meloman_GetTopPopularArtists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MelomanServer).GetTopPopularArtists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.meloman.Meloman/GetTopPopularArtists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MelomanServer).GetTopPopularArtists(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Meloman_ServiceDesc is the grpc.ServiceDesc for Meloman service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Meloman_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.moguchev.meloman.Meloman",
	HandlerType: (*MelomanServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Meloman_Ping_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _Meloman_Auth_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Meloman_CreateUser_Handler,
		},
		{
			MethodName: "GetUserByID",
			Handler:    _Meloman_GetUserByID_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _Meloman_GetUsers_Handler,
		},
		{
			MethodName: "UpdateUserRole",
			Handler:    _Meloman_UpdateUserRole_Handler,
		},
		{
			MethodName: "CreateArtist",
			Handler:    _Meloman_CreateArtist_Handler,
		},
		{
			MethodName: "GetArtists",
			Handler:    _Meloman_GetArtists_Handler,
		},
		{
			MethodName: "GetArtistByID",
			Handler:    _Meloman_GetArtistByID_Handler,
		},
		{
			MethodName: "GetArtistAlbums",
			Handler:    _Meloman_GetArtistAlbums_Handler,
		},
		{
			MethodName: "GetFormats",
			Handler:    _Meloman_GetFormats_Handler,
		},
		{
			MethodName: "GetLabels",
			Handler:    _Meloman_GetLabels_Handler,
		},
		{
			MethodName: "CreateAlbum",
			Handler:    _Meloman_CreateAlbum_Handler,
		},
		{
			MethodName: "CreateTrack",
			Handler:    _Meloman_CreateTrack_Handler,
		},
		{
			MethodName: "GetAlbumTracks",
			Handler:    _Meloman_GetAlbumTracks_Handler,
		},
		{
			MethodName: "GetAlbumsByFilter",
			Handler:    _Meloman_GetAlbumsByFilter_Handler,
		},
		{
			MethodName: "AddAlbum",
			Handler:    _Meloman_AddAlbum_Handler,
		},
		{
			MethodName: "RemoveAlbum",
			Handler:    _Meloman_RemoveAlbum_Handler,
		},
		{
			MethodName: "GetUserCollection",
			Handler:    _Meloman_GetUserCollection_Handler,
		},
		{
			MethodName: "GetTopPopularArtists",
			Handler:    _Meloman_GetTopPopularArtists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "meloman/meloman.proto",
}
