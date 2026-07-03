package sharing

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/togo-framework/togo"
)

func TestSharesAndTables(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/shares":
			_, _ = w.Write([]byte(`{"items":[{"name":"sales"}]}`))
		case strings.HasSuffix(r.URL.Path, "/tables"):
			_, _ = w.Write([]byte(`{"items":[{"share":"sales","schema":"public","name":"orders"}]}`))
		}
	}))
	defer srv.Close()
	d := &delta{base: srv.URL, hc: srv.Client()}
	shares, err := d.Shares(context.Background())
	if err != nil || len(shares) != 1 || shares[0].Name != "sales" {
		t.Fatalf("shares = %v err=%v", shares, err)
	}
	tbls, err := d.Tables(context.Background(), "sales", "public")
	if err != nil || len(tbls) != 1 || tbls[0].Name != "orders" {
		t.Fatalf("tables = %v err=%v", tbls, err)
	}
}

func TestRequiresEndpoint(t *testing.T) {
	if _, err := newDelta(togo.New()).Shares(context.Background()); err == nil {
		t.Fatal("expected error without SHARING_ENDPOINT")
	}
}
