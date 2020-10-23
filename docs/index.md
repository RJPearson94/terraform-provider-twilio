# Twilio Provider

The Twilio provider is used to interact with the many resources supported by Twilio. The provider needs to be configured with your Twilio credentials before it can be used.

> ⚠️ **Disclaimer**: This project is not an official Twilio project and is not supported or endorsed by Twilio in any way. It is maintained in [my](https://github.com/RJPearson94) free time.

## Installation

**NOTE:** This provider only supports Terraform 0.12+

### Terraform 0.13+

The provider has been published to the [Terraform Registry](https://registry.terraform.io/providers/RJPearson94/twilio/latest) you need to add the following code to your Terraform configuration and run terraform init. Terraform will take care of installing the provider for you.

```hcl
terraform {
  required_providers {
    twilio = {
      source  = "RJPearson94/twilio"
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

## Authentication

The Twilio provider offers a various way of providing credentials for authentication. The following methods are supported, in precedence order:

- Static credentials
  - API Key & Secret
  - Account SID & Auth Token
- Environment variables
  - API Key & Secret
  - Account SID & Auth Token

### Static credentials

!> This method is supported however it is not recommend for use as secrets could be leaked if the provider was committed to public version control.

#### API Key & Secret

Static credentials can be provided by adding an `account_sid`, `api_key` & `api_secret` in-line in the Twilio provider block:

Usage:

```hcl
provider "aws" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  api_key     = "SKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  api_secret  = "api-secret"
}
```

#### Account SID & Auth Token

Static credentials can be provided by adding an `account_sid` and `auth_token` in-line in the Twilio provider block:

Usage:

```hcl
provider "aws" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  auth_token  = "my-auth-token"
}
```

### Environment variables

#### API Key & Secret

You can provide your credentials via the `TWILIO_ACCOUNT_SID`, `TWILIO_API_KEY` and `TWILIO_API_SECRET` environment variables, representing your Twilio Account SID, API Key SID and API Secret respectively.

```hcl
provider "twilio" {}
```

Usage:

```sh
export TWILIO_ACCOUNT_SID="ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
export TWILIO_API_KEY="SKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
export TWILIO_API_SECRET="api-secret"
terraform plan
```

or

```sh
TWILIO_ACCOUNT_SID="ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" TWILIO_API_KEY="SKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" TWILIO_API_SECRET="api-secret" terraform plan
```

#### Account SID & Auth Token

You can provide your credentials via the `TWILIO_ACCOUNT_SID` and `TWILIO_AUTH_TOKEN` environment variables, representing your Twilio Account SID and Auth Token respectively.

```hcl
provider "twilio" {}
```

Usage:

```sh
export TWILIO_ACCOUNT_SID="ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
export TWILIO_AUTH_TOKEN="my-auth-token"
terraform plan
```

or

```sh
TWILIO_ACCOUNT_SID="ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" TWILIO_AUTH_TOKEN="my-auth-token" terraform plan
```

## Argument Reference

In addition to [generic provider arguments](https://www.terraform.io/docs/configuration/providers.html) the following arguments are supported:

- `account_sid` - (Optional) This is the Account Sid. This SID is mandatory, but it can also be retrieved from the `TWILIO_ACCOUNT_SID` environment variable
- `api_key` - (Optional) An API key SID associate with the account. This value can be retrieved from the `TWILIO_API_KEY` environment variable
- `api_secret` - (Optional) An secret value for the API Key. This value can be retrieved from the `TWILIO_API_SECRET` environment variable
- `auth_token` - (Optional) The Auth token for the account. This value can be retrieved from the `TWILIO_AUTH_TOKEN` environment variable

**NOTE:** A valid API Key and Secret or Auth Token must be supplied
