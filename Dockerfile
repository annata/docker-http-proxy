FROM golang:1.13-alpine as builder
WORKDIR /http_proxy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM scratch
COPY --from=builder /http_proxy/http_proxy /
CMD ["/http_proxy"]