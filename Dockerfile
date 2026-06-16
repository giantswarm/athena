FROM gsoci.azurecr.io/giantswarm/alpine:3.24.1

RUN apk add --no-cache ca-certificates

ADD ./athena /athena

ENTRYPOINT ["/athena"]
