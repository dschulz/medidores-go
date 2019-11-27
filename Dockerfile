FROM golang:latest as constructor

RUN mkdir -pv /go/src/medidores
WORKDIR /go/src/medidores

COPY . .
ENV GO111MODULE=on
RUN go mod tidy
RUN go build -a -o ./medidores ./cmd/app && go build  -a -o ./migrate ./cmd/migrate


# Etapa 2
FROM debian:latest

RUN apt-get update && apt-get install -y locales && rm -rf /var/lib/apt/lists/* \
    && localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8
ENV LANG en_US.utf8

ARG BUILD_TSTAMP
ARG VERSION_TAG

LABEL org.label-schema.name="medidores"
LABEL org.label-schema.version="$VERSION_TAG"
LABEL org.label-schema.description="API REST para catastro de medidores de agua y lecturas de consumo"
LABEL org.label-schema.vendor="8Bit S.A."
LABEL org.label-schema.build-date="$BUILD_TSTAMP"
LABEL org.label-schema.docker.cmd="docker run --name medidores01 --detach --publish 3000:3000 --mount type=bind,source=/ruta/absouta/config.json,target=/etc/config.json medidores:latest"
LABEL maintainer="admin@8bit.com.py"
LABEL org.label-schema.vcs-url="https://git.8bit.com.py/8Bit/medidores"

RUN mkdir -pv /aplicacion/bin

WORKDIR /aplicacion

COPY --from=constructor /go/src/medidores/medidores ./bin
COPY --from=constructor /go/src/medidores/migrate ./bin
COPY --from=constructor /go/src/medidores/docker/app/entrypoint.sh ./bin
COPY --from=constructor /go/src/medidores/migrations ./migrations


# Descomentar esta linea para correr en modo "production":
# ENV GO_ENV=prod

# Descomentar esta linea para correr en modo "development":
ENV GO_ENV=dev

# Escuchar en 0.0.0.0 para que sea accesible desde fuera del container
ENV LISTEN_ADDRESS=0.0.0.0

EXPOSE 3000

# Iniciar con exec para hacer fork/exit
#CMD exec /usr/local/bin/medidores

ENTRYPOINT ["/aplicacion/bin/entrypoint.sh"]