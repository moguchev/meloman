# Multi-stage builds
################################
# STEP 1 build executable binary
################################
FROM ubuntu:latest AS build

RUN apt-get update
RUN apt-get install -y wget git gcc g++ make bash

# Install golang
RUN wget -P /tmp https://dl.google.com/go/go1.15.1.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf /tmp/go1.15.1.linux-amd64.tar.gz
RUN rm /tmp/go1.15.1.linux-amd64.tar.gz
# Set golang env
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
# init work dir
RUN mkdir -p ${GOPATH}/src/github.com/moguchev/meloman
COPY . ${GOPATH}/src/github.com/moguchev/meloman
WORKDIR ${GOPATH}/src/github.com/moguchev/meloman
# build
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/meloman -v ./cmd/meloman

################################
# STEP 2 build a small image
############################
FROM scratch AS final

COPY --from=build /bin/meloman .

ADD ./db/migrations /db/migrations
ADD ./swaggerui ./swaggerui

ENV PG_HOST pg
ENV BOUNCER_HOST bouncer
 
EXPOSE 8080 8080
EXPOSE 8090 8090
# Run the executable
CMD ["./meloman"]