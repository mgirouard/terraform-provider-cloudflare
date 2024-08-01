// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneResultDataSourceEnvelope struct {
	Result ZoneDataSourceModel `json:"result,computed"`
}

type ZoneResultListDataSourceEnvelope struct {
	Result *[]*ZoneDataSourceModel `json:"result,computed"`
}

type ZoneDataSourceModel struct {
	ZoneID              types.String                                         `tfsdk:"zone_id" path:"zone_id"`
	ID                  types.String                                         `tfsdk:"id" json:"id,computed"`
	Account             customfield.NestedObject[ZoneAccountDataSourceModel] `tfsdk:"account" json:"account,computed"`
	ActivatedOn         timetypes.RFC3339                                    `tfsdk:"activated_on" json:"activated_on,computed"`
	CreatedOn           timetypes.RFC3339                                    `tfsdk:"created_on" json:"created_on,computed"`
	DevelopmentMode     types.Float64                                        `tfsdk:"development_mode" json:"development_mode,computed"`
	Meta                customfield.NestedObject[ZoneMetaDataSourceModel]    `tfsdk:"meta" json:"meta,computed"`
	ModifiedOn          timetypes.RFC3339                                    `tfsdk:"modified_on" json:"modified_on,computed"`
	Name                types.String                                         `tfsdk:"name" json:"name,computed"`
	NameServers         *[]types.String                                      `tfsdk:"name_servers" json:"name_servers,computed"`
	OriginalDnshost     types.String                                         `tfsdk:"original_dnshost" json:"original_dnshost,computed"`
	OriginalNameServers *[]types.String                                      `tfsdk:"original_name_servers" json:"original_name_servers,computed"`
	OriginalRegistrar   types.String                                         `tfsdk:"original_registrar" json:"original_registrar,computed"`
	Owner               customfield.NestedObject[ZoneOwnerDataSourceModel]   `tfsdk:"owner" json:"owner,computed"`
	VanityNameServers   *[]types.String                                      `tfsdk:"vanity_name_servers" json:"vanity_name_servers"`
	Filter              *ZoneFindOneByDataSourceModel                        `tfsdk:"filter"`
}

type ZoneAccountDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id"`
	Name types.String `tfsdk:"name" json:"name"`
}

type ZoneMetaDataSourceModel struct {
	CDNOnly                types.Bool  `tfsdk:"cdn_only" json:"cdn_only"`
	CustomCertificateQuota types.Int64 `tfsdk:"custom_certificate_quota" json:"custom_certificate_quota"`
	DNSOnly                types.Bool  `tfsdk:"dns_only" json:"dns_only"`
	FoundationDNS          types.Bool  `tfsdk:"foundation_dns" json:"foundation_dns"`
	PageRuleQuota          types.Int64 `tfsdk:"page_rule_quota" json:"page_rule_quota"`
	PhishingDetected       types.Bool  `tfsdk:"phishing_detected" json:"phishing_detected"`
	Step                   types.Int64 `tfsdk:"step" json:"step"`
}

type ZoneOwnerDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id"`
	Name types.String `tfsdk:"name" json:"name"`
	Type types.String `tfsdk:"type" json:"type"`
}

type ZoneFindOneByDataSourceModel struct {
	Account   *ZoneAccountDataSourceModel `tfsdk:"account" query:"account"`
	Direction types.String                `tfsdk:"direction" query:"direction"`
	Match     types.String                `tfsdk:"match" query:"match"`
	Name      types.String                `tfsdk:"name" query:"name"`
	Order     types.String                `tfsdk:"order" query:"order"`
	Page      types.Float64               `tfsdk:"page" query:"page"`
	PerPage   types.Float64               `tfsdk:"per_page" query:"per_page"`
	Status    types.String                `tfsdk:"status" query:"status"`
}