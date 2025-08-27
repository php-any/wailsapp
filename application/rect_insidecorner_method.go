package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type RectInsideCornerMethod struct {
	source *applicationsrc.Rect
}

func (h *RectInsideCornerMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.InsideCorner()
	return data.NewClassValue(NewPointClassFrom(&ret0), ctx), nil
}

func (h *RectInsideCornerMethod) GetName() string            { return "insideCorner" }
func (h *RectInsideCornerMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *RectInsideCornerMethod) GetIsStatic() bool          { return true }
func (h *RectInsideCornerMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *RectInsideCornerMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *RectInsideCornerMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
