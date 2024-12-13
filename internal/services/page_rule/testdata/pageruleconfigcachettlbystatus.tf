resource "cloudflare_page_rule" "%[3]s" {
  zone_id = "%[1]s"
  target  = "%[2]s"
  actions = {
    cache_ttl_by_status = {
      "200" : "0",
      # "200-299" = 300
      # "300-399" = 60
      # "400-403" = -1
      # "404"     = 30
      # "405-499" = -1
      # "500-599" = 0
    }
  }
}
