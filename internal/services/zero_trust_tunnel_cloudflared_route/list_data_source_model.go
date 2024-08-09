// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_route

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelCloudflaredRoutesResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustTunnelCloudflaredRoutesResultDataSourceModel `json:"result,computed"`
}

type ZeroTrustTunnelCloudflaredRoutesDataSourceModel struct {
	AccountID        types.String                                              `tfsdk:"account_id" path:"account_id"`
	Comment          types.String                                              `tfsdk:"comment" query:"comment"`
	ExistedAt        timetypes.RFC3339                                         `tfsdk:"existed_at" query:"existed_at"`
	IsDeleted        types.Bool                                                `tfsdk:"is_deleted" query:"is_deleted"`
	NetworkSubset    types.String                                              `tfsdk:"network_subset" query:"network_subset"`
	NetworkSuperset  types.String                                              `tfsdk:"network_superset" query:"network_superset"`
	RouteID          types.String                                              `tfsdk:"route_id" query:"route_id"`
	TunTypes         types.String                                              `tfsdk:"tun_types" query:"tun_types"`
	TunnelID         types.String                                              `tfsdk:"tunnel_id" query:"tunnel_id"`
	VirtualNetworkID types.String                                              `tfsdk:"virtual_network_id" query:"virtual_network_id"`
	MaxItems         types.Int64                                               `tfsdk:"max_items"`
	Result           *[]*ZeroTrustTunnelCloudflaredRoutesResultDataSourceModel `tfsdk:"result"`
}

type ZeroTrustTunnelCloudflaredRoutesResultDataSourceModel struct {
	ID                 types.String      `tfsdk:"id" json:"id"`
	Comment            types.String      `tfsdk:"comment" json:"comment"`
	CreatedAt          timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	DeletedAt          timetypes.RFC3339 `tfsdk:"deleted_at" json:"deleted_at,computed"`
	Network            types.String      `tfsdk:"network" json:"network"`
	TunType            types.String      `tfsdk:"tun_type" json:"tun_type"`
	TunnelID           types.String      `tfsdk:"tunnel_id" json:"tunnel_id"`
	TunnelName         types.String      `tfsdk:"tunnel_name" json:"tunnel_name"`
	VirtualNetworkID   types.String      `tfsdk:"virtual_network_id" json:"virtual_network_id"`
	VirtualNetworkName types.String      `tfsdk:"virtual_network_name" json:"virtual_network_name"`
}
