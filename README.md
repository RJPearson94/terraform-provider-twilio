# Terraform Provider for Twilio

![Terraform Provider Checks](https://github.com/RJPearson94/terraform-provider-twilio/workflows/Terraform%20Provider%20Checks/badge.svg)
[![Terraform Registry](https://img.shields.io/badge/registry-twilio-green?logo=terraform&style=flat)](https://registry.terraform.io/providers/RJPearson94/twilio/latest)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/RJPearson94/terraform-provider-twilio)](https://pkg.go.dev/github.com/RJPearson94/terraform-provider-twilio)
[![Release](https://img.shields.io/github/release/RJPearson94/terraform-provider-twilio.svg)](https://github.com/RJPearson94/terraform-provider-twilio/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/RJPearson94/terraform-provider-twilio)](https://goreportcard.com/report/github.com/RJPearson94/terraform-provider-twilio)
[![License](https://img.shields.io/github/license/RJPearson94/terraform-provider-twilio)](/LICENSE)

The Terraform Twilio provider is a plugin for [Terraform](https://www.terraform.io/) that allows for the lifecycle management of supported Twilio resources.

> ⚠️ **Disclaimer**: This project is not an official Twilio project and is not supported or endorsed by Twilio in any way. It is maintained in [my](https://github.com/RJPearson94) free time.

## Getting Started

- [Using Provider](./docs/index.md)
- [Terraform Registry](https://registry.terraform.io/providers/RJPearson94/twilio/latest)
- [Developing the Provider](./development.md)

**NOTE:** The default branch for this project is called `main`

## Documentation

Documentation of the provider and all supported resources can be found [here](./docs)
Documentation on managing sub-accounts can be found [here](./examples/account/sub-account/README.md)

## Installation

**NOTE:** This provider only supports Terraform 0.12+

### Terraform 0.13+

The provider has been published to the [Terraform Registry](https://registry.terraform.io/providers/RJPearson94/twilio/latest) you need to add the following code to your Terraform configuration and run terraform init. Terraform will take care of installing the provider for you.

```hcl
terraform {
  required_providers {
    twilio = {
      source = "RJPearson94/twilio"
      version = ">= 0.2.1"
    }
  }
}

provider "twilio" {
  # Configuration options
}
```

### Terraform 0.12

This is a bit more work as you have to download the [latest release](https://github.com/RJPearson94/terraform-provider-twilio/releases/latest) of the terraform provider which can run on you machine operating system/ processor architecture. Then unzip the provider and place the provider in the `~/.terraform.d/plugins` folder (on most operating systems) and `%APPDATA%\terraform.d\plugins` on Windows. For more information see the [terraform docs](https://www.terraform.io/docs/extend/how-terraform-works.html#plugin-locations)

## Dependencies

- [Twilio Go SDK](https://github.com/RJPearson94/twilio-sdk-go)
- [Terraform Plugin SDK](https://github.com/hashicorp/terraform-plugin-sdk)
- [go-homdir](https://github.com/mitchellh/go-homedir)
