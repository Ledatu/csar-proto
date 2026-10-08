# csar-proto Agent Summary

## Role In Prod
`csar-proto` is the source of truth for CSAR gRPC API definitions and generated Go stubs. In prod today it backs the authz service/authn client path and carries the forward-compatible notify ingest contract. It is the only place where new gRPC contracts should live.

## Runtime Entry Points
- `csar/authz/v1/*.proto` defines the authz service and messages.
- `csar/notify/v1/*.proto` defines the notification ingest service and messages.
- Generated Go code under the matching package paths is consumed by downstream services.
- `buf generate` is the canonical regeneration path.

## Trust And Auth Model
- This repo does not run a service by itself; it defines wire contracts for shared gRPC APIs.
- Any authn/authz behavior changes here are security-sensitive because they change the on-wire API used by the router and downstream services.

## Critical Flows
- `CheckAccess` request/response contracts.
- Role, permission, and assignment management RPC shapes.
- Forward-compatible notify ingest request/response contracts.
- Generated client/server stubs that downstream modules compile against.
- `csar/audit/v1/AuditEvent.id` is additive field 14 carrying a stable event UUID;
  omitted by legacy producers. Consumers must upgrade before relying on replay
  deduplication; earlier consumers do not persist this field.

## Dependencies
- `buf` for linting and generation.
- `protoc` with Go and gRPC plugins if generation is done manually.
- Downstream consumers: `csar-authz` server, `csar-authn` client, notify producers/consumers as they adopt gRPC, and any future service that calls shared proto APIs.

## Config And Secrets
- No runtime config or secrets live here.
- Schema changes are the sensitive artifact; they must be coordinated with downstream consumers and router config expectations.

## Audit Hotspots
- A proto change is a blast-radius event: downstream modules may need regeneration and retesting.
- Do not duplicate authz proto definitions in service repos; keep the contract centralized here.

## First Files To Read
- `README.md`
- `csar/authz/v1/authz.proto`
- `csar/notify/v1/notify.proto`
- Any generated Go package under `csar/authz/v1`
- Any generated Go package under `csar/notify/v1`
- `buf.yaml` and `buf.gen.yaml` if regeneration behavior needs to change

## DRY / Extraction Candidates
- All new gRPC API definitions for the ecosystem belong here, not in service repos.
- Shared message shapes should be consolidated here before being copied into application code.

## Required Quality Gates
- `buf lint`
- `buf generate`
