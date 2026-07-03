// Package sharing is togo's zero-copy data-sharing capability — a client for the
// open OpenSharing / Delta Sharing REST protocol. Consume datasets (and, on
// OpenSharing servers, agent skills) shared by partners without copying data.
// Select/point it with `togo provider:use share delta`.
//
//   SHARING_ENDPOINT  Delta Sharing endpoint base URL
//   SHARING_TOKEN     bearer token (secret)
package sharing

import (
	"context"

	"github.com/togo-framework/providers"
	"github.com/togo-framework/togo"
)

type Share struct {
	Name string `json:"name"`
}

type SharedTable struct {
	Share  string `json:"share"`
	Schema string `json:"schema"`
	Name   string `json:"name"`
}

// Sharing lists shares/tables offered by an open-sharing server.
type Sharing interface {
	Shares(ctx context.Context) ([]Share, error)
	Tables(ctx context.Context, share, schema string) ([]SharedTable, error)
}

func init() {
	togo.RegisterProviderFunc("sharing", togo.PriorityService, func(k *togo.Kernel) error {
		providers.Use(k, providers.CapShare, "delta", newDelta(k), true)
		if k.Log != nil {
			k.Log.Info("plugin active", "plugin", "sharing")
		}
		return nil
	})
}

// FromKernel returns the active Sharing client (or nil).
func FromKernel(k *togo.Kernel) Sharing {
	if v, ok := k.Get(providers.CapShare); ok {
		if s, ok := v.(Sharing); ok {
			return s
		}
	}
	return nil
}
