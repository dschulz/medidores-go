#!/usr/bin/env sh

IMAGEN="medidores"
VERSION_TAG="1.0"
REMOTE="dschulzg/${IMAGEN}:${VERSION_TAG}"

# Construir imagen pasando la fecha/hora como build arg
#docker build --no-cache=true --build-arg BUILD_TSTAMP=$(date -u +'%Y-%m-%dT%H:%M:%SZ')  --build-arg VERSION_TAG=$VERSION_TAG -t ${IMAGEN}:${VERSION_TAG} .


echo -e "\033[01;32mdocker tag ${IMAGEN}:${VERSION_TAG}\033[0m"
docker tag ${IMAGEN}:${VERSION_TAG} ${REMOTE}
echo -e "\033[01;32mdocker push ${REMOTE}\033[0m"
docker push ${REMOTE}
