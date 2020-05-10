# Build container
FROM golang:latest AS api_builder
ADD . /api
WORKDIR /api
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o ./bin/api_server cmd/main.go

# Run container
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=api_builder /api/bin/ ./api
WORKDIR /api/
RUN chmod +x ./api_server
EXPOSE 8080
CMD ./api_server