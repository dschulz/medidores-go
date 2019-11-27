# medidores
## API para recolección de lecturas de medidores de agua

### NOTA: ESTO ES EXPERIMENTAL. Faltan muchas cosas, esto es solo un esqueleto por ahora.

*Entre otras cosas, faltan:*

- [ ] Los modelos de datos y sus migraciones
- [ ] Las definiciones de rutas
- [ ] Los handlers para las definiciones de rutas
- [ ] Arreglar docker-compose.yml para que se pueda levantar un servidor PostgreSQL containerizado y luego la aplicación
- [ ] Documentar
- [ ] Implementar integración continua
- [ ] Tests de todo

*Lo que ya está mas o menos definido y funcional:*

- [x] Organización de packages
- [x] Packages básicos para validación, logging, routers, configuración, etc.
- [x] Dockerfile funcional
- [x] Scripts para construir imágenes y subir a hub.docker.com

### Construir

```
git clone git@git.8bit.com.py:8Bit/medidores.git
cd medidores
go mod download

go build -a -o migrate cmd/migrate
go build -a -o medidores cmd/app
```


### Ejecutar


- Ejecutar

```

./local.sh

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
