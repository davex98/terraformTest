provider "aws" {
  region = "us-east-2"
}

resource "aws_instance" "hello_world" {
  ami                    = "ami-0d5d9d301c853a04a"
  instance_type          = "t2.micro"
  vpc_security_group_ids = [aws_security_group.hello_world.id]

  user_data =  <<EOF
#!/bin/bash
echo "Hello Terraform!" > index.html
nohup busybox httpd -f -p 8888 &
EOF
}

resource "aws_security_group" "hello_world" {
  ingress {
    from_port   = 8888
    to_port     = 8888
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}