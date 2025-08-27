package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ContextMenuAddSeparatorMethod struct {
	source *applicationsrc.ContextMenu
}

func (h *ContextMenuAddSeparatorMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.AddSeparator()
	return nil, nil
}

func (h *ContextMenuAddSeparatorMethod) GetName() string            { return "addSeparator" }
func (h *ContextMenuAddSeparatorMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ContextMenuAddSeparatorMethod) GetIsStatic() bool          { return true }
func (h *ContextMenuAddSeparatorMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *ContextMenuAddSeparatorMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *ContextMenuAddSeparatorMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
