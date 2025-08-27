package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowIsMinimisedMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowIsMinimisedMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.IsMinimised()
	return data.NewBoolValue(ret0), nil
}

func (h *WebviewWindowIsMinimisedMethod) GetName() string            { return "isMinimised" }
func (h *WebviewWindowIsMinimisedMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowIsMinimisedMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowIsMinimisedMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowIsMinimisedMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowIsMinimisedMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
