# Deploy Backend - Terraform + AWS EC2 + Docker

Este entorno despliega el backend del proyecto en una instancia EC2 de AWS usando **Terraform** y **Docker**, asegurando una infraestructura reproducible y escalable.

## 1. Prerrequisitos

Antes de comenzar asegúrate de tener:

* private y public key del usuario en AWS
* Terraform instalado

## 2. Configura las variables de entorno

Crea el archivo `terraform.tfvars` o copialo de `terraform.tfvars.example` con la información de las variables de entorno del back y las claves del usuario con la siguiente estructura:

```env
aws_access_key = ""
aws_secret_key = ""
env_vars = {
  DATABASE_URL = ""
  API_URL      = ""
  API_TOKEN    = ""
}
```

## 3. genera un key pair para poder acceder a la instancia por SSH

Los archivos de la clave deben estar creadas en el directorio `../terraform/`

```bash
ssh-keygen -t rsa -b 4096 -f backend-key
```

## 4. Despliegue

Desde eñ directorio  `../terraform/`

```bash
terraform init
```

```bash
terraform apply -auto-approve
```

## 5. Verifica el despliegue

Una vez levantados los servios la terminar mostrará la dirección IP con la que se podrá acceder al servidor en el puero `8080`.

para probar el funcionamiento puede hacerce desde el navegador en la ruta `/stocks`

```bash
http://<IP-PUBLICA>:8080/stocks
```

o tambien con `curl`:

```bash
curl http://<IP-PUBLICA>:8080/stocks
```
