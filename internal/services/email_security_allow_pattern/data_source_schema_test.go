// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_allow_pattern_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_security_allow_pattern"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestEmailSecurityAllowPatternDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*email_security_allow_pattern.EmailSecurityAllowPatternDataSourceModel)(nil)
	schema := email_security_allow_pattern.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}