package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"terraform-sysint-os-dns/opnsense"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: opnsense.Provider})
}
