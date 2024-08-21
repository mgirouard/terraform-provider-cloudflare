// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = &ZeroTrustTunnelCloudflaredResource{}
var _ resource.ResourceWithModifyPlan = &ZeroTrustTunnelCloudflaredResource{}
var _ resource.ResourceWithImportState = &ZeroTrustTunnelCloudflaredResource{}

func NewResource() resource.Resource {
	return &ZeroTrustTunnelCloudflaredResource{}
}

// ZeroTrustTunnelCloudflaredResource defines the resource implementation.
type ZeroTrustTunnelCloudflaredResource struct {
	client *cloudflare.Client
}

func (r *ZeroTrustTunnelCloudflaredResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zero_trust_tunnel_cloudflared"
}

func (r *ZeroTrustTunnelCloudflaredResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *cloudflare.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ZeroTrustTunnelCloudflaredResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ZeroTrustTunnelCloudflaredModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := apijson.Marshal(data)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := ZeroTrustTunnelCloudflaredResultEnvelope{*data}
	_, err = r.client.ZeroTrust.Tunnels.New(
		ctx,
		zero_trust.TunnelNewParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustTunnelCloudflaredResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ZeroTrustTunnelCloudflaredModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *ZeroTrustTunnelCloudflaredModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := apijson.MarshalForUpdate(data, state)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := ZeroTrustTunnelCloudflaredResultEnvelope{*data}
	_, err = r.client.ZeroTrust.Tunnels.Edit(
		ctx,
		data.ID.ValueString(),
		zero_trust.TunnelEditParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustTunnelCloudflaredResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ZeroTrustTunnelCloudflaredModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := ZeroTrustTunnelCloudflaredResultEnvelope{*data}
	_, err := r.client.ZeroTrust.Tunnels.Get(
		ctx,
		data.ID.ValueString(),
		zero_trust.TunnelGetParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustTunnelCloudflaredResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ZeroTrustTunnelCloudflaredModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.ZeroTrust.Tunnels.Delete(
		ctx,
		data.ID.ValueString(),
		zero_trust.TunnelDeleteParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustTunnelCloudflaredResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *ZeroTrustTunnelCloudflaredModel

	path_account_id := ""
	path_tunnel_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<tunnel_id>",
		&path_account_id,
		&path_tunnel_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := ZeroTrustTunnelCloudflaredResultEnvelope{*data}
	_, err := r.client.ZeroTrust.Tunnels.Get(
		ctx,
		path_tunnel_id,
		zero_trust.TunnelGetParams{
			AccountID: cloudflare.F(path_account_id),
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustTunnelCloudflaredResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}