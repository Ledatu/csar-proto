# csar-proto

Shared Protocol Buffer definitions and generated Go stubs for the CSAR service family.

```
go get github.com/ledatu/csar-proto
```

Requires Go 1.25+.

---

## Services

### `authz/v1` -- Authorization Service

RBAC-based authorization decisions and policy management. Used by csar-authz (server) and csar-authn (client).

```go
import pb "github.com/ledatu/csar-proto/authz/v1"

client := pb.NewAuthzServiceClient(conn)
resp, err := client.CheckAccess(ctx, &pb.CheckAccessRequest{
    Subject:  "user-uuid",
    Resource: "/api/v1/documents/123",
    Action:   "DELETE",
})
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
       authz/v1/authz.proto
```

---

## License

MIT
