FROM golang:alpine as builder

WORKDIR /build

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./webserver main.go

FROM scratch

COPY --from=builder /build/webserver /app/webserver

EXPOSE 8080
CMD ["/app/webserver"]
