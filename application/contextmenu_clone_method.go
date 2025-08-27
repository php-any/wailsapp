package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ContextMenuCloneMethod struct {
	source *applicationsrc.ContextMenu
}

func (h *ContextMenuCloneMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Clone()
	return data.NewClassValue(NewMenuClassFrom(ret0), ctx), nil
}

func (h *ContextMenuCloneMethod) GetName() string            { return "clone" }
func (h *ContextMenuCloneMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ContextMenuCloneMethod) GetIsStatic() bool          { return true }
func (h *ContextMenuCloneMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *ContextMenuCloneMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *ContextMenuCloneMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
