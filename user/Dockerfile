ARG GOBIN=/go/bin

FROM golang:1.12.5-alpine3.9

WORKDIR /go/src/app
ENV GO111MODULE on
RUN apk update && \
    apk add git make gcc g++
COPY . /go/src/app
RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o $GOBIN/user

FROM alpine:3.9
ENV TZ Asia/Tokyo
ENV ACTIVITY_SERVICE_ADDR activity:50051
ENV PROJECT_SERVICE_ADDR project:50051
COPY --from=0 $GOBIN/user /go/bin/user
ENTRYPOINT ["/go/bin/user"]
