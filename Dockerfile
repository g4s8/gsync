FROM golang:1.9-alpine as build
WORKDIR /build
COPY main.go main.go
RUN go build -o gsync-server .

FROM alpine:latest
WORKDIR /app
COPY --from=build /build/gsync-server gsync-server
COPY scripts /scripts
RUN chmod +x /scripts/gsync.sh
RUN apk --update --no-cache add git
EXPOSE 8042
CMD ["/app/gsync-server"]
