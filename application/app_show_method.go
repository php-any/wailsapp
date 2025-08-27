package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type AppShowMethod struct {
	source *applicationsrc.App
}

func (h *AppShowMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Show()
	return nil, nil
}

func (h *AppShowMethod) GetName() string            { return "show" }
func (h *AppShowMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *AppShowMethod) GetIsStatic() bool          { return true }
func (h *AppShowMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *AppShowMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *AppShowMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
