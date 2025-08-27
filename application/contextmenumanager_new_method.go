package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ContextMenuManagerNewMethod struct {
	source *applicationsrc.ContextMenuManager
}

func (h *ContextMenuManagerNewMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.New()
	return data.NewClassValue(NewContextMenuClassFrom(ret0), ctx), nil
}

func (h *ContextMenuManagerNewMethod) GetName() string            { return "new" }
func (h *ContextMenuManagerNewMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ContextMenuManagerNewMethod) GetIsStatic() bool          { return true }
func (h *ContextMenuManagerNewMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *ContextMenuManagerNewMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *ContextMenuManagerNewMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
