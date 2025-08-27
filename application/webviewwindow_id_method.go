package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowIDMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowIDMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.ID()
	return data.NewAnyValue(ret0), nil
}

func (h *WebviewWindowIDMethod) GetName() string            { return "iD" }
func (h *WebviewWindowIDMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowIDMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowIDMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowIDMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowIDMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
