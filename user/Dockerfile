FROM golang:1.12.5-alpine3.9

RUN apk update && \
    apk add git gcc g++
ENV PKG_PATH /go/src/github.com/uqichi/go-grpc-microservices
ENV SERVICE_PATH $PKG_PATH/user
ENV GO111MODULE on
WORKDIR $PKG_PATH
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . $PKG_PATH
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go install $SERVICE_PATH

FROM alpine:3.9

ENV TZ Asia/Tokyo
COPY --from=0 /go/bin/user /go/bin/user
ENTRYPOINT ["/go/bin/user"]
