#!/usr/bin/env sh

IMAGEN="medidores"
VERSION_TAG="latest"
REMOTE="dschulzg/${IMAGEN}:${VERSION_TAG}"

# Construir imagen pasando la fecha/hora como build arg
docker build --no-cache=true --build-arg BUILD_TSTAMP=$(date -u +'%Y-%m-%dT%H:%M:%SZ')  --build-arg VERSION_TAG=$VERSION_TAG -t ${IMAGEN}:${VERSION_TAG} .

docker tag ${IMAGEN}:${VERSION_TAG} ${REMOTE}


# Push?

if [ $# -gt 0 ]; then
case $1 in
    push)
        docker push ${REMOTE}
        ;;
    *)
        echo "Argumento no reconocido: $1"
        ;;
esac 

fi



