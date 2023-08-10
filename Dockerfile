FROM quay.io/giantswarm/alpine:3.18.3

RUN apk add --no-cache ca-certificates

ADD ./athena /athena

ENTRYPOINT ["/athena"]
