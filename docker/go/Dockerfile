FROM golang

ENV GO111MODULE=on
ENV ROOT=/go/src/app
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR ${ROOT}
COPY ./src/api /go/src/app

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]