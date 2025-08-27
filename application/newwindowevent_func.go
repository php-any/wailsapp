package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type NewWindowEventFunction struct{}

func NewNewWindowEventFunction() data.FuncStmt {
	return &NewWindowEventFunction{}
}

func (h *NewWindowEventFunction) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := applicationsrc.NewWindowEvent()
	return data.NewClassValue(NewWindowEventClassFrom(ret0), ctx), nil
}

func (h *NewWindowEventFunction) GetName() string            { return "wails\\application\\newWindowEvent" }
func (h *NewWindowEventFunction) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *NewWindowEventFunction) GetIsStatic() bool          { return true }
func (h *NewWindowEventFunction) GetParams() []data.GetValue {
	return []data.GetValue{}
}
func (h *NewWindowEventFunction) GetVariables() []data.Variable {
	return []data.Variable{}
}
func (h *NewWindowEventFunction) GetReturnType() data.Types { return data.NewBaseType("void") }
