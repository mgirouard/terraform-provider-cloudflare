---
page_title: "cloudflare_hyperdrive_config Resource - Cloudflare"
subcategory: ""
description: |-
  The Hyperdrive Config https://developers.cloudflare.com/hyperdrive/ resource allows you to manage Cloudflare Hyperdrive Configs.
---

# cloudflare_hyperdrive_config (Resource)

The [Hyperdrive Config](https://developers.cloudflare.com/hyperdrive/) resource allows you to manage Cloudflare Hyperdrive Configs.

## Example Usage

```terraform
resource "cloudflare_hyperdrive_config" "no_defaults" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "my-hyperdrive-config"
  origin = {
    database = "postgres"
    password = "my-password"
    host     = "my-database.example.com"
    port     = 5432
    scheme   = "postgres"
    user     = "my-user"
  }
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) The account identifier to target for the resource.
- `name` (String) The name of the Hyperdrive configuration.
- `origin` (Attributes) The origin details for the Hyperdrive configuration. (see [below for nested schema](#nestedatt--origin))

### Optional

- `caching` (Attributes) The caching details for the Hyperdrive configuration. (see [below for nested schema](#nestedatt--caching))
- `id` (String) The identifier of this resource. This is the hyperdrive config value.

<a id="nestedatt--origin"></a>
### Nested Schema for `origin`

Required:

- `database` (String) The name of your origin database.
- `host` (String) The host (hostname or IP) of your origin database.
- `password` (String, Sensitive) The password of the Hyperdrive configuration.
- `port` (Number) The port (default: 5432 for Postgres) of your origin database.
- `scheme` (String) Specifies the URL scheme used to connect to your origin database.
- `user` (String) The user of your origin database.


<a id="nestedatt--caching"></a>
### Nested Schema for `caching`

Optional:

- `disabled` (Boolean) Disable caching for this Hyperdrive configuration.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_hyperdrive_config.example <account_id>/<hyperdrive_config_id>
```