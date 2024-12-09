FROM golang:1.23.3-alpine as builder

WORKDIR /usr/local/src

RUN apk --no-cache add  bash  git  make  gcc  gettext musl-dev

#dependencies
COPY app/go.mod app/go.sum ./
RUN go mod download

#build
COPY app ./
RUN go build -o ./bin/app cmd/main.go

FROM alpine

COPY --from=builder /usr/local/src/bin/app /
COPY --from=builder /usr/local/src/.env /.env

CMD ["/app"]