# csar-proto

Shared Protocol Buffer definitions and generated Go stubs for the CSAR service family.

```
go get github.com/ledatu/csar-proto
```

Requires Go 1.26+. CI and release builds use the patched Go 1.26.11 toolchain.

---

## Services

### `csar/authz/v1` -- Authorization Service

RBAC-based authorization decisions and policy management. Used by csar-authz (server) and csar-authn (client).

```go
import pb "github.com/ledatu/csar-proto/csar/authz/v1"

client := pb.NewAuthzServiceClient(conn)
resp, err := client.CheckAccess(ctx, &pb.CheckAccessRequest{
    Subject:  "user-uuid",
    Resource: "/api/v1/documents/123",
    Action:   "DELETE",
})
```

### `csar/notify/v1` -- Notification Ingest Service

Forward-compatible notification ingest contract. Phase 1 producers still use
HTTP ingest through the csar router, but the proto package and generated Go
stubs are kept here for future gRPC producers.

```go
import notifyv1 "github.com/ledatu/csar-proto/csar/notify/v1"

_ = notifyv1.SendNotificationRequest{}
```

---

## Regenerating Stubs

Install [buf](https://buf.build/docs/installation/):

```bash
buf generate
```

Or with protoc directly:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       csar/authz/v1/authz.proto csar/notify/v1/notify.proto
```

---

## License

MIT

## Stored credential rejection reports

`csar/credential/v1` defines the additive `CredentialService.ReportRejection`
contract for trusted router reports of actual upstream authentication 401s. It
carries only the injected ref, opaque version and credential fingerprint, never
the token. The coordinator must authenticate the router transport, compare the
reported value against current rotation state and preserve fail-closed behavior.
It does not replay a request or mint a replacement. Existing AuthService contracts
remain compatible; deploy the report server before relying on durable rejection.

`csar/throttle/v1.RequestQuota` defines rolling request caps for coordinator
traffic policies. CSAR owns enforcement and capability negotiation; workers use
its retry headers. Limits, scope keys and windows survive configuration transport.
