// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dns_location

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDNSLocationResultEnvelope struct {
	Result ZeroTrustDNSLocationModel `json:"result,computed"`
}

type ZeroTrustDNSLocationModel struct {
	ID                    types.String                          `tfsdk:"id" json:"id,computed"`
	AccountID             types.String                          `tfsdk:"account_id" path:"account_id"`
	Name                  types.String                          `tfsdk:"name" json:"name"`
	ClientDefault         types.Bool                            `tfsdk:"client_default" json:"client_default"`
	DNSDestinationIPsID   types.String                          `tfsdk:"dns_destination_ips_id" json:"dns_destination_ips_id"`
	ECSSupport            types.Bool                            `tfsdk:"ecs_support" json:"ecs_support"`
	Endpoints             *ZeroTrustDNSLocationEndpointsModel   `tfsdk:"endpoints" json:"endpoints"`
	Networks              *[]*ZeroTrustDNSLocationNetworksModel `tfsdk:"networks" json:"networks"`
	CreatedAt             timetypes.RFC3339                     `tfsdk:"created_at" json:"created_at,computed"`
	DOHSubdomain          types.String                          `tfsdk:"doh_subdomain" json:"doh_subdomain,computed"`
	IP                    types.String                          `tfsdk:"ip" json:"ip,computed"`
	IPV4Destination       types.String                          `tfsdk:"ipv4_destination" json:"ipv4_destination,computed"`
	IPV4DestinationBackup types.String                          `tfsdk:"ipv4_destination_backup" json:"ipv4_destination_backup,computed"`
	UpdatedAt             timetypes.RFC3339                     `tfsdk:"updated_at" json:"updated_at,computed"`
}

type ZeroTrustDNSLocationEndpointsModel struct {
	DOH  *ZeroTrustDNSLocationEndpointsDOHModel  `tfsdk:"doh" json:"doh"`
	DOT  *ZeroTrustDNSLocationEndpointsDOTModel  `tfsdk:"dot" json:"dot"`
	IPV4 *ZeroTrustDNSLocationEndpointsIPV4Model `tfsdk:"ipv4" json:"ipv4"`
	IPV6 *ZeroTrustDNSLocationEndpointsIPV6Model `tfsdk:"ipv6" json:"ipv6"`
}

type ZeroTrustDNSLocationEndpointsDOHModel struct {
	Enabled      types.Bool                                        `tfsdk:"enabled" json:"enabled"`
	Networks     *[]*ZeroTrustDNSLocationEndpointsDOHNetworksModel `tfsdk:"networks" json:"networks"`
	RequireToken types.Bool                                        `tfsdk:"require_token" json:"require_token"`
}

type ZeroTrustDNSLocationEndpointsDOHNetworksModel struct {
	Network types.String `tfsdk:"network" json:"network"`
}

type ZeroTrustDNSLocationEndpointsDOTModel struct {
	Enabled  types.Bool                                        `tfsdk:"enabled" json:"enabled"`
	Networks *[]*ZeroTrustDNSLocationEndpointsDOTNetworksModel `tfsdk:"networks" json:"networks"`
}

type ZeroTrustDNSLocationEndpointsDOTNetworksModel struct {
	Network types.String `tfsdk:"network" json:"network"`
}

type ZeroTrustDNSLocationEndpointsIPV4Model struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}

type ZeroTrustDNSLocationEndpointsIPV6Model struct {
	Enabled  types.Bool                                         `tfsdk:"enabled" json:"enabled"`
	Networks *[]*ZeroTrustDNSLocationEndpointsIPV6NetworksModel `tfsdk:"networks" json:"networks"`
}

type ZeroTrustDNSLocationEndpointsIPV6NetworksModel struct {
	Network types.String `tfsdk:"network" json:"network"`
}

type ZeroTrustDNSLocationNetworksModel struct {
	Network types.String `tfsdk:"network" json:"network"`
}