package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowRunMethod struct {
	source applicationsrc.Window
}

func (h *WindowRunMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Run()
	return nil, nil
}

func (h *WindowRunMethod) GetName() string            { return "run" }
func (h *WindowRunMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowRunMethod) GetIsStatic() bool          { return true }
func (h *WindowRunMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowRunMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowRunMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
