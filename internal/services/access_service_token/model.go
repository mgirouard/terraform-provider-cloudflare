// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_service_token

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessServiceTokenResultEnvelope struct {
	Result AccessServiceTokenModel `json:"result,computed"`
}

type AccessServiceTokenResultDataSourceEnvelope struct {
	Result AccessServiceTokenDataSourceModel `json:"result,computed"`
}

type AccessServiceTokensResultDataSourceEnvelope struct {
	Result AccessServiceTokensDataSourceModel `json:"result,computed"`
}

type AccessServiceTokenModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	AccountID    types.String `tfsdk:"account_id" path:"account_id"`
	ZoneID       types.String `tfsdk:"zone_id" path:"zone_id"`
	Name         types.String `tfsdk:"name" json:"name"`
	Duration     types.String `tfsdk:"duration" json:"duration"`
	ClientID     types.String `tfsdk:"client_id" json:"client_id,computed"`
	ClientSecret types.String `tfsdk:"client_secret" json:"client_secret,computed"`
	CreatedAt    types.String `tfsdk:"created_at" json:"created_at,computed"`
	UpdatedAt    types.String `tfsdk:"updated_at" json:"updated_at,computed"`
}

type AccessServiceTokenDataSourceModel struct {
}

type AccessServiceTokensDataSourceModel struct {
}