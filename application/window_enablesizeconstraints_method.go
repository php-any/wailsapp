package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowEnableSizeConstraintsMethod struct {
	source applicationsrc.Window
}

func (h *WindowEnableSizeConstraintsMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.EnableSizeConstraints()
	return nil, nil
}

func (h *WindowEnableSizeConstraintsMethod) GetName() string            { return "enableSizeConstraints" }
func (h *WindowEnableSizeConstraintsMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowEnableSizeConstraintsMethod) GetIsStatic() bool          { return true }
func (h *WindowEnableSizeConstraintsMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowEnableSizeConstraintsMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowEnableSizeConstraintsMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
