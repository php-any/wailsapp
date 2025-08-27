package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowNameMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowNameMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Name()
	return data.NewStringValue(ret0), nil
}

func (h *WebviewWindowNameMethod) GetName() string            { return "name" }
func (h *WebviewWindowNameMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowNameMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowNameMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowNameMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowNameMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
