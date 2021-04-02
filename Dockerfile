FROM quay.io/giantswarm/alpine:3.12

USER root

RUN apk add --no-cache ca-certificates

RUN mkdir -p /opt
ADD ./athena /opt/athena

USER giantswarm

EXPOSE 8000
ENTRYPOINT ["/opt/athena"]
