FROM golang:alpine as build_base

WORKDIR /app

# Copy go mod and sum files
COPY go.mod .
COPY go.sum .
RUN go mod download

# This image builds the weavaite server
FROM build_base AS server_builder
# Here we copy the rest of the source code
COPY . .

RUN go build -o main ./main.go
RUN apk add --no-cache chromium

EXPOSE 80

CMD ["./main", "grpc"]