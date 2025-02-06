This library serves as an internal library as a client for the StackState platform receiver.

Part of the receiver API is compatible with DataDog, those parts are extracted here, another part is new and comes from stackstate-openapi.

### Bumping the openapi version

- Change the version/branch/commit sha in the `stackstate_openapi/openapi_version` file
- Run `nix develop -c ./scripts/generate_receiver_api.sh`
- Commit the generated code