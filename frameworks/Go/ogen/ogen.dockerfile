FROM docker.io/golang:1.20

WORKDIR /ogen

COPY ./src /ogen

RUN GOAMD64=v3 go build -ldflags="-s -w" -o ogen-server ./cmd/ogen-server

EXPOSE 8080

CMD ./ogen-server
