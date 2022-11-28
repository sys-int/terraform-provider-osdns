package opnsense

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"terraform-sysint-os-dns/opnsense/unbound"
)

const OPN_URL = "opn_url"
const OPN_NOSSLVERIFY = "opn_nosslverify"
const OPN_APIKEY = "opn_apikey"
const OPN_APISECRET = "opn_apisecret"

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ConfigureFunc: providerConfigure,

		Schema: map[string]*schema.Schema{
			OPN_URL: {
				Type:     schema.TypeString,
				Required: true,
			},

			OPN_APIKEY: {
				Type:     schema.TypeString,
				Required: true,
			},

			OPN_APISECRET: {
				Type:     schema.TypeString,
				Required: true,
			},

			OPN_NOSSLVERIFY: {
				Type:     schema.TypeBool,
				Required: false,
				Default:  false,
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"dns_override": initHostOverride(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"dns_override": initHostOverride(),
		},
	}
}

func initHostOverride() *schema.Resource {
	return &schema.Resource{
		Create: unbound.CreateOverrideHost,
		Read:   unbound.ReadOverrideHost,
		Update: unbound.UpdateOverrideHost,
		Delete: unbound.DeleteOverrideHost,
		Exists: unbound.ExistsOverrideHost,
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
			},
			"host": {
				Type:     schema.TypeString,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rr": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mx": {
				Type:     schema.TypeString,
				Required: false,
			},
			"mxprio": {
				Type:     schema.TypeString,
				Required: false,
			},
			"description": {
				Type:     schema.TypeString,
				Required: false,
			},
			"uuid": {
				Type:     schema.TypeString,
				Required: false,
			},
		},
	}
}
