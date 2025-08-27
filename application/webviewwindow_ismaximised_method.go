package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowIsMaximisedMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowIsMaximisedMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.IsMaximised()
	return data.NewBoolValue(ret0), nil
}

func (h *WebviewWindowIsMaximisedMethod) GetName() string            { return "isMaximised" }
func (h *WebviewWindowIsMaximisedMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowIsMaximisedMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowIsMaximisedMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowIsMaximisedMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowIsMaximisedMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
