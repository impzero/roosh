FROM golang:1.21.5-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN apk --no-cache add ca-certificates
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -a -o roosh-server ./cmd/server

FROM alpine
WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/roosh-server /app

EXPOSE 8080

CMD [ "./roosh-server", "web" ]
