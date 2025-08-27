package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowHeightMethod struct {
	source applicationsrc.Window
}

func (h *WindowHeightMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Height()
	return data.NewIntValue(ret0), nil
}

func (h *WindowHeightMethod) GetName() string            { return "height" }
func (h *WindowHeightMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowHeightMethod) GetIsStatic() bool          { return true }
func (h *WindowHeightMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowHeightMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowHeightMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
