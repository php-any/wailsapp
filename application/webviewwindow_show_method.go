package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowShowMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowShowMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Show()
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WebviewWindowShowMethod) GetName() string            { return "show" }
func (h *WebviewWindowShowMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowShowMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowShowMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowShowMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowShowMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
