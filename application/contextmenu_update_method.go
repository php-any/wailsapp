package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ContextMenuUpdateMethod struct {
	source *applicationsrc.ContextMenu
}

func (h *ContextMenuUpdateMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Update()
	return nil, nil
}

func (h *ContextMenuUpdateMethod) GetName() string            { return "update" }
func (h *ContextMenuUpdateMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ContextMenuUpdateMethod) GetIsStatic() bool          { return true }
func (h *ContextMenuUpdateMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *ContextMenuUpdateMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *ContextMenuUpdateMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
