package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuCloneMethod struct {
	source *applicationsrc.Menu
}

func (h *MenuCloneMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Clone()
	return data.NewClassValue(NewMenuClassFrom(ret0), ctx), nil
}

func (h *MenuCloneMethod) GetName() string            { return "clone" }
func (h *MenuCloneMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuCloneMethod) GetIsStatic() bool          { return true }
func (h *MenuCloneMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *MenuCloneMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *MenuCloneMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
