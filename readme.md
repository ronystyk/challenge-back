# Challenge - Backend con Go + Gorilla Mux + GORM + CockroachDB + MVC + Unit tests

Este proyecto es una solución al reto propuesto. Conecta a una API externa, almacena los datos en CockroachDB usando GORM, y expone una API RESTful para consultar y sincronizar los datos.

## Configuración

### 1. Clona el repositorio

```bash
git clone https://github.com/ronystyk/challenge-back.git
cd challenge-back
```

### 2. Configura las variables de entorno

Crea un archivo `.env` o cópialo desde `.env.example` en la raíz del proyecto con esta estructura y asigna el valor de las variables.

```bash
DATABASE_URL=...
API_URL=...
API_KEY=...
```

### 3. Instala las dependencias

```go
go mod tidy
```

### 4. Ejecutar el servidor

```go
go run main.go
```

El servidor quedará disponible en `http://localhost:8080`

## Sincronizar los datos de la API con la base de datos

Consulta todos los registros de la API y los inserta en la base de datos de Cockroachlabs

```bash
curl --location --request POST 'http://localhost:8080/sync-stocks'
```

## Ejecutar pruebas unitarias

```bash
go test .\tests\
```

## Terraform IaC

Para el despliegue de la Infractructura como Código (IaC) revisa el archivo `readme.md` del directorio `./terraform/`
