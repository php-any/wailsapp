package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowIDMethod struct {
	source applicationsrc.Window
}

func (h *WindowIDMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.ID()
	return data.NewAnyValue(ret0), nil
}

func (h *WindowIDMethod) GetName() string            { return "iD" }
func (h *WindowIDMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowIDMethod) GetIsStatic() bool          { return true }
func (h *WindowIDMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowIDMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowIDMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
