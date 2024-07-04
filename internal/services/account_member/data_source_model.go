// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountMemberResultDataSourceEnvelope struct {
	Result AccountMemberDataSourceModel `json:"result,computed"`
}

type AccountMemberResultListDataSourceEnvelope struct {
	Result *[]*AccountMemberDataSourceModel `json:"result,computed"`
}

type AccountMemberDataSourceModel struct {
	AccountID types.String                             `tfsdk:"account_id" path:"account_id"`
	MemberID  types.String                             `tfsdk:"member_id" path:"member_id"`
	ID        types.String                             `tfsdk:"id" json:"id"`
	Policies  *[]*AccountMemberPoliciesDataSourceModel `tfsdk:"policies" json:"policies"`
	Roles     *[]*AccountMemberRolesDataSourceModel    `tfsdk:"roles" json:"roles"`
	Status    types.String                             `tfsdk:"status" json:"status"`
	User      *AccountMemberUserDataSourceModel        `tfsdk:"user" json:"user"`
	FindOneBy *AccountMemberFindOneByDataSourceModel   `tfsdk:"find_one_by"`
}

type AccountMemberPoliciesDataSourceModel struct {
	ID               types.String                                             `tfsdk:"id" json:"id,computed"`
	Access           types.String                                             `tfsdk:"access" json:"access"`
	PermissionGroups *[]*AccountMemberPoliciesPermissionGroupsDataSourceModel `tfsdk:"permission_groups" json:"permission_groups"`
	ResourceGroups   *[]*AccountMemberPoliciesResourceGroupsDataSourceModel   `tfsdk:"resource_groups" json:"resource_groups"`
}

type AccountMemberPoliciesPermissionGroupsDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Meta types.String `tfsdk:"meta" json:"meta"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type AccountMemberPoliciesResourceGroupsDataSourceModel struct {
	ID    types.String                                                `tfsdk:"id" json:"id,computed"`
	Scope *[]*AccountMemberPoliciesResourceGroupsScopeDataSourceModel `tfsdk:"scope" json:"scope"`
	Meta  types.String                                                `tfsdk:"meta" json:"meta"`
	Name  types.String                                                `tfsdk:"name" json:"name,computed"`
}

type AccountMemberPoliciesResourceGroupsScopeDataSourceModel struct {
	Key     types.String                                                       `tfsdk:"key" json:"key,computed"`
	Objects *[]*AccountMemberPoliciesResourceGroupsScopeObjectsDataSourceModel `tfsdk:"objects" json:"objects"`
}

type AccountMemberPoliciesResourceGroupsScopeObjectsDataSourceModel struct {
	Key types.String `tfsdk:"key" json:"key,computed"`
}

type AccountMemberRolesDataSourceModel struct {
	ID          types.String `tfsdk:"id" json:"id"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
	Permissions types.String `tfsdk:"permissions" json:"permissions,computed"`
}

type AccountMemberUserDataSourceModel struct {
	Email                          types.String `tfsdk:"email" json:"email"`
	ID                             types.String `tfsdk:"id" json:"id,computed"`
	FirstName                      types.String `tfsdk:"first_name" json:"first_name"`
	LastName                       types.String `tfsdk:"last_name" json:"last_name"`
	TwoFactorAuthenticationEnabled types.Bool   `tfsdk:"two_factor_authentication_enabled" json:"two_factor_authentication_enabled,computed"`
}

type AccountMemberFindOneByDataSourceModel struct {
	AccountID types.String  `tfsdk:"account_id" path:"account_id"`
	Direction types.String  `tfsdk:"direction" query:"direction"`
	Order     types.String  `tfsdk:"order" query:"order"`
	Page      types.Float64 `tfsdk:"page" query:"page"`
	PerPage   types.Float64 `tfsdk:"per_page" query:"per_page"`
	Status    types.String  `tfsdk:"status" query:"status"`
}