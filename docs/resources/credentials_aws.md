---
page_title: "Twilio AWS Credentials"
subcategory: "Credentials"
---

# twilio_credentials_aws Resource

Manages an AWS credential resource. This resource allows you to upload a set of AWS credentials to Twilio for various services to use to access resources in your AWS account

!> If the `account_sid` is managed via Terraform and the `account_sid` is removed from the configuration file. The old value will be retained on the next apply.

## Example Usage

### Basic

```hcl
resource "twilio_credentials_aws" "aws" {
  friendly_name         = "Test AWS Credential Resource"
  aws_access_key_id     = "aws_access_key_id"
  aws_secret_access_key = "aws_secret_access_key"
}
```

### With AWS & Twilio providers

```hcl
resource "aws_iam_user" "iam_user" {
  name = "test-iam-user"
}

resource "aws_iam_access_key" "access_key" {
  user = aws_iam_user.iam_user.name
}

resource "twilio_credentials_aws" "aws" {
  friendly_name         = "Test AWS Credential Resource"
  aws_access_key_id     = aws_iam_access_key.access_key.id
  aws_secret_access_key = aws_iam_access_key.access_key.secret
}
```

## Argument Reference

The following arguments are supported:

- `account_sid` - (Optional) The SID of a sub account to associate the public key resource with. Changing this forces a new resource to be created
- `aws_secret_access_key` - (Mandatory) The AWS Secret Access Key to associate with the AWS credential resource. Changing this forces a new resource to be created
- `aws_access_key_id` - (Mandatory) The AWS Access Key ID to associate with the AWS credential resource. Changing this forces a new resource to be created
- `friendly_name` - (Optional) The friendly name of the AWS credential resource

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the AWS credential resource (Same as the `sid`)
- `sid` - The SID of the AWS credential resource (Same as the `id`)
- `account_sid` - The account SID associated with the AWS credential resource
- `friendly_name` - The friendly name of the AWS credential resource
- `date_created` - The date in RFC3339 format that the AWS credential resource was created
- `date_updated` - The date in RFC3339 format that the AWS credential resource was updated
- `url` - The URL of the AWS credential resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the AWS credential resource
- `update` - (Defaults to 10 minutes) Used when updating the AWS credential resource
- `read` - (Defaults to 5 minutes) Used when retrieving the AWS credential resource
- `delete` - (Defaults to 10 minutes) Used when deleting the AWS credential resource
