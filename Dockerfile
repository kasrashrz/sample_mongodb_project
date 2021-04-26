FROM golang:1.15.2-alpine

RUN apk add --no-cache git

COPY  . /app

WORKDIR /app

# COPY . .

#RUN go mod download

RUN go build -o main main.go

EXPOSE 3000

CMD ["./main"]