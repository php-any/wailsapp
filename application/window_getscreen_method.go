package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowGetScreenMethod struct {
	source applicationsrc.Window
}

func (h *WindowGetScreenMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0, err := h.source.GetScreen()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	return data.NewClassValue(NewScreenClassFrom(ret0), ctx), nil
}

func (h *WindowGetScreenMethod) GetName() string            { return "getScreen" }
func (h *WindowGetScreenMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowGetScreenMethod) GetIsStatic() bool          { return true }
func (h *WindowGetScreenMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowGetScreenMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowGetScreenMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
