package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/std/context"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type AppContextMethod struct {
	source *applicationsrc.App
}

func (h *AppContextMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Context()
	return data.NewClassValue(context.NewContextClassFrom(ret0), ctx), nil
}

func (h *AppContextMethod) GetName() string            { return "context" }
func (h *AppContextMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *AppContextMethod) GetIsStatic() bool          { return true }
func (h *AppContextMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *AppContextMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *AppContextMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
