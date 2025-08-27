package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type EnvironmentManagerOpenFileManagerMethod struct {
	source *applicationsrc.EnvironmentManager
}

func (h *EnvironmentManagerOpenFileManagerMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("EnvironmentManager.OpenFileManager 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("EnvironmentManager.OpenFileManager 缺少参数, index: 1"))
	}

	arg0 := a0.(*data.StringValue).AsString()
	arg1, err := a1.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	if err := h.source.OpenFileManager(arg0, arg1); err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	return nil, nil
}

func (h *EnvironmentManagerOpenFileManagerMethod) GetName() string { return "openFileManager" }
func (h *EnvironmentManagerOpenFileManagerMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *EnvironmentManagerOpenFileManagerMethod) GetIsStatic() bool { return true }
func (h *EnvironmentManagerOpenFileManagerMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "path", 0, nil, nil),
		node.NewParameter(nil, "selectFile", 1, nil, nil),
	}
}

func (h *EnvironmentManagerOpenFileManagerMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "path", 0, nil),
		node.NewVariable(nil, "selectFile", 1, nil),
	}
}

func (h *EnvironmentManagerOpenFileManagerMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
