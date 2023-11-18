variable "vpc_id" {}
variable "prefix" {}
variable "cluster_name" {}
variable "logs_retention_days" {
  default = 30
}
variable "subnet_ids" {
  type = list(string)
}
variable "desired_size" {
  default = 1
}
variable "max_size" {
  default = 2
}
variable "min_size" {
  default = 1
}
