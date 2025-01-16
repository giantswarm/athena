FROM quay.io/giantswarm/alpine:3.21.2

RUN apk add --no-cache ca-certificates

ADD ./athena /athena

ENTRYPOINT ["/athena"]
