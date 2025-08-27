package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowReloadMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowReloadMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Reload()
	return nil, nil
}

func (h *WebviewWindowReloadMethod) GetName() string            { return "reload" }
func (h *WebviewWindowReloadMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowReloadMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowReloadMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowReloadMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowReloadMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
