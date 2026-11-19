# csar-proto summary

This shared module owns protobuf wire contracts and generated Go stubs. It has
no runtime or credential storage. Existing packages cover authorization, audit
and notification; `csar/credential/v1` adds trusted-router rejection reporting.

Credential reports contain ref, opaque version and a value fingerprint, never
the token. Servers must authenticate the router, preserve rotation ordering and
fail closed. No report authorizes replay or token generation. Upgrade readers
before relying on an additive protocol field or service.

Use `buf lint` and `buf generate` after protocol edits. Local generation may use
the installed protoc Go plugins with source-relative paths. Run Go build/test
and the baseline golangci-lint configuration for generated Go changes. Downstream
consumers must pin the published revision rather than depend on a workspace-only
replacement.
