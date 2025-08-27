package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuAddSeparatorMethod struct {
	source *applicationsrc.Menu
}

func (h *MenuAddSeparatorMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.AddSeparator()
	return nil, nil
}

func (h *MenuAddSeparatorMethod) GetName() string            { return "addSeparator" }
func (h *MenuAddSeparatorMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuAddSeparatorMethod) GetIsStatic() bool          { return true }
func (h *MenuAddSeparatorMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *MenuAddSeparatorMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *MenuAddSeparatorMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
