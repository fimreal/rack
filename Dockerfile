FROM golang:latest as builder

COPY . /srv/rack

RUN apk --no-cache add ca-certificates
RUN cd /srv/rack &&\
    make docker-build &&\
    ls -l bin 



FROM scratch

LABEL source.url="https://github.com/fimreal/rack"

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /srv/rack/bin/rack /rack

CMD [ "/rack" ]