package main

import (
	"context"
	"flag"
	"log"

	"github.com/RJPearson94/terraform-provider-twilio/twilio"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: twilio.Provider,
	}

	if debugMode {
		// For information on debugging the provider see, https://www.terraform.io/docs/extend/debugging.html
		if err := plugin.Debug(context.Background(), "registry.terraform.io/RJPearson94/terraform-provider-twilio", opts); err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	plugin.Serve(opts)
}
