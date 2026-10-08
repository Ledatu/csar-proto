# csar-proto

Shared Protocol Buffer definitions and generated Go stubs for the CSAR service family.

```
go get github.com/ledatu/csar-proto
```

Requires Go 1.25+.

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
