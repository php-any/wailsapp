package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type AppHideMethod struct {
	source *applicationsrc.App
}

func (h *AppHideMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Hide()
	return nil, nil
}

func (h *AppHideMethod) GetName() string            { return "hide" }
func (h *AppHideMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *AppHideMethod) GetIsStatic() bool          { return true }
func (h *AppHideMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *AppHideMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *AppHideMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
