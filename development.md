# Developing the Provider

These instruction are here to help aid you in setting up your development environment to allow you to build and test the Twilio provider.

**NOTE:** Currently you will have to set it up on your own machine as [Github Codespaces](https://github.com/features/codespaces/) are currently in Beta. Once this feature becomes available I will try to add support ASAP.

## Prerequisites

- [Terraform](https://www.terraform.io/downloads.html) 0.12+
- [Go](https://golang.org/doc/install) 1.16 (to build the provider plugin)

**Note:** This project uses [Go Modules](https://blog.golang.org/using-go-modules)

## Getting started

This project can either be cloned inside or outside your [GOPATH](http://golang.org/doc/code.html#GOPATH) The example will show cloning within your GOPATH

Clone repository to: `$GOPATH/src/github.com/RJPearson94/terraform-provider-twilio`

```sh
mkdir -p $GOPATH/src/github.com/RJPearson94; cd $GOPATH/src/github.com/RJPearson94
$ git clone git@github.com:RJPearson94/terraform-provider-twilio
```

Enter the provider directory and run `make tools`. To download all the tools necessary to build & test the provider.

```sh
make tools
```

To build the provider binary to the `$GOPATH/bin` directory run `make build`

```sh
make build
...
$GOPATH/bin/terraform-provider-twilio
...
```

## Testing

In order to test the provider, run the following command

```sh
make test
```

In order to run the Acceptance tests, run the following command

**Warning:** These test will provision real resources on Twilio and could cost money.

```sh
make testacc
```
