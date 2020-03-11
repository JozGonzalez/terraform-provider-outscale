package outscale

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	oscgo "github.com/marinsalinas/osc-sdk-go"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceOutscaleOAPIVirtualGateways() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOutscaleOAPIVirtualGatewaysRead,

		Schema: map[string]*schema.Schema{
			"filters": dataSourceFiltersSchema(),
			"virtual_gateway_id": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"virtual_gateways": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"connection_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"virtual_gateway_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"net_to_virtual_gateway_links": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"state": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"net_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"tags": tagsOAPIListSchemaComputed(),
					},
				},
			},
			"request_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceOutscaleOAPIVirtualGatewaysRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).OSCAPI

	filters, filtersOk := d.GetOk("filters")
	_, vpnOk := d.GetOk("virtual_gateway_id")

	if !filtersOk && !vpnOk {
		return fmt.Errorf("One of virtual_gateway_id or filters must be assigned")
	}

	params := oscgo.ReadVirtualGatewaysRequest{}

	if filtersOk {
		params.SetFilters(buildOutscaleAPIVirtualGatewayFilters(filters.(*schema.Set)))
	}

	var resp oscgo.ReadVirtualGatewaysResponse
	var err error

	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		resp, _, err = conn.VirtualGatewayApi.ReadVirtualGateways(context.Background(), &oscgo.ReadVirtualGatewaysOpts{ReadVirtualGatewaysRequest: optional.NewInterface(params)})
		if err != nil {
			if strings.Contains(err.Error(), "RequestLimitExceeded:") {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return resource.NonRetryableError(err)
	})

	if err != nil {
		return err
	}
	if resp.GetVirtualGateways() == nil || len(resp.GetVirtualGateways()) == 0 {
		return fmt.Errorf("no matching VPN gateway found: %#v", params)
	}

	vpns := make([]map[string]interface{}, len(resp.GetVirtualGateways()))

	for k, v := range resp.GetVirtualGateways() {
		vpn := make(map[string]interface{})
		vs := make([]map[string]interface{}, len(v.GetNetToVirtualGatewayLinks()))

		for k, v1 := range v.GetNetToVirtualGatewayLinks() {
			vp := make(map[string]interface{})
			vp["state"] = v1.GetState()
			vp["net_id"] = v1.GetNetId()

			vs[k] = vp
		}
		vpn["net_to_virtual_gateway_links"] = vs
		vpn["state"] = v.GetState()
		vpn["connection_type"] = v.GetConnectionType()
		vpn["virtual_gateway_id"] = v.GetVirtualGatewayId()
		vpn["tags"] = tagsOSCAPIToMap(v.GetTags())

		vpns[k] = vpn
	}
	d.Set("virtual_gateways", vpns)
	d.Set("request_id", resp.ResponseContext.RequestId)
	d.SetId(resource.UniqueId())

	return nil
}
