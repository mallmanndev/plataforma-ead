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

module "ecr" {
  source = "./modules/ecr"
}

module "ecs" {
  source = "./modules/ecs"
  prefix = var.prefix
  vpc_id = module.new-vpc.vpc_id
  subnet_ids = module.new-vpc.public_subnet_ids
}
