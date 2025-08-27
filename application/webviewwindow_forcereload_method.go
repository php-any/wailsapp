package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowForceReloadMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowForceReloadMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.ForceReload()
	return nil, nil
}

func (h *WebviewWindowForceReloadMethod) GetName() string            { return "forceReload" }
func (h *WebviewWindowForceReloadMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowForceReloadMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowForceReloadMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowForceReloadMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowForceReloadMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
