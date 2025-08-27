package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuManagerNewMethod struct {
	source *applicationsrc.MenuManager
}

func (h *MenuManagerNewMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.New()
	return data.NewClassValue(NewMenuClassFrom(ret0), ctx), nil
}

func (h *MenuManagerNewMethod) GetName() string            { return "new" }
func (h *MenuManagerNewMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuManagerNewMethod) GetIsStatic() bool          { return true }
func (h *MenuManagerNewMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *MenuManagerNewMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *MenuManagerNewMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
