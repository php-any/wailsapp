package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuManagerGetApplicationMenuMethod struct {
	source *applicationsrc.MenuManager
}

func (h *MenuManagerGetApplicationMenuMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.GetApplicationMenu()
	return data.NewClassValue(NewMenuClassFrom(ret0), ctx), nil
}

func (h *MenuManagerGetApplicationMenuMethod) GetName() string            { return "getApplicationMenu" }
func (h *MenuManagerGetApplicationMenuMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuManagerGetApplicationMenuMethod) GetIsStatic() bool          { return true }
func (h *MenuManagerGetApplicationMenuMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *MenuManagerGetApplicationMenuMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *MenuManagerGetApplicationMenuMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
