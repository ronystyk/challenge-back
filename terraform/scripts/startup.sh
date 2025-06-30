#!/bin/bash
sudo apt-get update -y
sudo apt-get install -y docker.io
sudo systemctl start docker
sudo systemctl enable docker

# Crear el directorio si no existe
mkdir -p /home/ubuntu/challenge

# Variables de entorno
echo "🔧 Generando archivo .env..."
cat <<EOF > /home/ubuntu/challenge/.env
%{ for key, value in env_vars ~}
${key}=${value}
%{ endfor ~}
EOF

# Aseguramos permisos adecuados
chmod 600 /home/ubuntu/challenge/.env

# Ejecutamos el contenedor usando el .env
echo "🚀 Ejecutando contenedor..."
docker run -d \
  --name backend \
  --env-file /home/ubuntu/challenge/.env \
  -p 8080:8080 \
  ${docker_image}