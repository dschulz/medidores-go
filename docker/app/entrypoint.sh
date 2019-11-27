#!/usr/bin/env bash


ls
echo "Esperando base de datos..."
sleep 1

if [ x"$GO_ENV" == x"dev" ]; then
    /aplicacion/bin/migrate down
    /aplicacion/bin/migrate up
fi

if [ x"$GO_ENV" == x"prod" ]; then
    /aplicacion/bin/migrate up
fi

exec /aplicacion/bin/medidores
