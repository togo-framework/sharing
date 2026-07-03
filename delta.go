package sharing

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/togo-framework/providers"
	"github.com/togo-framework/togo"
)

// delta is a client for the Delta Sharing / OpenSharing REST protocol.
type delta struct {
	base  string
	token string
	hc    *http.Client
}

func newDelta(k *togo.Kernel) *delta {
	return &delta{
		base:  strings.TrimRight(providers.Value(k, providers.CapShare, "delta", "endpoint", "", false), "/"),
		token: providers.Value(k, providers.CapShare, "delta", "token", "", true),
		hc:    &http.Client{Timeout: 30 * time.Second},
	}
}

func (d *delta) get(ctx context.Context, path string, dst any) error {
	if d.base == "" {
		return fmt.Errorf("sharing: set SHARING_ENDPOINT")
	}
	req, _ := http.NewRequestWithContext(ctx, "GET", d.base+path, nil)
	if d.token != "" {
		req.Header.Set("Authorization", "Bearer "+d.token)
	}
	resp, err := d.hc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return fmt.Errorf("sharing %s: %s", path, resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(dst)
}

func (d *delta) Shares(ctx context.Context) ([]Share, error) {
	var out struct {
		Items []Share `json:"items"`
	}
	if err := d.get(ctx, "/shares", &out); err != nil {
		return nil, err
	}
	return out.Items, nil
}

func (d *delta) Tables(ctx context.Context, share, schema string) ([]SharedTable, error) {
	var out struct {
		Items []SharedTable `json:"items"`
	}
	if err := d.get(ctx, "/shares/"+share+"/schemas/"+schema+"/tables", &out); err != nil {
		return nil, err
	}
	return out.Items, nil
}
