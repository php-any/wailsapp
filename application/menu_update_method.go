package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuUpdateMethod struct {
	source *applicationsrc.Menu
}

func (h *MenuUpdateMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Update()
	return nil, nil
}

func (h *MenuUpdateMethod) GetName() string            { return "update" }
func (h *MenuUpdateMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuUpdateMethod) GetIsStatic() bool          { return true }
func (h *MenuUpdateMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *MenuUpdateMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *MenuUpdateMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
