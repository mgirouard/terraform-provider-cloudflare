// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_ssl

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/custom_certificates"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type CustomSSLDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &CustomSSLDataSource{}

func NewCustomSSLDataSource() datasource.DataSource {
	return &CustomSSLDataSource{}
}

func (d *CustomSSLDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_ssl"
}

func (r *CustomSSLDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *CustomSSLDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *CustomSSLDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Filter == nil {
		res := new(http.Response)
		env := CustomSSLResultDataSourceEnvelope{*data}
		_, err := r.client.CustomCertificates.Get(
			ctx,
			data.CustomCertificateID.ValueString(),
			custom_certificates.CustomCertificateGetParams{
				ZoneID: cloudflare.F(data.ZoneID.ValueString()),
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
	} else {
		items := &[]*CustomSSLDataSourceModel{}
		env := CustomSSLResultListDataSourceEnvelope{items}

		page, err := r.client.CustomCertificates.List(ctx, custom_certificates.CustomCertificateListParams{
			ZoneID:  cloudflare.F(data.Filter.ZoneID.ValueString()),
			Match:   cloudflare.F(custom_certificates.CustomCertificateListParamsMatch(data.Filter.Match.ValueString())),
			Page:    cloudflare.F(data.Filter.Page.ValueFloat64()),
			PerPage: cloudflare.F(data.Filter.PerPage.ValueFloat64()),
			Status:  cloudflare.F(custom_certificates.CustomCertificateListParamsStatus(data.Filter.Status.ValueString())),
		})
		if err != nil {
			resp.Diagnostics.AddError("failed to make http request", err.Error())
			return
		}

		bytes := []byte(page.JSON.RawJSON())
		err = apijson.Unmarshal(bytes, &env)
		if err != nil {
			resp.Diagnostics.AddError("failed to unmarshal http request", err.Error())
			return
		}

		if count := len(*items); count != 1 {
			resp.Diagnostics.AddError("failed to find exactly one result", fmt.Sprint(count)+" found")
			return
		}
		data = (*items)[0]
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}