FROM golang:1.16-alpine
RUN apk add build-base

WORKDIR /apps

ARG DB_HOST

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

EXPOSE 8080

ENV DB_HOST=$DB_HOST
ENV DB_UB_NAME=ub
ENV DB_SSLMODE=disable
ENV DB_BASELEDGER_NAME=baseledger
ENV DB_BASELEDGER_USER=baseledger
ENV NATS_URL=nats:4222
ENV BLOCKCHAIN_APP_URL=blockchain_app:1317
ENV TENDERMINT_API_URL=blockchain_app:26657
ENV API_CONCIRCLE_USER=<user>

RUN go build -o /proxy_app httpd/main.go

CMD ["/proxy_app"]