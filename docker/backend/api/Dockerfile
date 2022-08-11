FROM golang:latest as builder

ENV REPOSITORY github.com/norun9/postmantest
ENV GOOS=linux
ENV GOARCH=amd64
ENV WORKDIR $GOPATH/src/$REPOSITORY
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOFLAGS "-mod=vendor"

ADD . $GOPATH/src/$REPOSITORY/
WORKDIR $WORKDIR

RUN cd $GOPATH/src/$REPOSITORY/backend/cmd/api \
    && go mod vendor \
    && make build

FROM alpine:3.9
WORKDIR /app
RUN apk add --no-cache --update ca-certificates tzdata bash mysql-client
COPY --from=builder /go/src/github.com/norun9/postmantest/backend/cmd/api/api /usr/local/bin/
# https://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
ADD backend/config.yaml /app/
ADD backend/testdata/fixture/common/posts.yml /app/

ENV PWD $WORKDIR

EXPOSE 3306

RUN ls

CMD api
