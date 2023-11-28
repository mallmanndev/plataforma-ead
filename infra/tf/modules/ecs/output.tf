output "backend_sg_id" {
  value = aws_security_group.backend-sg.id
}

output "frontend_sg_id" {
  value = aws_security_group.frontend-sg.id
}
