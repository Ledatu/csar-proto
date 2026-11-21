package auditv1_test

import (
	"testing"

	auditv1 "github.com/ledatu/csar-proto/csar/audit/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"
)

func TestOptionalIdentitySurvivesLegacyWireDecoder(t *testing.T) {
	input := &auditv1.AuditEvent{Id: "c30ca535-d731-4559-9238-d2b6906f1129", Actor: "legacy-actor",
		Service: "owner-service", Caller: &auditv1.CallerIdentity{Version: 1, CallerKind: "service",
			CallerId: "svc:reporting", EffectiveSubject: "svc:reporting", CredentialKind: "sts", Assurance: "verified"},
		OriginCaller: &auditv1.CallerIdentity{Version: 1, CallerKind: "agent", CallerId: "agent:1", OwnerSubject: "user:1",
			EffectiveSubject: "user:1", CredentialKind: "api_key", CredentialId: "key:1", Assurance: "verified"},
		OperationId: "operation:1", ParentOperationId: "operation:0", TraceId: "trace:1"}
	wire, err := proto.Marshal(input)
	if err != nil {
		t.Fatal(err)
	}
	file := protodesc.ToFileDescriptorProto(auditv1.File_csar_audit_v1_audit_proto)
	// Model the old protobuf decoder, not the old service's JSON/store adapters:
	// only fields 1-14 existed in AuditEvent before this additive change.
	file.MessageType[0].Field = file.MessageType[0].Field[:14]
	oldFile, err := protodesc.NewFile(file, protoregistry.GlobalFiles)
	if err != nil {
		t.Fatal(err)
	}
	legacy := dynamicpb.NewMessage(oldFile.Messages().ByName("AuditEvent"))
	if err := proto.Unmarshal(wire, legacy); err != nil {
		t.Fatal(err)
	}
	if legacy.Get(legacy.Descriptor().Fields().ByName("actor")).String() != input.Actor {
		t.Fatal("legacy actor changed")
	}
	replayed, err := proto.Marshal(legacy)
	if err != nil {
		t.Fatal(err)
	}
	var output auditv1.AuditEvent
	if err := proto.Unmarshal(replayed, &output); err != nil || !proto.Equal(input, &output) {
		t.Fatalf("optional fields lost on protobuf replay: %v", err)
	}
}

func TestLegacyEventDoesNotAcquireIdentity(t *testing.T) {
	wire, err := proto.Marshal(&auditv1.AuditEvent{Actor: "old-user", Id: "c30ca535-d731-4559-9238-d2b6906f1129"})
	if err != nil {
		t.Fatal(err)
	}
	var event auditv1.AuditEvent
	if err := proto.Unmarshal(wire, &event); err != nil {
		t.Fatal(err)
	}
	if event.Caller != nil || event.OriginCaller != nil || event.Actor != "old-user" {
		t.Fatal("legacy event identity was inferred")
	}
}
