package opnsense

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	opn_unbound "github.com/sys-int/opnsense-api/api/unbound"
)

func resourceHostOverrideCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Debug(ctx, "creating unbound override host")
	var diags diag.Diagnostics
	client := meta.(ProviderClient)
	tflog.Debug(ctx, fmt.Sprintf("conn url=\"%s\" key=\"%s\" secret=\"%s\"", client.Conn.BaseUrl.String(), client.Conn.ApiKey, client.Conn.ApiSecret))
	api := opn_unbound.UnboundApi{client.Conn}

	host := unmarshalHost(ctx, d, &opn_unbound.HostOverride{})
	tflog.Debug(ctx, "pre-create", map[string]interface{}{
		"enabled":     host.Enabled,
		"uuid":        host.Uuid,
		"host":        host.Host,
		"domain":      host.Domain,
		"ip":          host.Ip,
		"mx":          host.Mx,
		"mxprio":      host.Mxprio,
		"description": host.Description,
	})
	tflog.Debug(ctx, "now calling host to create override host")
	uuid, err := api.HostOverrideCreate(*host)
	tflog.Debug(ctx, "processed uuid="+uuid)

	if err != nil {
		tflog.Error(ctx, "error creating override host "+err.Error())
		return diag.FromErr(err)
	}

	d.SetId(uuid)
	resourceHostOverrideRead(ctx, d, meta)
	return diags
}

func resourceHostOverrideRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Debug(ctx, "read unbound override host")
	var diags diag.Diagnostics
	client := meta.(ProviderClient)
	api := opn_unbound.UnboundApi{client.Conn}
	host, err := api.HostEntryGetByUuid(d.Id())
	host.Uuid = d.Id()
	tflog.Debug(ctx, "read host", map[string]interface{}{
		"enabled":     host.Enabled,
		"uuid":        host.Uuid,
		"host":        host.Host,
		"domain":      host.Domain,
		"ip":          host.Ip,
		"mx":          host.Mx,
		"mxprio":      host.Mxprio,
		"description": host.Description,
	})
	if err != nil {
		return diag.FromErr(err)
	}
	marshalHost(ctx, d, host)
	return diags
}

func resourceHostOverrideDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Debug(ctx, "delete unbound override host")
	var diags diag.Diagnostics
	client := meta.(ProviderClient)
	api := opn_unbound.UnboundApi{client.Conn}
	err := api.HostEntryRemove(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceHostOverrideUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Debug(ctx, "update unbound override host")
	var diags diag.Diagnostics
	client := meta.(ProviderClient)
	api := opn_unbound.UnboundApi{client.Conn}
	host := unmarshalHost(ctx, d, &opn_unbound.HostOverride{})
	uuid, err := api.HostOverrideUpdate(*host)
	host.Uuid = uuid
	marshalHost(ctx, d, *host)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

//
//func ExistsOverrideHost(ctx context.Context, d *schema.ResourceData, meta interface{}) (bool, diag.Diagnostics) {
//
//	var diags diag.Diagnostics
//	client := meta.(ProviderClient)
//	api := opn_unbound.UnboundApi{client.Conn}
//	host, err := api.HostEntryGetByFQDN(d.Get("host").(string), d.Get("domain").(string))
//	if err != nil {
//		return true, diag.FromErr(err)
//	}
//	if err != nil {
//		return false, nil
//	}
//	if &host != nil {
//		return true, nil
//	}
//	return false, diags
//}
