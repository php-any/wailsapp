package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ContextMenuManagerGetAllMethod struct {
	source *applicationsrc.ContextMenuManager
}

func (h *ContextMenuManagerGetAllMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.GetAll()
	return data.NewAnyValue(ret0), nil
}

func (h *ContextMenuManagerGetAllMethod) GetName() string            { return "getAll" }
func (h *ContextMenuManagerGetAllMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ContextMenuManagerGetAllMethod) GetIsStatic() bool          { return true }
func (h *ContextMenuManagerGetAllMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *ContextMenuManagerGetAllMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *ContextMenuManagerGetAllMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
