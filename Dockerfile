FROM golang:latest as builder
COPY . /srv/rack
# ENV GOPROXY="https://goproxy.cn,direct"
RUN cd /srv/rack &&\
    make docker-build &&\
    ls -l bin 

# 下载证书
FROM alpine:latest as ca
# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk --no-cache add ca-certificates

# 
FROM scratch
LABEL source.url="https://github.com/fimreal/rack"

COPY --from=ca /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /srv/rack/bin/rack /rack

ENTRYPOINT [ "/rack" ]