FROM quay.io/giantswarm/alpine:3.20.1

RUN apk add --no-cache ca-certificates

ADD ./athena /athena

ENTRYPOINT ["/athena"]
