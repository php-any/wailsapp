package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SystemTrayManagerNewMethod struct {
	source *applicationsrc.SystemTrayManager
}

func (h *SystemTrayManagerNewMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.New()
	return data.NewClassValue(NewSystemTrayClassFrom(ret0), ctx), nil
}

func (h *SystemTrayManagerNewMethod) GetName() string            { return "new" }
func (h *SystemTrayManagerNewMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *SystemTrayManagerNewMethod) GetIsStatic() bool          { return true }
func (h *SystemTrayManagerNewMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *SystemTrayManagerNewMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *SystemTrayManagerNewMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
