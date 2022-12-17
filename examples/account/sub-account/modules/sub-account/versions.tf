terraform {
  required_version = ">= 0.13.0"

  required_providers {
    twilio = {
      source  = "RJPearson94/twilio"
      version = ">= 0.21.0"
    }
  }
}
