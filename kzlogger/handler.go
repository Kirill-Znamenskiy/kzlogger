package kzlogger

import (
	"log/slog"

	"github.com/Kirill-Znamenskiy/kzlogger/lgl"
)

type LeveledHandler struct {
	slog.Handler
	Level *slog.LevelVar
}

func (h *LeveledHandler) Enabled(ctx Ctx, lvl lgl.Level) bool {
	return lvl >= h.Level.Level()
}
