package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SystemTrayHideMethod struct {
	source *applicationsrc.SystemTray
}

func (h *SystemTrayHideMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Hide()
	return nil, nil
}

func (h *SystemTrayHideMethod) GetName() string            { return "hide" }
func (h *SystemTrayHideMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *SystemTrayHideMethod) GetIsStatic() bool          { return true }
func (h *SystemTrayHideMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *SystemTrayHideMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *SystemTrayHideMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
