terraform {
  required_version = ">= 0.13"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 3.51.0"
    }
    twilio = {
      source  = "RJPearson94/twilio"
      version = ">= 0.14.0"
    }
  }
}

