resource "aws_instance" "backend" {
  ami           = data.aws_ami.ubuntu.id
  instance_type = var.instance_type
  user_data = templatefile("${path.module}/scripts/startup.sh", {
    docker_image = var.docker_image
    env_vars     = var.env_vars
  })
  key_name                    = aws_key_pair.deployer.key_name
  associate_public_ip_address = true
  vpc_security_group_ids      = [aws_security_group.backend_sg.id]

  tags = {
    Name = "backend-instance"
  }
}

data "aws_ami" "ubuntu" {
  most_recent = true
  owners      = ["099720109477"] # Canonical

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
  }
}