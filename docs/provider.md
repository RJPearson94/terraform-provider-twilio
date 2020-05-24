# Twilio Provider

The Twilio provider is used to interact with the many resources supported by Twilio. The provider needs to be configured with the proper credentials before it can be used.

View the [resources](./resources) folder and the [data sources](./data_sources) folder to see all the resources that are available.

## Authentication

The Twilio provider offers a various way of providing credentials for
authentication. The following methods are supported, in precedence order:

- Static credentials *(Not Recommended)*
- Environment variables

### Static credentials *(Not Recommended)*

**Warning:** This method is supported however it is not recommend for use as secrets could be leaked if the provider was committed to public version control.

Static credentials can be provided by adding an `account_sid` and `auth_token` in-line in the Twilio provider block:

Usage:

```hcl
provider "aws" {
  account_sid = "ACxxxxxxxxxxxxx"
  auth_token = "my-auth-token"
}
```

### Environment variables

You can provide your credentials via the `TWILIO_ACCOUNT_SID` and `TWILIO_AUTH_TOKEN` environment variables, representing your Twilio Account SID and Auth Token respectively.

```hcl
provider "twilio" {}
```

Usage:

```sh
export TWILIO_ACCOUNT_SID="ACxxxxxxxxxxxxx"
export TWILIO_AUTH_TOKEN="my-auth-token"
terraform plan
```

or

```sh
TWILIO_ACCOUNT_SID="ACxxxxxxxxxxxxx" TWILIO_AUTH_TOKEN="my-auth-token" terraform plan
```

## Argument Reference

In addition to [generic provider arguments](https://www.terraform.io/docs/configuration/providers.html) the following arguments are supported:

- `account_sid` - (Optional) This is the Account Sid. This SID is mandatory, but it can also be retrieved from the `TWILIO_ACCOUNT_SID` environment variable

- `auth_token` - (Optional) This is the Auth token. This token is mandatory, but it can also be retrieved from the `TWILIO_AUTH_TOKEN` environment variable.

**Note:** In the future there is plans to support API Key and Secret Authentication too
