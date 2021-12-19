###################################################
# Plain env as basement and for local development #
###################################################
FROM golang:1.17-alpine as env

MAINTAINER ayubirazeh@gmail.com

# Add support for Delve debugger.
RUN apk add --no-cache ca-certificates git

##########################################################
# Prepare a build container with all dependencies inside #
##########################################################
FROM env as builder
WORKDIR $GOPATH/src/github.com/ayubirz/redishop
COPY ./ .
RUN go build -o /go/bin/main main.go

###########################################
# Create clean container with binary only #
###########################################
FROM alpine as exec

RUN apk add --update bash ca-certificates

WORKDIR /app
COPY --from=builder /go/bin/main ./

ENTRYPOINT ["/app/main"]