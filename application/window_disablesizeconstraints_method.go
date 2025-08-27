package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowDisableSizeConstraintsMethod struct {
	source applicationsrc.Window
}

func (h *WindowDisableSizeConstraintsMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.DisableSizeConstraints()
	return nil, nil
}

func (h *WindowDisableSizeConstraintsMethod) GetName() string            { return "disableSizeConstraints" }
func (h *WindowDisableSizeConstraintsMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowDisableSizeConstraintsMethod) GetIsStatic() bool          { return true }
func (h *WindowDisableSizeConstraintsMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowDisableSizeConstraintsMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowDisableSizeConstraintsMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
