variable "vpc_id" {}
variable "prefix" {}
variable "subnet_ids" {
  type = list(string)
}
