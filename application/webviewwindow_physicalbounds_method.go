package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowPhysicalBoundsMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowPhysicalBoundsMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.PhysicalBounds()
	return data.NewClassValue(NewRectClassFrom(&ret0), ctx), nil
}

func (h *WebviewWindowPhysicalBoundsMethod) GetName() string            { return "physicalBounds" }
func (h *WebviewWindowPhysicalBoundsMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowPhysicalBoundsMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowPhysicalBoundsMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowPhysicalBoundsMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowPhysicalBoundsMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
