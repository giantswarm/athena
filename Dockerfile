FROM gsoci.azurecr.io/giantswarm/alpine:3.23.4

RUN apk add --no-cache ca-certificates

ADD ./athena /athena

ENTRYPOINT ["/athena"]
