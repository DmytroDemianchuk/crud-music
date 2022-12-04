FROM golang:1.19.3-alpine as builder

RUN mkdir /build
COPY . /build
WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /bin/musics-app cmd/main.go

# generate clean, final image for end users
FROM alpine:3.11.3
COPY --from=builder /bin/musics-app /bin/musics-app

# executable
ENTRYPOINT [ "/bin/musics-app" ]