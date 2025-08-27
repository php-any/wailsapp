package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowHandleDragAndDropMessageMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowHandleDragAndDropMessageMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.HandleDragAndDropMessage 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.HandleDragAndDropMessage 缺少参数, index: 1"))
	}

	arg0 := make([]string, 0)
	for _, v := range a0.(*data.ArrayValue).Value {
		switch elemVal := v.(type) {
		case *data.ClassValue:
			if p, ok := elemVal.Class.(interface{ GetSource() any }); ok {
				// 检查 GetSource 返回的类型，如果是指针则解引用
				if src := p.GetSource(); src != nil {
					if ptr, ok := src.(*string); ok {
						arg0 = append(arg0, *ptr)
					} else {
						arg0 = append(arg0, src.(string))
					}
				}
			} else {
				return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.HandleDragAndDropMessage 参数类型不支持, index: 0"))
			}
		case *data.ProxyValue:
			if p, ok := elemVal.Class.(interface{ GetSource() any }); ok {
				if src := p.GetSource(); src != nil {
					if ptr, ok := src.(*string); ok {
						arg0 = append(arg0, *ptr)
					} else {
						arg0 = append(arg0, src.(string))
					}
				}
			} else {
				return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.HandleDragAndDropMessage 参数类型不支持, index: 0"))
			}
		case *data.AnyValue:
			arg0 = append(arg0, elemVal.Value.(string))
		}
	}
	var arg1 *applicationsrc.DropZoneDetails
	switch v := a1.(type) {
	case *data.ClassValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			// 检查 GetSource 返回的类型，如果是指针则解引用
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(**applicationsrc.DropZoneDetails); ok {
					arg1 = *ptr
				} else {
					arg1 = src.(*applicationsrc.DropZoneDetails)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.HandleDragAndDropMessage 参数类型不支持, index: 1"))
		}
	case *data.ProxyValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(**applicationsrc.DropZoneDetails); ok {
					arg1 = *ptr
				} else {
					arg1 = src.(*applicationsrc.DropZoneDetails)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.HandleDragAndDropMessage 参数类型不支持, index: 1"))
		}
	case *data.AnyValue:
		arg1 = v.Value.(*applicationsrc.DropZoneDetails)
	default:
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.HandleDragAndDropMessage 参数类型不支持, index: 1"))
	}

	h.source.HandleDragAndDropMessage(arg0, arg1)
	return nil, nil
}

func (h *WebviewWindowHandleDragAndDropMessageMethod) GetName() string {
	return "handleDragAndDropMessage"
}
func (h *WebviewWindowHandleDragAndDropMessageMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *WebviewWindowHandleDragAndDropMessageMethod) GetIsStatic() bool { return true }
func (h *WebviewWindowHandleDragAndDropMessageMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "filenames", 0, nil, nil),
		node.NewParameter(nil, "dropZone", 1, nil, nil),
	}
}

func (h *WebviewWindowHandleDragAndDropMessageMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "filenames", 0, nil),
		node.NewVariable(nil, "dropZone", 1, nil),
	}
}

func (h *WebviewWindowHandleDragAndDropMessageMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
