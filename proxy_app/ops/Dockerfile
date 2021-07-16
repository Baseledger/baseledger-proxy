FROM golang:1.16-alpine

WORKDIR /apps

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /proxy_app httpd/main.go

CMD ["/proxy_app"]