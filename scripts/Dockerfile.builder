FROM golang:1.20.2-alpine3.16
RUN apk -U add bash git gcc musl-dev make docker-cli curl ca-certificates
WORKDIR /go/src/github.com/terraform-providers/terraform-provider-rancher2