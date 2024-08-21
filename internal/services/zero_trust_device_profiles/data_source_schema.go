// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_profiles

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &ZeroTrustDeviceProfilesDataSource{}

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"policy_id": schema.StringAttribute{
				Description: "Device ID.",
				Computed:    true,
				Optional:    true,
			},
			"allow_mode_switch": schema.BoolAttribute{
				Description: "Whether to allow the user to switch WARP between modes.",
				Optional:    true,
			},
			"allow_updates": schema.BoolAttribute{
				Description: "Whether to receive update notifications when a new version of the client is available.",
				Optional:    true,
			},
			"allowed_to_leave": schema.BoolAttribute{
				Description: "Whether to allow devices to leave the organization.",
				Optional:    true,
			},
			"auto_connect": schema.Float64Attribute{
				Description: "The amount of time in minutes to reconnect after having been disabled.",
				Optional:    true,
			},
			"captive_portal": schema.Float64Attribute{
				Description: "Turn on the captive portal after the specified amount of time.",
				Optional:    true,
			},
			"default": schema.BoolAttribute{
				Description: "Whether the policy is the default policy for an account.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the policy.",
				Optional:    true,
			},
			"disable_auto_fallback": schema.BoolAttribute{
				Description: "If the `dns_server` field of a fallback domain is not present, the client will fall back to a best guess of the default/system DNS resolvers unless this policy option is set to `true`.",
				Optional:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Whether the policy will be applied to matching devices.",
				Optional:    true,
			},
			"exclude_office_ips": schema.BoolAttribute{
				Description: "Whether to add Microsoft IPs to Split Tunnel exclusions.",
				Optional:    true,
			},
			"gateway_unique_id": schema.StringAttribute{
				Optional: true,
			},
			"lan_allow_minutes": schema.Float64Attribute{
				Description: "The amount of time in minutes a user is allowed access to their LAN. A value of 0 will allow LAN access until the next WARP reconnection, such as a reboot or a laptop waking from sleep. Note that this field is omitted from the response if null or unset.",
				Optional:    true,
			},
			"lan_allow_subnet_size": schema.Float64Attribute{
				Description: "The size of the subnet for the local access network. Note that this field is omitted from the response if null or unset.",
				Optional:    true,
			},
			"match": schema.StringAttribute{
				Description: "The wirefilter expression to match devices.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the device settings profile.",
				Optional:    true,
			},
			"precedence": schema.Float64Attribute{
				Description: "The precedence of the policy. Lower values indicate higher precedence. Policies will be evaluated in ascending order of this field.",
				Optional:    true,
			},
			"support_url": schema.StringAttribute{
				Description: "The URL to launch when the Send Feedback button is clicked.",
				Optional:    true,
			},
			"switch_locked": schema.BoolAttribute{
				Description: "Whether to allow the user to turn off the WARP switch and disconnect the client.",
				Optional:    true,
			},
			"tunnel_protocol": schema.StringAttribute{
				Description: "Determines which tunnel protocol to use.",
				Optional:    true,
			},
			"exclude": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							Description: "The address in CIDR format to exclude from the tunnel. If `address` is present, `host` must not be present.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "A description of the Split Tunnel item, displayed in the client UI.",
							Computed:    true,
						},
						"host": schema.StringAttribute{
							Description: "The domain name to exclude from the tunnel. If `host` is present, `address` must not be present.",
							Computed:    true,
							Optional:    true,
						},
					},
				},
			},
			"fallback_domains": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"suffix": schema.StringAttribute{
							Description: "The domain suffix to match when resolving locally.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "A description of the fallback domain, displayed in the client UI.",
							Computed:    true,
							Optional:    true,
						},
						"dns_server": schema.ListAttribute{
							Description: "A list of IP addresses to handle domain resolution.",
							Computed:    true,
							Optional:    true,
							ElementType: types.StringType,
						},
					},
				},
			},
			"include": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							Description: "The address in CIDR format to include in the tunnel. If address is present, host must not be present.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "A description of the split tunnel item, displayed in the client UI.",
							Computed:    true,
						},
						"host": schema.StringAttribute{
							Description: "The domain name to include in the tunnel. If host is present, address must not be present.",
							Computed:    true,
							Optional:    true,
						},
					},
				},
			},
			"service_mode_v2": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"mode": schema.StringAttribute{
						Description: "The mode to run the WARP client under.",
						Computed:    true,
						Optional:    true,
					},
					"port": schema.Float64Attribute{
						Description: "The port number when used with proxy mode.",
						Computed:    true,
						Optional:    true,
					},
				},
			},
			"target_tests": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The id of the DEX test targeting this policy",
							Computed:    true,
							Optional:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the DEX test targeting this policy",
							Computed:    true,
							Optional:    true,
						},
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
	}
}

func (d *ZeroTrustDeviceProfilesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustDeviceProfilesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("policy_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("policy_id")),
	}
}