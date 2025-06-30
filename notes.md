# Apuntes

## Go

1. Inicializar un nuevo proyecto `go mod init challenge-back`
2. Librería para variables de entorno `go get github.com/joho/godotenv`
3. Instalar el Paquete resty para manejo de peticiones HTTP `go get github.com/go-resty/resty/v2`
4. Instalar postgreSQL `go get github.com/jackc/pgx/v5/pgxpool`
5. Ejecutar el programa `go run .`
6. Instalar framework web Gorilla Mux `go get github.com/gorilla/mux`
7. Instalar un ORM para la gestión de los datos (GORM) `go get -u gorm.io/gorm` y el driver para PostgreSQL `go get -u gorm.io/driver/postgres` y un driver para una base de datos SQLite para testings `go get github.com/glebarez/sqlite`
8. para cors `go get github.com/gorilla/handlers`

## Terraform

1. `terraform init` Inicializa Terraform
2. `terraform plan` Revisa el plan de ejecución
3. `terraform apply` Aplica los cambios (despliegue) `-auto-approve` para evitar el mensaje de confirmación
4. `terraform fmt` Aplica formato a los archivos
5. `terraform destroy` Elimina toda la insfractructura completa
6. `ssh-keygen -t rsa -b 4096 -f backend-key` generar private and public key para SSH si es necesario
7. por si necesito revisar el log del servidor `cat /var/log/cloud-init-output.log`
8. Verificar que se esté escuchando el puerto `sudo ss -tuln | grep :8080`
