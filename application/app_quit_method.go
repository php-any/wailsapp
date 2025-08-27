package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type AppQuitMethod struct {
	source *applicationsrc.App
}

func (h *AppQuitMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Quit()
	return nil, nil
}

func (h *AppQuitMethod) GetName() string            { return "quit" }
func (h *AppQuitMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *AppQuitMethod) GetIsStatic() bool          { return true }
func (h *AppQuitMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *AppQuitMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *AppQuitMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
