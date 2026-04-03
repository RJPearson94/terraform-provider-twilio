---
page_title: "twilio_credentials_aws Resource - twilio"
subcategory: "Credentials"
description: |-
  
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

## Schema

### Required

- `aws_access_key_id` (String) The AWS access key ID used to authenticate with AWS services. Changing this forces a new resource
- `aws_secret_access_key` (String, Sensitive) The AWS secret access key used to authenticate with AWS services. Sensitive -- will not be shown in logs or plans. Changing this forces a new resource

### Optional

- `account_sid` (String) The SID of the account that owns this AWS credential. Changing this forces a new resource
- `friendly_name` (String) A human-readable label for the AWS credential
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `date_created` (String) The date and time the AWS credential was created, in RFC 3339 format
- `date_updated` (String) The date and time the AWS credential was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this AWS credential by Twilio
- `url` (String) The absolute URL of the AWS credential resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)
