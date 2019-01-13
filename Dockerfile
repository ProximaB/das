# Step 1: build executable binary
FROM golang:alpine AS builder
# Install Git
RUN apk update && apk add git
# Copy source code to /src in Docker image
COPY . $GOPATH/src/github.com/DancesportSoftware/das/
WORKDIR $GOPATH/src/github.com/DancesportSoftware/das/
# Get all depdendencies
RUN go get -d ./...
# Build executable binary
RUN go build -o /go/bin/das
# Step 2: build a small image. scratch is an empty image
FROM alpine
WORKDIR /bin
# Copy the static executable
COPY --from=builder /go/bin/das /go/bin/das
# Set environment variable. TODO: it doesn't work as the app will look for PSQL within the image
ENV POSTGRES_CONNECTION user\=dasdev password\=dAs\!@#\$1234 dbname\=das sslmode\=disable
# Server File
EXPOSE 8080
ENTRYPOINT ["/go/bin/das"]