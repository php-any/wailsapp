package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowCloseMethod struct {
	source applicationsrc.Window
}

func (h *WindowCloseMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Close()
	return nil, nil
}

func (h *WindowCloseMethod) GetName() string            { return "close" }
func (h *WindowCloseMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowCloseMethod) GetIsStatic() bool          { return true }
func (h *WindowCloseMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowCloseMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowCloseMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
