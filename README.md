# Base datos Mongo

Docker
```
Crear el contenedor
docker run --name mongodb-local -d --rm -e MONGO_INITDB_ROOT_USERNAME=user -e MONGO_INITDB_ROOT_PASSWORD=password -p 27017:27017 mongo:3.6.8

docker ps -a
CONTAINER_ID c59244668908

Copiar .json en la ruta /tmp/products.json dentro del contenedor
docker cp products.json c59244668908:/tmp/products.json

Valiadar
docker exec -it c59244668908 sh
cat /tmp/products.json

importar registros en la bd
docker exec c59244668908 mongoimport -u user -p password --authenticationDatabase admin -d products -c products --file /tmp/products.json

Valiadar
docker exec -it c59244668908 sh
cat /tmp/products.json
```



# Git
```
gh auth login
gh repo create yogparra/go-fiber-rest-api
```

# Go
```
go mod init github/yogparra/go-fiber-rest-api
go mod tidy
```