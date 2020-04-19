package socks5

import (
	"context"
	"testing"
)

func TestPermitCommand(t *testing.T) {
	ctx := context.Background()
	r := &PermitCommand{true, false, false}

	if _, ok := r.Allow(ctx, &Request{Header: Header{Command: ConnectCommand}}); !ok {
		t.Fatalf("expect connect")
	}

	if _, ok := r.Allow(ctx, &Request{Header: Header{Command: BindCommand}}); ok {
		t.Fatalf("do not expect bind")
	}

	if _, ok := r.Allow(ctx, &Request{Header: Header{Command: AssociateCommand}}); ok {
		t.Fatalf("do not expect associate")
	}
}
