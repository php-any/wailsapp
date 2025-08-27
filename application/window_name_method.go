package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowNameMethod struct {
	source applicationsrc.Window
}

func (h *WindowNameMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Name()
	return data.NewStringValue(ret0), nil
}

func (h *WindowNameMethod) GetName() string            { return "name" }
func (h *WindowNameMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowNameMethod) GetIsStatic() bool          { return true }
func (h *WindowNameMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowNameMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowNameMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
