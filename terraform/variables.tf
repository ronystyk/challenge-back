variable "aws_region" {
  default = "us-east-1"
}

variable "aws_access_key" {
  type      = string
  sensitive = true
}

variable "aws_secret_key" {
  type      = string
  sensitive = true
}

variable "instance_type" {
  default = "t2.micro"
}

variable "docker_image" {
  description = "Docker image from Docker Hub"
  default     = "ronystyk/challenge-back:latest"
}

variable "env_vars" {
  description = "Environment variables for backend"
  type        = map(string)
}