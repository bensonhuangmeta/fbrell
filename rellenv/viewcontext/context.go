// Package viewcontext provides context related HTTP Handlers for Rell.
package viewcontext

import (
	"net/http"

	"github.com/daaku/rell/internal/github.com/daaku/go.httpdev"
	"github.com/daaku/rell/internal/golang.org/x/net/context"
	"github.com/daaku/rell/rellenv"
)

var version string

type Handler struct{}

// Handler for /info/ to see a JSON view of some server context.
func (h *Handler) Info(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	env, err := rellenv.FromContext(ctx)
	if err != nil {
		return err
	}
	info := map[string]interface{}{
		"context":    env,
		"pageTabURL": env.PageTabURL("/"),
		"canvasURL":  env.CanvasURL("/"),
		"sdkURL":     env.SdkURL(),
		"version":    version,
	}
	httpdev.Info(info, w, r)
	return nil
}
