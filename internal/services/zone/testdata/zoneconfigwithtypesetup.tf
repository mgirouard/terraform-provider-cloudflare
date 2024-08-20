resource "cloudflare_zone" "%[1]s" {
    name = "%[2]s"
    account = {
        id = "%[3]s"
    }
    plan = "%[4]s"
    type = "%[5]s"
}
