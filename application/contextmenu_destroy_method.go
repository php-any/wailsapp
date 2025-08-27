package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ContextMenuDestroyMethod struct {
	source *applicationsrc.ContextMenu
}

func (h *ContextMenuDestroyMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Destroy()
	return nil, nil
}

func (h *ContextMenuDestroyMethod) GetName() string            { return "destroy" }
func (h *ContextMenuDestroyMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ContextMenuDestroyMethod) GetIsStatic() bool          { return true }
func (h *ContextMenuDestroyMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *ContextMenuDestroyMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *ContextMenuDestroyMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
