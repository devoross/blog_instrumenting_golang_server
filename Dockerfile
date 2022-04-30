FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

ADD api ./api
ADD server ./server
COPY *.go ./

RUN go build -o instrumented_http_server

CMD [ "/instrumented_http_server" ]