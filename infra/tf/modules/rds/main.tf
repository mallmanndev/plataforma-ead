provider "aws" {
  region = "us-east-1" # ou substitua pela região desejada
}

resource "aws_security_group" "db-sg" {
  vpc_id = var.vpc_id
  egress {
    from_port       = 0
    to_port         = 0
    protocol        = "-1"
    cidr_blocks     = ["0.0.0.0/0"]
    prefix_list_ids = []
  }
  tags = {
    Name = "${var.db_name}-sg"
  }
}

resource "aws_db_subnet_group" "db_subnet" {
  name       = "db_subnet"
  subnet_ids = var.subnet_ids
}

resource "aws_db_instance" "postgres" {
  identifier              = var.db_name
  allocated_storage       = var.allocated_storage
  storage_type            = "gp2"
  engine                  = "postgres"
  engine_version          = "13.4"
  instance_class          = "db.t2.micro"
  db_name                 = var.db_name
  username                = var.username
  password                = var.password
  publicly_accessible     = false
  backup_retention_period = 7
  skip_final_snapshot     = true
  db_subnet_group_name    = aws_db_subnet_group.db_subnet.name
  vpc_security_group_ids  = [aws_security_group.db-sg.id]

  tags = {
    Name = var.db_name
  }
}
