package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuManagerShowAboutMethod struct {
	source *applicationsrc.MenuManager
}

func (h *MenuManagerShowAboutMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.ShowAbout()
	return nil, nil
}

func (h *MenuManagerShowAboutMethod) GetName() string            { return "showAbout" }
func (h *MenuManagerShowAboutMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuManagerShowAboutMethod) GetIsStatic() bool          { return true }
func (h *MenuManagerShowAboutMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *MenuManagerShowAboutMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *MenuManagerShowAboutMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
