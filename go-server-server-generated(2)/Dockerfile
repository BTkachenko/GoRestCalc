FROM golang:1.18
RUN apt update -y && apt install -y collectd collectd-utils
WORKDIR /app
COPY ./go ./go
COPY main.go ./
COPY ./go.mod ./
COPY ./go.sum ./
COPY ./*.go ./


ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go mod tidy | go build -o ./gocalc


EXPOSE 8080
ENTRYPOINT ["./gocalc"]
