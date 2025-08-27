package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type RectIsEmptyMethod struct {
	source *applicationsrc.Rect
}

func (h *RectIsEmptyMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.IsEmpty()
	return data.NewBoolValue(ret0), nil
}

func (h *RectIsEmptyMethod) GetName() string            { return "isEmpty" }
func (h *RectIsEmptyMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *RectIsEmptyMethod) GetIsStatic() bool          { return true }
func (h *RectIsEmptyMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *RectIsEmptyMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *RectIsEmptyMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
