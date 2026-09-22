This library serves as an internal library as a client for the StackState platform receiver.

Part of the receiver API is compatible with DataDog, those parts are extracted here, another part is new and comes from stackstate-openapi.

### OpenAPI connection options

`NewOpenAPIClientWithOptions` constructs an authenticated client with its own HTTP transport and a required request timeout. Supply the final Receiver URL: redirects are rejected, including same-host HTTP-to-HTTPS and trailing-slash redirects. Set `ConnectionOptions.ProxyURL` explicitly when a proxy is required; the transport does not read proxy environment variables. `CABundlePEM` adds trusted certificates to the system roots.

Provide exactly one authentication source: `APIKey` or `ServiceAccountToken`. The token callback is read on every request, including after rotation; empty credentials and header delimiters are rejected. The existing `NewOpenAPIClient` signature, `Connect()` method and legacy authentication behavior remain available unchanged.

### Bumping the openapi version

- Change the version/branch/commit sha in the `stackstate_openapi/openapi_version` file
- Run `nix develop -c ./scripts/generate_receiver_api.sh`
- Commit the generated code