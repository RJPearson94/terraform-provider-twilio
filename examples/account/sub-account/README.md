# Managing Sub-Accounts

This example shows how you can split your sub-account and sub-account resources into separate state files, then use terragrunt to manage the dependencies and automatically configure the sub-account resources provider with the created sub-account

In this example, there is a [modules](./modules/) which contains the terraform code to manage a [sub account](./modules/sub-account/main.tf) resource and another folder which contains the code to manage a [verify service](./modules/sub-account/main.tf) resource. Both modules contain a prefix variable and the sub-account module outputs the account details i.e. SID and auth_token and the verify module outputs the SID of the verify service.

**NOTE:**: The verify service was used as an example, you can configure 1 or more resources that are supported by sub-accounts

The terragrunt code can be found in the [deployment](./deployment/) folder. There is a folder and terragrunt file per module, which contains the necessary config to configure/ manage that resource. A terragrunt.hcl file also exists at the root of the folder to include common config and inputs/ variables.

In the [verify terragrunt.hcl](./deployment/verify/terragrunt.hcl) file there is config to create a link/ dependency on the sub-account resources. This ensures that Terragrunt will configure the sub-account resources before provisioning the verify resources. To allow you to plan and validate the Terraform/ Terragrunt config, defaults are added. When applying the changes, the sub-account resources will be provisioned, and then using the outputs from the module terragrunt will generate a provider.tf file that has the account sid and auth token values substituted in. This file is written into the module allowing the terraform code to use the provider that has been configured for the sub-account.

## Running the code

### Prerequisite

- A Twilio account
  - The Account SID set as the TWILIO_ACCOUNT_SID environment variable
  - The Auth Token set as the TWILIO_AUTH_TOKEN environment variable
- Terraform (v1.3.4 was used during testing)
- Terragrunt (v0.37.0 was used during testing)

### Plan

To plan the changes, you will need to cd into the deployments folder and run the following command

```sh
terragrunt run-all plan
```

This will perform a dry-run and tell you what resources will be provisions

### Apply

To apply the changes, you will need to cd into the deployments folder and run the following command

```sh
terragrunt run-all apply
```

You will need to type `y` when the prompted

This will create the sub-account and verify service in the sub-account

### Destroy

To destroy the resources, you will need to cd into the deployments folder and run the following command

```sh
terragrunt run-all destroy
```

You will need to type `y` when the prompted

This will delete the verify service and sub-account
