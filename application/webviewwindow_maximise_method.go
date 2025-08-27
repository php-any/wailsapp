package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowMaximiseMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowMaximiseMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Maximise()
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WebviewWindowMaximiseMethod) GetName() string            { return "maximise" }
func (h *WebviewWindowMaximiseMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowMaximiseMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowMaximiseMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowMaximiseMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowMaximiseMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
