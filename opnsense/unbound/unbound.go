package unbound

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func HostOverride() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceHostOverrideCreate,
		ReadContext:   resourceHostOverrideRead,
		UpdateContext: resourceHostOverrideUpdate,
		DeleteContext: resourceHostOverrideDelete,
		Schema:        hostOverrideSchema(),
	}
}

func ReadHostOverride() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceHostOverrideRead,
		Schema:      hostOverrideSchema(),
	}
}

func hostOverrideSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"host": {
			Type:     schema.TypeString,
			Required: true,
		},
		"domain": {
			Type:     schema.TypeString,
			Required: true,
		},
		"ip": {
			Type:     schema.TypeString,
			Required: true,
		},
		"mx": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"mxprio": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}
