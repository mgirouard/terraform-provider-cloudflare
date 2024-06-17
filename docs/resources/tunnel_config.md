---
page_title: "cloudflare_tunnel_config Resource - Cloudflare"
subcategory: ""
description: |-
  Provides a Cloudflare Tunnel configuration resource.
---

# cloudflare_tunnel_config (Resource)

Provides a Cloudflare Tunnel configuration resource.

## Example Usage

```terraform
resource "cloudflare_tunnel" "example_tunnel" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "example_tunnel"
  secret     = "<32 character secret>"
}

resource "cloudflare_tunnel_config" "example_config" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  tunnel_id  = cloudflare_tunnel.example_tunnel.id

  config {
    warp_routing {
      enabled = true
    }
    origin_request {
      connect_timeout          = "1m0s"
      tls_timeout              = "1m0s"
      tcp_keep_alive           = "1m0s"
      no_happy_eyeballs        = false
      keep_alive_connections   = 1024
      keep_alive_timeout       = "1m0s"
      http_host_header         = "baz"
      origin_server_name       = "foobar"
      ca_pool                  = "/path/to/unsigned/ca/pool"
      no_tls_verify            = false
      disable_chunked_encoding = false
      bastion_mode             = false
      proxy_address            = "10.0.0.1"
      proxy_port               = "8123"
      proxy_type               = "socks"
      ip_rules {
        prefix = "/web"
        ports  = [80, 443]
        allow  = false
      }
    }
    ingress_rule {
      hostname = "foo"
      path     = "/bar"
      service  = "http://10.0.0.2:8080"
      origin_request {
        connect_timeout = "2m0s"
        access {
          required  = true
          team_name = "terraform"
          aud_tag   = ["AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"]
        }
      }
    }
    ingress_rule {
      service = "https://10.0.0.3:8081"
    }
  }
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) The account identifier to target for the resource.
- `config` (Block List, Min: 1, Max: 1) Configuration block for Tunnel Configuration. (see [below for nested schema](#nestedblock--config))
- `tunnel_id` (String) Identifier of the Tunnel to target for this configuration.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--config"></a>
### Nested Schema for `config`

Required:

- `ingress_rule` (Block List, Min: 1) Each incoming request received by cloudflared causes cloudflared to send a request to a local service. This section configures the rules that determine which requests are sent to which local services. Last rule must match all requests, e.g `service = "http_status:503"`. [Read more](https://developers.cloudflare.com/cloudflare-one/connections/connect-apps/install-and-setup/tunnel-guide/local/local-management/ingress/). (see [below for nested schema](#nestedblock--config--ingress_rule))

Optional:

- `origin_request` (Block List, Max: 1) (see [below for nested schema](#nestedblock--config--origin_request))
- `warp_routing` (Block List, Max: 1) If you're exposing a [private network](https://developers.cloudflare.com/cloudflare-one/connections/connect-apps/private-net/), you need to add the `warp-routing` key and set it to `true`. (see [below for nested schema](#nestedblock--config--warp_routing))

<a id="nestedblock--config--ingress_rule"></a>
### Nested Schema for `config.ingress_rule`

Required:

- `service` (String) Name of the service to which the request will be sent.

Optional:

- `hostname` (String) Hostname to match the incoming request with. If the hostname matches, the request will be sent to the service.
- `origin_request` (Block List, Max: 1) (see [below for nested schema](#nestedblock--config--ingress_rule--origin_request))
- `path` (String) Path of the incoming request. If the path matches, the request will be sent to the local service.

<a id="nestedblock--config--ingress_rule--origin_request"></a>
### Nested Schema for `config.ingress_rule.origin_request`

Optional:

- `access` (Block List, Max: 1) Access rules for the ingress service. (see [below for nested schema](#nestedblock--config--ingress_rule--origin_request--access))
- `bastion_mode` (Boolean) Runs as jump host.
- `ca_pool` (String) Path to the certificate authority (CA) for the certificate of your origin. This option should be used only if your certificate is not signed by Cloudflare. Defaults to `""`.
- `connect_timeout` (String) Timeout for establishing a new TCP connection to your origin server. This excludes the time taken to establish TLS, which is controlled by `tlsTimeout`. Defaults to `30s`.
- `disable_chunked_encoding` (Boolean) Disables chunked transfer encoding. Useful if you are running a Web Server Gateway Interface (WSGI) server. Defaults to `false`.
- `http2_origin` (Boolean) Enables HTTP/2 support for the origin connection. Defaults to `false`.
- `http_host_header` (String) Sets the HTTP Host header on requests sent to the local service. Defaults to `""`.
- `ip_rules` (Block Set) IP rules for the proxy service. (see [below for nested schema](#nestedblock--config--ingress_rule--origin_request--ip_rules))
- `keep_alive_connections` (Number) Maximum number of idle keepalive connections between Tunnel and your origin. This does not restrict the total number of concurrent connections. Defaults to `100`.
- `keep_alive_timeout` (String) Timeout after which an idle keepalive connection can be discarded. Defaults to `1m30s`.
- `no_happy_eyeballs` (Boolean) Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local network has misconfigured one of the protocols. Defaults to `false`.
- `no_tls_verify` (Boolean) Disables TLS verification of the certificate presented by your origin. Will allow any certificate from the origin to be accepted. Defaults to `false`.
- `origin_server_name` (String) Hostname that cloudflared should expect from your origin server certificate. Defaults to `""`.
- `proxy_address` (String) cloudflared starts a proxy server to translate HTTP traffic into TCP when proxying, for example, SSH or RDP. This configures the listen address for that proxy. Defaults to `127.0.0.1`.
- `proxy_port` (Number) cloudflared starts a proxy server to translate HTTP traffic into TCP when proxying, for example, SSH or RDP. This configures the listen port for that proxy. If set to zero, an unused port will randomly be chosen. Defaults to `0`.
- `proxy_type` (String) cloudflared starts a proxy server to translate HTTP traffic into TCP when proxying, for example, SSH or RDP. This configures what type of proxy will be started. Available values: `""`, `socks`. Defaults to `""`.
- `tcp_keep_alive` (String) The timeout after which a TCP keepalive packet is sent on a connection between Tunnel and the origin server. Defaults to `30s`.
- `tls_timeout` (String) Timeout for completing a TLS handshake to your origin server, if you have chosen to connect Tunnel to an HTTPS server. Defaults to `10s`.

<a id="nestedblock--config--ingress_rule--origin_request--access"></a>
### Nested Schema for `config.ingress_rule.origin_request.access`

Optional:

- `aud_tag` (Set of String) Audience tags of the access rule.
- `required` (Boolean) Whether the access rule is required.
- `team_name` (String) Name of the team to which the access rule applies.


<a id="nestedblock--config--ingress_rule--origin_request--ip_rules"></a>
### Nested Schema for `config.ingress_rule.origin_request.ip_rules`

Optional:

- `allow` (Boolean) Whether to allow the IP prefix.
- `ports` (List of Number) Ports to use within the IP rule.
- `prefix` (String) IP rule prefix.




<a id="nestedblock--config--origin_request"></a>
### Nested Schema for `config.origin_request`

Optional:

- `access` (Block List, Max: 1) Access rules for the ingress service. (see [below for nested schema](#nestedblock--config--origin_request--access))
- `bastion_mode` (Boolean) Runs as jump host.
- `ca_pool` (String) Path to the certificate authority (CA) for the certificate of your origin. This option should be used only if your certificate is not signed by Cloudflare. Defaults to `""`.
- `connect_timeout` (String) Timeout for establishing a new TCP connection to your origin server. This excludes the time taken to establish TLS, which is controlled by `tlsTimeout`. Defaults to `30s`.
- `disable_chunked_encoding` (Boolean) Disables chunked transfer encoding. Useful if you are running a Web Server Gateway Interface (WSGI) server. Defaults to `false`.
- `http2_origin` (Boolean) Enables HTTP/2 support for the origin connection. Defaults to `false`.
- `http_host_header` (String) Sets the HTTP Host header on requests sent to the local service. Defaults to `""`.
- `ip_rules` (Block Set) IP rules for the proxy service. (see [below for nested schema](#nestedblock--config--origin_request--ip_rules))
- `keep_alive_connections` (Number) Maximum number of idle keepalive connections between Tunnel and your origin. This does not restrict the total number of concurrent connections. Defaults to `100`.
- `keep_alive_timeout` (String) Timeout after which an idle keepalive connection can be discarded. Defaults to `1m30s`.
- `no_happy_eyeballs` (Boolean) Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local network has misconfigured one of the protocols. Defaults to `false`.
- `no_tls_verify` (Boolean) Disables TLS verification of the certificate presented by your origin. Will allow any certificate from the origin to be accepted. Defaults to `false`.
- `origin_server_name` (String) Hostname that cloudflared should expect from your origin server certificate. Defaults to `""`.
- `proxy_address` (String) cloudflared starts a proxy server to translate HTTP traffic into TCP when proxying, for example, SSH or RDP. This configures the listen address for that proxy. Defaults to `127.0.0.1`.
- `proxy_port` (Number) cloudflared starts a proxy server to translate HTTP traffic into TCP when proxying, for example, SSH or RDP. This configures the listen port for that proxy. If set to zero, an unused port will randomly be chosen. Defaults to `0`.
- `proxy_type` (String) cloudflared starts a proxy server to translate HTTP traffic into TCP when proxying, for example, SSH or RDP. This configures what type of proxy will be started. Available values: `""`, `socks`. Defaults to `""`.
- `tcp_keep_alive` (String) The timeout after which a TCP keepalive packet is sent on a connection between Tunnel and the origin server. Defaults to `30s`.
- `tls_timeout` (String) Timeout for completing a TLS handshake to your origin server, if you have chosen to connect Tunnel to an HTTPS server. Defaults to `10s`.

<a id="nestedblock--config--origin_request--access"></a>
### Nested Schema for `config.origin_request.access`

Optional:

- `aud_tag` (Set of String) Audience tags of the access rule.
- `required` (Boolean) Whether the access rule is required.
- `team_name` (String) Name of the team to which the access rule applies.


<a id="nestedblock--config--origin_request--ip_rules"></a>
### Nested Schema for `config.origin_request.ip_rules`

Optional:

- `allow` (Boolean) Whether to allow the IP prefix.
- `ports` (List of Number) Ports to use within the IP rule.
- `prefix` (String) IP rule prefix.



<a id="nestedblock--config--warp_routing"></a>
### Nested Schema for `config.warp_routing`

Optional:

- `enabled` (Boolean) Whether WARP routing is enabled.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_tunnel_config.example <account_id>/<tunnel_id>
```