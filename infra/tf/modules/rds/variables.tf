variable "db_name" {}
variable "allocated_storage" {
  default = 20
}
variable "username" {}
variable "password" {}
variable "vpc_id" {
  
}
variable "subnet_ids" {
  type = list(string)
}
