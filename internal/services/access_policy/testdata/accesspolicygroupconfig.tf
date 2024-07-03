
    resource "cloudflare_access_application" "%[1]s" {
      name       = "%[1]s"
      account_id = "%[3]s"
      domain     = "%[1]s.%[2]s"
    }

    resource "cloudflare_access_group" "%[1]s" {
        account_id     = "%[3]s"
        name           = "%[1]s"

        include {
          ip = ["127.0.0.1/32"]
        }
    }

    resource "cloudflare_access_policy" "%[1]s" {
      application_id = "${cloudflare_access_application.%[1]s.id}"
      name           = "%[1]s"
      account_id     = "%[3]s"
      decision       = "non_identity"
      precedence     = "10"

      include {
        group = [cloudflare_access_group.%[1]s.id]
      }
    }

  