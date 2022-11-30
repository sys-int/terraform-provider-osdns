package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"terraform-sysint-os-dns/opnsense"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: opnsense.Provider})
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return opnsense.Provider()
		},
	})
}
