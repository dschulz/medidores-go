# medidores
## API para recolección de lecturas de medidores de agua

### Construir

```
  go get -u -v git.8bit.com.py/8Bit/medidores
  go mod download
  go build -a -o medidores cmd/app
```


### Ejecutar

- Primero crear un archivo con la configuración
 ```
    (pendiente)
```
- Ejecutar
```
source config
./medidores
```
   
## Docker

- Iniciar aplicación:
```
docker run -it --rm -p 3000:3000 dschulzg/medidores:latest
```

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
