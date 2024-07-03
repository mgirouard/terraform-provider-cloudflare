// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel_route

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TunnelRouteResultEnvelope struct {
	Result TunnelRouteModel `json:"result,computed"`
}

type TunnelRoutesResultDataSourceEnvelope struct {
	Result TunnelRoutesDataSourceModel `json:"result,computed"`
}

type TunnelRouteModel struct {
	ID               types.String `tfsdk:"id" json:"id,computed"`
	AccountID        types.String `tfsdk:"account_id" path:"account_id"`
	Network          types.String `tfsdk:"network" json:"network"`
	Comment          types.String `tfsdk:"comment" json:"comment"`
	VirtualNetworkID types.String `tfsdk:"virtual_network_id" json:"virtual_network_id"`
	CreatedAt        types.String `tfsdk:"created_at" json:"created_at,computed"`
	DeletedAt        types.String `tfsdk:"deleted_at" json:"deleted_at,computed"`
	TunnelID         types.String `tfsdk:"tunnel_id" json:"tunnel_id,computed"`
}

type TunnelRoutesDataSourceModel struct {
}