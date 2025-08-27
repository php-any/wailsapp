package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowMinimiseMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowMinimiseMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Minimise()
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WebviewWindowMinimiseMethod) GetName() string            { return "minimise" }
func (h *WebviewWindowMinimiseMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowMinimiseMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowMinimiseMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowMinimiseMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowMinimiseMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
