This library serves as an internal library as a client for the StackState platform receiver.

Part of the receiver API is compatible with DataDog, those parts are extracted here, another part is new and comes from stackstate-openapi.

### OpenAPI connection and capability queries

Supply the final Receiver URL: redirects are rejected, including same-host HTTP-to-HTTPS and trailing-slash redirects. Set `ConnectionOptions.ProxyURL` explicitly when a proxy is required; this client does not read proxy environment variables. API-key and service-account authentication are mutually exclusive.

Feature queries bound attempts, elapsed time and response bodies (1 MiB). Set `QueryOptions.BooleanCapabilities` to the keys the consumer requires as booleans, such as `otel-logs` or `k8s-rbac`. Missing keys are valid, unrelated values are preserved, and an empty list checks only the response object. A 404 means the endpoint was not found; it cannot distinguish an older Receiver from an incorrect base path. Query results expose status and outcome without raw URLs, errors or bodies.

### Bumping the openapi version

- Change the version/branch/commit sha in the `stackstate_openapi/openapi_version` file
- Run `nix develop -c ./scripts/generate_receiver_api.sh`
- Commit the generated code
