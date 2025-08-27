package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowEventContextDropZoneDetailsMethod struct {
	source *applicationsrc.WindowEventContext
}

func (h *WindowEventContextDropZoneDetailsMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.DropZoneDetails()
	return data.NewClassValue(NewDropZoneDetailsClassFrom(ret0), ctx), nil
}

func (h *WindowEventContextDropZoneDetailsMethod) GetName() string { return "dropZoneDetails" }
func (h *WindowEventContextDropZoneDetailsMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *WindowEventContextDropZoneDetailsMethod) GetIsStatic() bool { return true }
func (h *WindowEventContextDropZoneDetailsMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowEventContextDropZoneDetailsMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowEventContextDropZoneDetailsMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
