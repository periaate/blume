The {{ .Module }} was built specifically with Caddy in mind. The following example implements a Caddy instance secured by {{ .Module }}.
```Caddyfile
:8901 {
	forward_auth {env.FW_AUTH_ADDR} {
		uri /fw-auth/{http.request.host}{http.request.uri}
	}

	respond "Hello, World! You have been authenticated!"
}
```
