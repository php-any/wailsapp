package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SystemTrayOpenMenuMethod struct {
	source *applicationsrc.SystemTray
}

func (h *SystemTrayOpenMenuMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.OpenMenu()
	return nil, nil
}

func (h *SystemTrayOpenMenuMethod) GetName() string            { return "openMenu" }
func (h *SystemTrayOpenMenuMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *SystemTrayOpenMenuMethod) GetIsStatic() bool          { return true }
func (h *SystemTrayOpenMenuMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *SystemTrayOpenMenuMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *SystemTrayOpenMenuMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
