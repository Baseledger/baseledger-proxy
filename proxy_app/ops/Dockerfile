FROM golang:1.16-alpine

WORKDIR /apps

ARG DB_HOST
ARG DB_UB_USER
ARG DB_UB_PWD

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

EXPOSE 8080

ENV DB_HOST=$DB_HOST
ENV DB_UB_USER=$DB_UB_USER
ENV DB_UB_PWD=$DB_UB_PWD
ENV DB_UB_NAME=ub
ENV DB_SSLMODE=disable
ENV DB_BASELEDGER_NAME=baseledger
ENV DB_BASELEDGER_USER=baseledger
ENV NATS_URL=nats:4222

RUN go build -o /proxy_app httpd/main.go

CMD ["/proxy_app"]