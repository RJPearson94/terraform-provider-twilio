terraform {
  source = "../..//modules/verify"
}

include {
  path = find_in_parent_folders()
}

dependency "sub_account" {
  config_path = "../sub-account"

  mock_outputs = {
    account_details = {
      sid        = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
      auth_token = "auth_token"
    }
  }
  mock_outputs_allowed_terraform_commands = ["validate", "plan"]
}

generate "provider" {
  path      = "provider.tf"
  if_exists = "overwrite"
  contents = <<EOF
provider "twilio" {
  account_sid = "${dependency.sub_account.outputs.account_details.sid}"
  auth_token = "${dependency.sub_account.outputs.account_details.auth_token}"
}
EOF
}