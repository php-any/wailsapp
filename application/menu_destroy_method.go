package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuDestroyMethod struct {
	source *applicationsrc.Menu
}

func (h *MenuDestroyMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Destroy()
	return nil, nil
}

func (h *MenuDestroyMethod) GetName() string            { return "destroy" }
func (h *MenuDestroyMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuDestroyMethod) GetIsStatic() bool          { return true }
func (h *MenuDestroyMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *MenuDestroyMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *MenuDestroyMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
