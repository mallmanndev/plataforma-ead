resource "aws_security_group" "backend-sg" {
  name        = "backend-sg"
  description = "Security Group for Backend Services"

  vpc_id = var.vpc_id # Substitua pelo ID da sua VPC

  // Regras de entrada
  ingress {
    from_port   = 3000
    to_port     = 3000
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  // Regras de saída (ajuste conforme necessário)
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "frontend-sg" {
  name        = "frontend-sg"
  description = "Security Group for Frontend Services"

  vpc_id = var.vpc_id # Substitua pelo ID da sua VPC

  // Regras de entrada
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  // Regras de saída (ajuste conforme necessário)
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_ecs_cluster" "plataforma-ead-cluster" {
  name = "plataforma-ead-cluster"
}

resource "aws_ecs_task_definition" "backend_task" {
  family                   = "backend-task"
  network_mode             = "bridge"
  requires_compatibilities = ["EC2"]
  container_definitions = jsonencode([
    {
      name      = "backend"
      image     = "315655993037.dkr.ecr.us-east-1.amazonaws.com/backend:latest"
      cpu       = 1
      memory    = 512
      essential = true
      portMappings = [
        {
          containerPort = 3000
          hostPort      = 3000
        }
      ]
    }
  ])
}

resource "aws_ecs_service" "backend_service" {
  name            = "backend-service"
  cluster         = aws_ecs_cluster.plataforma-ead-cluster.id
  task_definition = aws_ecs_task_definition.backend_task.arn
  launch_type     = "EC2"
  desired_count   = 1 # número de instâncias do frontend

  network_configuration {
    subnets = var.subnet_ids  # substitua pela sua subnet
  }
}
