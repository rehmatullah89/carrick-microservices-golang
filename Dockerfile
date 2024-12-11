FROM golang:1.16-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
#RUN CGO_ENABLED=0 go test -v

RUN go build -o ./out/carrick-js-api .

# Start fresh from a smaller image
FROM alpine:latest
RUN apk add ca-certificates

WORKDIR /root/

COPY --from=build_base /tmp/app/out/carrick-js-api .

# This container exposes port 8000 to the outside world
EXPOSE 5000

# Run the binary program produced by `go build`
CMD ["./carrick-js-api"]