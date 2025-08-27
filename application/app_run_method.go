package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type AppRunMethod struct {
	source *applicationsrc.App
}

func (h *AppRunMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	if err := h.source.Run(); err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	return nil, nil
}

func (h *AppRunMethod) GetName() string            { return "run" }
func (h *AppRunMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *AppRunMethod) GetIsStatic() bool          { return true }
func (h *AppRunMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *AppRunMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *AppRunMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
