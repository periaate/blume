# fwauth
fwauth is a forward auth implementation of [auth](https://github.com/periaate/auth) with persistence implemented with [blob](https://github.com/periaate/blob).

## Usage
The fwauth server starts two endpoints, one inward enpoint, used to generate links to authenticate users, the other to validate said sessions. For example, starting `fwauth` with the following command:
```
fwauth :8900 :8950
```
Will start a local link generation service at `127.0.0.1:8900`, and the validation service at `127.0.0.1:8950`. The services have the following routes:
```
:8900 GET /gen/{host}/{linkDuration}/{linkUses}/{sessionDuration}
:8950 GET /fw-auth/{host}/{hash...}
```
We will create a link for `example.tld` which can be used for the next 5 minutes up to 2 times. The generated session will last 30 minutes.
```
127.0.0.1:8900/gen/example.tld/5/2/30
```
Which returns something akin to:
```
https://example.tld/Olr3_bXyQljdPLKJiy8ePqFgCphJx2PrSBYgzsI0peg=
```

### Caddy example
The fwauth was built specifically with Caddy in mind. The following example implements a Caddy instance secured by fwauth.
```Caddyfile
:8901 {
	forward_auth {env.FW_AUTH_ADDR} {
		uri /fw-auth/{http.request.host}{http.request.uri}
	}

	respond "Hello, World! You have been authenticated!"
}
```
