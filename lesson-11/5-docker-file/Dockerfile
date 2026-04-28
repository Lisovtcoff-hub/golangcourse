FROM golang:1.24.6-alpine3.22 AS build

WORKDIR /app

# Modules layer
COPY go.mod go.sum ./
RUN go mod download

# Build layer
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /myapp ./main.go

FROM alpine:3.22 AS run

COPY --from=build /myapp /myapp

EXPOSE 8080

CMD ["/myapp"]