package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SystemTrayDestroyMethod struct {
	source *applicationsrc.SystemTray
}

func (h *SystemTrayDestroyMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Destroy()
	return nil, nil
}

func (h *SystemTrayDestroyMethod) GetName() string            { return "destroy" }
func (h *SystemTrayDestroyMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *SystemTrayDestroyMethod) GetIsStatic() bool          { return true }
func (h *SystemTrayDestroyMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *SystemTrayDestroyMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *SystemTrayDestroyMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
