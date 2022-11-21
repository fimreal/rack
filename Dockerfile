FROM golang:latest as builder

COPY . /srv/rack

RUN cd /srv/rack &&\
    make docker-build &&\
    ls -l bin


FROM scratch

LABEL source.url="https://github.com/fimreal/rack"

COPY --from=builder /srv/rack/bin/rack /rack

CMD [ "/rack" ]