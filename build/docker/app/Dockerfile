FROM golang:1.17.3-alpine3.14 as builder
RUN apk update && apk add make

ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./

CMD ["go", "run", "./cmd/main.go"]
#RUN make

#FROM gcr.io/distroless/base:1e4a8bb3ad03f71b572cbcb3bbc25f3fd8d0ff14
#EXPOSE 8080
#COPY --from=builder /usr/src/app/bin/app /usr/local/bin
#CMD ["/usr/local/bin"]
