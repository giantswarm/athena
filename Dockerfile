FROM quay.io/giantswarm/alpine:3.16.0

RUN apk add --no-cache ca-certificates

ADD ./athena /athena

ENTRYPOINT ["/athena"]
