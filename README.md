# medidores
## API para recolección de lecturas de medidores de agua

### Construir

```
git clone git@git.8bit.com.py:8Bit/medidores.git
cd medidores
go mod download
go build -a -o medidores cmd/app
```


### Ejecutar

- Primero crear un archivo con la configuración. Se necesita tener un servidor PostgreSQL corriendo. La configuración debe contener usuario, password, host, etc.

```
    (pendiente)
```

- Ejecutar

```
source config
./medidores

```
   
## Docker

Para iniciar aplicación con la última versión subida al [registry de Docker](htps://hub.docker.com)

```

docker run -it --rm -p 3000:3000 dschulzg/medidores:latest

```

La opción `--rm` elimina el container una vez que la aplicación termina.

*PENDIENTE:* se debe tener un volumen con el archivo de configuración y pasar con el argumento `-v` o `--mount type=bind,source...,target...`


### docker-compose

- Iniciar aplicación y servidor PostgreSQL

```  
docker-compose up
```

- Detener aplicación y servidor PostgreSQL

```  
docker-compose down
```
