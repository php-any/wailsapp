package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuClearMethod struct {
	source *applicationsrc.Menu
}

func (h *MenuClearMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Clear()
	return nil, nil
}

func (h *MenuClearMethod) GetName() string            { return "clear" }
func (h *MenuClearMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuClearMethod) GetIsStatic() bool          { return true }
func (h *MenuClearMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *MenuClearMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *MenuClearMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
