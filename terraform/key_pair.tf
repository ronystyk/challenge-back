resource "aws_key_pair" "deployer" {
  key_name   = "backend-key"
  public_key = file("${path.module}/backend-key.pub")
}