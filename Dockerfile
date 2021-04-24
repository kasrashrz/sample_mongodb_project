FROM golang:1.15.2-alpine

RUN apk add --no-cache git

COPY  . /app/go-sample-app
# Set the Current Working Directory inside the container
WORKDIR /app/go-sample-app

# We want to populate the module cache based on the go.{mod,sum} files.
RUN go mod download

# Build the Go app
RUN go build -o ./out/go-sample-app .
# This container exposes port 8080 to the outside world
EXPOSE 8000

RUN RUN echo $(ls -1)

# Run the binary program produced by `go install`
CMD ["./out/go-sample-app"]



