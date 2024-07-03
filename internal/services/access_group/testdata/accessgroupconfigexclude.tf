
resource "cloudflare_access_group" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s"

  include {
    email = ["%[3]s"]
	email_domain = ["example.com"]
  }

  exclude {
    email = ["%[3]s"]
  }
}