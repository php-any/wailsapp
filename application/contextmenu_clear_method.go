package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ContextMenuClearMethod struct {
	source *applicationsrc.ContextMenu
}

func (h *ContextMenuClearMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Clear()
	return nil, nil
}

func (h *ContextMenuClearMethod) GetName() string            { return "clear" }
func (h *ContextMenuClearMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ContextMenuClearMethod) GetIsStatic() bool          { return true }
func (h *ContextMenuClearMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *ContextMenuClearMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *ContextMenuClearMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
