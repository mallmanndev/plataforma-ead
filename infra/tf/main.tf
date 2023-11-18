terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
    local = ">=2.1.0"
  }
  backend "s3" {
    bucket  = "mallmanndev-tf-state"
    key     = "terraform.tfstate"
    region  = "us-east-1"
    profile = "matheus-admin"
  }
}

provider "aws" {
  region  = "us-east-1"
  profile = "matheus-admin"
}

module "new-vpc" {
  source         = "./modules/vpc"
  prefix         = var.prefix
  vpc_cidr_block = var.vpc_cidr_block
}

/*
module "eks" {
  source              = "./modules/eks"
  prefix              = var.prefix
  vpc_id              = module.new-vpc.vpc_id
  subnet_ids          = module.new-vpc.public_subnet_ids
  cluster_name        = var.cluster_name
  logs_retention_days = var.logs_retention_days
  desired_size        = var.desired_size
  max_size            = var.max_size
  min_size            = var.min_size
}
*/

module "ecr" {
  source = "./modules/ecr"
}

module "rds" {
  source     = "./modules/rds"
  db_name    = "service-core-db"
  vpc_id     = module.new-vpc.vpc_id
  subnet_ids = module.new-vpc.private_subnet_ids
  username   = "postgres"
  password   = "2RYoaq4iL&P#5gd$x"
}
