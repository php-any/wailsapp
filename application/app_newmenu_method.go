package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type AppNewMenuMethod struct {
	source *applicationsrc.App
}

func (h *AppNewMenuMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.NewMenu()
	return data.NewClassValue(NewMenuClassFrom(ret0), ctx), nil
}

func (h *AppNewMenuMethod) GetName() string            { return "newMenu" }
func (h *AppNewMenuMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *AppNewMenuMethod) GetIsStatic() bool          { return true }
func (h *AppNewMenuMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *AppNewMenuMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *AppNewMenuMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
