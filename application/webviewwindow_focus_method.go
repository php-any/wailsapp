package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowFocusMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowFocusMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Focus()
	return nil, nil
}

func (h *WebviewWindowFocusMethod) GetName() string            { return "focus" }
func (h *WebviewWindowFocusMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowFocusMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowFocusMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowFocusMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowFocusMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
