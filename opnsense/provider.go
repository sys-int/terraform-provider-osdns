package opnsense

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-sysint-os-dns/opnsense/unbound"
)

const OPN_URL = "opnsense_url"
const OPN_NOSSLVERIFY = "opnsense_nosslverify"
const OPN_APIKEY = "opnsense_apikey"
const OPN_APISECRET = "opnsense_apisecret"

func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureContextFunc: providerConfigure,

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
				Optional: true,
				Default:  false,
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"opnsense_unbound_hostoverride": unbound.ReadHostOverride(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"opnsense_unbound_hostoverride": unbound.HostOverride(),
		},
	}
}
