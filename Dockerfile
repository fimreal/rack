FROM golang:latest as builder

COPY . /srv/rack

RUN cd /srv/rack &&\
    make docker-build &&\
    ls -l bin 

# 下载证书
FROM alpine:latest as ca

RUN apk --no-cache add ca-certificates

# 
FROM scratch

LABEL source.url="https://github.com/fimreal/rack"

COPY --from=ca /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /srv/rack/bin/rack /rack

CMD [ "/rack" ]