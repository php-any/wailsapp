package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowEventContextDroppedFilesMethod struct {
	source *applicationsrc.WindowEventContext
}

func (h *WindowEventContextDroppedFilesMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.DroppedFiles()
	return data.NewAnyValue(ret0), nil
}

func (h *WindowEventContextDroppedFilesMethod) GetName() string { return "droppedFiles" }
func (h *WindowEventContextDroppedFilesMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *WindowEventContextDroppedFilesMethod) GetIsStatic() bool { return true }
func (h *WindowEventContextDroppedFilesMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowEventContextDroppedFilesMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowEventContextDroppedFilesMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
