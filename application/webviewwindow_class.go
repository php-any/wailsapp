package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewWebviewWindowClass() data.ClassStmt {
	return &WebviewWindowClass{
		source:                         nil,
		construct:                      &WebviewWindowConstructMethod{source: nil},
		bounds:                         &WebviewWindowBoundsMethod{source: nil},
		callError:                      &WebviewWindowCallErrorMethod{source: nil},
		callResponse:                   &WebviewWindowCallResponseMethod{source: nil},
		center:                         &WebviewWindowCenterMethod{source: nil},
		close:                          &WebviewWindowCloseMethod{source: nil},
		dialogError:                    &WebviewWindowDialogErrorMethod{source: nil},
		dialogResponse:                 &WebviewWindowDialogResponseMethod{source: nil},
		disableSizeConstraints:         &WebviewWindowDisableSizeConstraintsMethod{source: nil},
		dispatchWailsEvent:             &WebviewWindowDispatchWailsEventMethod{source: nil},
		emitEvent:                      &WebviewWindowEmitEventMethod{source: nil},
		enableSizeConstraints:          &WebviewWindowEnableSizeConstraintsMethod{source: nil},
		error:                          &WebviewWindowErrorMethod{source: nil},
		execJS:                         &WebviewWindowExecJSMethod{source: nil},
		flash:                          &WebviewWindowFlashMethod{source: nil},
		focus:                          &WebviewWindowFocusMethod{source: nil},
		forceReload:                    &WebviewWindowForceReloadMethod{source: nil},
		fullscreen:                     &WebviewWindowFullscreenMethod{source: nil},
		getBorderSizes:                 &WebviewWindowGetBorderSizesMethod{source: nil},
		getScreen:                      &WebviewWindowGetScreenMethod{source: nil},
		getZoom:                        &WebviewWindowGetZoomMethod{source: nil},
		handleDragAndDropMessage:       &WebviewWindowHandleDragAndDropMessageMethod{source: nil},
		handleKeyEvent:                 &WebviewWindowHandleKeyEventMethod{source: nil},
		handleMessage:                  &WebviewWindowHandleMessageMethod{source: nil},
		handleWindowEvent:              &WebviewWindowHandleWindowEventMethod{source: nil},
		height:                         &WebviewWindowHeightMethod{source: nil},
		hide:                           &WebviewWindowHideMethod{source: nil},
		hideMenuBar:                    &WebviewWindowHideMenuBarMethod{source: nil},
		iD:                             &WebviewWindowIDMethod{source: nil},
		info:                           &WebviewWindowInfoMethod{source: nil},
		initiateFrontendDropProcessing: &WebviewWindowInitiateFrontendDropProcessingMethod{source: nil},
		isFocused:                      &WebviewWindowIsFocusedMethod{source: nil},
		isFullscreen:                   &WebviewWindowIsFullscreenMethod{source: nil},
		isIgnoreMouseEvents:            &WebviewWindowIsIgnoreMouseEventsMethod{source: nil},
		isMaximised:                    &WebviewWindowIsMaximisedMethod{source: nil},
		isMinimised:                    &WebviewWindowIsMinimisedMethod{source: nil},
		isVisible:                      &WebviewWindowIsVisibleMethod{source: nil},
		maximise:                       &WebviewWindowMaximiseMethod{source: nil},
		minimise:                       &WebviewWindowMinimiseMethod{source: nil},
		name:                           &WebviewWindowNameMethod{source: nil},
		nativeWindow:                   &WebviewWindowNativeWindowMethod{source: nil},
		onWindowEvent:                  &WebviewWindowOnWindowEventMethod{source: nil},
		openContextMenu:                &WebviewWindowOpenContextMenuMethod{source: nil},
		openDevTools:                   &WebviewWindowOpenDevToolsMethod{source: nil},
		physicalBounds:                 &WebviewWindowPhysicalBoundsMethod{source: nil},
		position:                       &WebviewWindowPositionMethod{source: nil},
		print:                          &WebviewWindowPrintMethod{source: nil},
		registerHook:                   &WebviewWindowRegisterHookMethod{source: nil},
		relativePosition:               &WebviewWindowRelativePositionMethod{source: nil},
		reload:                         &WebviewWindowReloadMethod{source: nil},
		resizable:                      &WebviewWindowResizableMethod{source: nil},
		restore:                        &WebviewWindowRestoreMethod{source: nil},
		run:                            &WebviewWindowRunMethod{source: nil},
		setAlwaysOnTop:                 &WebviewWindowSetAlwaysOnTopMethod{source: nil},
		setBackgroundColour:            &WebviewWindowSetBackgroundColourMethod{source: nil},
		setBounds:                      &WebviewWindowSetBoundsMethod{source: nil},
		setCloseButtonState:            &WebviewWindowSetCloseButtonStateMethod{source: nil},
		setContentProtection:           &WebviewWindowSetContentProtectionMethod{source: nil},
		setEnabled:                     &WebviewWindowSetEnabledMethod{source: nil},
		setFrameless:                   &WebviewWindowSetFramelessMethod{source: nil},
		setHTML:                        &WebviewWindowSetHTMLMethod{source: nil},
		setIgnoreMouseEvents:           &WebviewWindowSetIgnoreMouseEventsMethod{source: nil},
		setMaxSize:                     &WebviewWindowSetMaxSizeMethod{source: nil},
		setMaximiseButtonState:         &WebviewWindowSetMaximiseButtonStateMethod{source: nil},
		setMenu:                        &WebviewWindowSetMenuMethod{source: nil},
		setMinSize:                     &WebviewWindowSetMinSizeMethod{source: nil},
		setMinimiseButtonState:         &WebviewWindowSetMinimiseButtonStateMethod{source: nil},
		setPhysicalBounds:              &WebviewWindowSetPhysicalBoundsMethod{source: nil},
		setPosition:                    &WebviewWindowSetPositionMethod{source: nil},
		setRelativePosition:            &WebviewWindowSetRelativePositionMethod{source: nil},
		setResizable:                   &WebviewWindowSetResizableMethod{source: nil},
		setSize:                        &WebviewWindowSetSizeMethod{source: nil},
		setTitle:                       &WebviewWindowSetTitleMethod{source: nil},
		setURL:                         &WebviewWindowSetURLMethod{source: nil},
		setZoom:                        &WebviewWindowSetZoomMethod{source: nil},
		show:                           &WebviewWindowShowMethod{source: nil},
		showMenuBar:                    &WebviewWindowShowMenuBarMethod{source: nil},
		size:                           &WebviewWindowSizeMethod{source: nil},
		snapAssist:                     &WebviewWindowSnapAssistMethod{source: nil},
		toggleFrameless:                &WebviewWindowToggleFramelessMethod{source: nil},
		toggleFullscreen:               &WebviewWindowToggleFullscreenMethod{source: nil},
		toggleMaximise:                 &WebviewWindowToggleMaximiseMethod{source: nil},
		toggleMenuBar:                  &WebviewWindowToggleMenuBarMethod{source: nil},
		unFullscreen:                   &WebviewWindowUnFullscreenMethod{source: nil},
		unMaximise:                     &WebviewWindowUnMaximiseMethod{source: nil},
		unMinimise:                     &WebviewWindowUnMinimiseMethod{source: nil},
		width:                          &WebviewWindowWidthMethod{source: nil},
		zoom:                           &WebviewWindowZoomMethod{source: nil},
		zoomIn:                         &WebviewWindowZoomInMethod{source: nil},
		zoomOut:                        &WebviewWindowZoomOutMethod{source: nil},
		zoomReset:                      &WebviewWindowZoomResetMethod{source: nil},
	}
}

func NewWebviewWindowClassFrom(source *applicationsrc.WebviewWindow) data.ClassStmt {
	return &WebviewWindowClass{
		source:                         source,
		construct:                      &WebviewWindowConstructMethod{source: source},
		bounds:                         &WebviewWindowBoundsMethod{source: source},
		callError:                      &WebviewWindowCallErrorMethod{source: source},
		callResponse:                   &WebviewWindowCallResponseMethod{source: source},
		center:                         &WebviewWindowCenterMethod{source: source},
		close:                          &WebviewWindowCloseMethod{source: source},
		dialogError:                    &WebviewWindowDialogErrorMethod{source: source},
		dialogResponse:                 &WebviewWindowDialogResponseMethod{source: source},
		disableSizeConstraints:         &WebviewWindowDisableSizeConstraintsMethod{source: source},
		dispatchWailsEvent:             &WebviewWindowDispatchWailsEventMethod{source: source},
		emitEvent:                      &WebviewWindowEmitEventMethod{source: source},
		enableSizeConstraints:          &WebviewWindowEnableSizeConstraintsMethod{source: source},
		error:                          &WebviewWindowErrorMethod{source: source},
		execJS:                         &WebviewWindowExecJSMethod{source: source},
		flash:                          &WebviewWindowFlashMethod{source: source},
		focus:                          &WebviewWindowFocusMethod{source: source},
		forceReload:                    &WebviewWindowForceReloadMethod{source: source},
		fullscreen:                     &WebviewWindowFullscreenMethod{source: source},
		getBorderSizes:                 &WebviewWindowGetBorderSizesMethod{source: source},
		getScreen:                      &WebviewWindowGetScreenMethod{source: source},
		getZoom:                        &WebviewWindowGetZoomMethod{source: source},
		handleDragAndDropMessage:       &WebviewWindowHandleDragAndDropMessageMethod{source: source},
		handleKeyEvent:                 &WebviewWindowHandleKeyEventMethod{source: source},
		handleMessage:                  &WebviewWindowHandleMessageMethod{source: source},
		handleWindowEvent:              &WebviewWindowHandleWindowEventMethod{source: source},
		height:                         &WebviewWindowHeightMethod{source: source},
		hide:                           &WebviewWindowHideMethod{source: source},
		hideMenuBar:                    &WebviewWindowHideMenuBarMethod{source: source},
		iD:                             &WebviewWindowIDMethod{source: source},
		info:                           &WebviewWindowInfoMethod{source: source},
		initiateFrontendDropProcessing: &WebviewWindowInitiateFrontendDropProcessingMethod{source: source},
		isFocused:                      &WebviewWindowIsFocusedMethod{source: source},
		isFullscreen:                   &WebviewWindowIsFullscreenMethod{source: source},
		isIgnoreMouseEvents:            &WebviewWindowIsIgnoreMouseEventsMethod{source: source},
		isMaximised:                    &WebviewWindowIsMaximisedMethod{source: source},
		isMinimised:                    &WebviewWindowIsMinimisedMethod{source: source},
		isVisible:                      &WebviewWindowIsVisibleMethod{source: source},
		maximise:                       &WebviewWindowMaximiseMethod{source: source},
		minimise:                       &WebviewWindowMinimiseMethod{source: source},
		name:                           &WebviewWindowNameMethod{source: source},
		nativeWindow:                   &WebviewWindowNativeWindowMethod{source: source},
		onWindowEvent:                  &WebviewWindowOnWindowEventMethod{source: source},
		openContextMenu:                &WebviewWindowOpenContextMenuMethod{source: source},
		openDevTools:                   &WebviewWindowOpenDevToolsMethod{source: source},
		physicalBounds:                 &WebviewWindowPhysicalBoundsMethod{source: source},
		position:                       &WebviewWindowPositionMethod{source: source},
		print:                          &WebviewWindowPrintMethod{source: source},
		registerHook:                   &WebviewWindowRegisterHookMethod{source: source},
		relativePosition:               &WebviewWindowRelativePositionMethod{source: source},
		reload:                         &WebviewWindowReloadMethod{source: source},
		resizable:                      &WebviewWindowResizableMethod{source: source},
		restore:                        &WebviewWindowRestoreMethod{source: source},
		run:                            &WebviewWindowRunMethod{source: source},
		setAlwaysOnTop:                 &WebviewWindowSetAlwaysOnTopMethod{source: source},
		setBackgroundColour:            &WebviewWindowSetBackgroundColourMethod{source: source},
		setBounds:                      &WebviewWindowSetBoundsMethod{source: source},
		setCloseButtonState:            &WebviewWindowSetCloseButtonStateMethod{source: source},
		setContentProtection:           &WebviewWindowSetContentProtectionMethod{source: source},
		setEnabled:                     &WebviewWindowSetEnabledMethod{source: source},
		setFrameless:                   &WebviewWindowSetFramelessMethod{source: source},
		setHTML:                        &WebviewWindowSetHTMLMethod{source: source},
		setIgnoreMouseEvents:           &WebviewWindowSetIgnoreMouseEventsMethod{source: source},
		setMaxSize:                     &WebviewWindowSetMaxSizeMethod{source: source},
		setMaximiseButtonState:         &WebviewWindowSetMaximiseButtonStateMethod{source: source},
		setMenu:                        &WebviewWindowSetMenuMethod{source: source},
		setMinSize:                     &WebviewWindowSetMinSizeMethod{source: source},
		setMinimiseButtonState:         &WebviewWindowSetMinimiseButtonStateMethod{source: source},
		setPhysicalBounds:              &WebviewWindowSetPhysicalBoundsMethod{source: source},
		setPosition:                    &WebviewWindowSetPositionMethod{source: source},
		setRelativePosition:            &WebviewWindowSetRelativePositionMethod{source: source},
		setResizable:                   &WebviewWindowSetResizableMethod{source: source},
		setSize:                        &WebviewWindowSetSizeMethod{source: source},
		setTitle:                       &WebviewWindowSetTitleMethod{source: source},
		setURL:                         &WebviewWindowSetURLMethod{source: source},
		setZoom:                        &WebviewWindowSetZoomMethod{source: source},
		show:                           &WebviewWindowShowMethod{source: source},
		showMenuBar:                    &WebviewWindowShowMenuBarMethod{source: source},
		size:                           &WebviewWindowSizeMethod{source: source},
		snapAssist:                     &WebviewWindowSnapAssistMethod{source: source},
		toggleFrameless:                &WebviewWindowToggleFramelessMethod{source: source},
		toggleFullscreen:               &WebviewWindowToggleFullscreenMethod{source: source},
		toggleMaximise:                 &WebviewWindowToggleMaximiseMethod{source: source},
		toggleMenuBar:                  &WebviewWindowToggleMenuBarMethod{source: source},
		unFullscreen:                   &WebviewWindowUnFullscreenMethod{source: source},
		unMaximise:                     &WebviewWindowUnMaximiseMethod{source: source},
		unMinimise:                     &WebviewWindowUnMinimiseMethod{source: source},
		width:                          &WebviewWindowWidthMethod{source: source},
		zoom:                           &WebviewWindowZoomMethod{source: source},
		zoomIn:                         &WebviewWindowZoomInMethod{source: source},
		zoomOut:                        &WebviewWindowZoomOutMethod{source: source},
		zoomReset:                      &WebviewWindowZoomResetMethod{source: source},
	}
}

type WebviewWindowClass struct {
	node.Node
	source                         *applicationsrc.WebviewWindow
	construct                      data.Method
	bounds                         data.Method
	callError                      data.Method
	callResponse                   data.Method
	center                         data.Method
	close                          data.Method
	dialogError                    data.Method
	dialogResponse                 data.Method
	disableSizeConstraints         data.Method
	dispatchWailsEvent             data.Method
	emitEvent                      data.Method
	enableSizeConstraints          data.Method
	error                          data.Method
	execJS                         data.Method
	flash                          data.Method
	focus                          data.Method
	forceReload                    data.Method
	fullscreen                     data.Method
	getBorderSizes                 data.Method
	getScreen                      data.Method
	getZoom                        data.Method
	handleDragAndDropMessage       data.Method
	handleKeyEvent                 data.Method
	handleMessage                  data.Method
	handleWindowEvent              data.Method
	height                         data.Method
	hide                           data.Method
	hideMenuBar                    data.Method
	iD                             data.Method
	info                           data.Method
	initiateFrontendDropProcessing data.Method
	isFocused                      data.Method
	isFullscreen                   data.Method
	isIgnoreMouseEvents            data.Method
	isMaximised                    data.Method
	isMinimised                    data.Method
	isVisible                      data.Method
	maximise                       data.Method
	minimise                       data.Method
	name                           data.Method
	nativeWindow                   data.Method
	onWindowEvent                  data.Method
	openContextMenu                data.Method
	openDevTools                   data.Method
	physicalBounds                 data.Method
	position                       data.Method
	print                          data.Method
	registerHook                   data.Method
	relativePosition               data.Method
	reload                         data.Method
	resizable                      data.Method
	restore                        data.Method
	run                            data.Method
	setAlwaysOnTop                 data.Method
	setBackgroundColour            data.Method
	setBounds                      data.Method
	setCloseButtonState            data.Method
	setContentProtection           data.Method
	setEnabled                     data.Method
	setFrameless                   data.Method
	setHTML                        data.Method
	setIgnoreMouseEvents           data.Method
	setMaxSize                     data.Method
	setMaximiseButtonState         data.Method
	setMenu                        data.Method
	setMinSize                     data.Method
	setMinimiseButtonState         data.Method
	setPhysicalBounds              data.Method
	setPosition                    data.Method
	setRelativePosition            data.Method
	setResizable                   data.Method
	setSize                        data.Method
	setTitle                       data.Method
	setURL                         data.Method
	setZoom                        data.Method
	show                           data.Method
	showMenuBar                    data.Method
	size                           data.Method
	snapAssist                     data.Method
	toggleFrameless                data.Method
	toggleFullscreen               data.Method
	toggleMaximise                 data.Method
	toggleMenuBar                  data.Method
	unFullscreen                   data.Method
	unMaximise                     data.Method
	unMinimise                     data.Method
	width                          data.Method
	zoom                           data.Method
	zoomIn                         data.Method
	zoomOut                        data.Method
	zoomReset                      data.Method
}

func (s *WebviewWindowClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewWebviewWindowClassFrom(&applicationsrc.WebviewWindow{}), ctx.CreateBaseContext()), nil
}
func (s *WebviewWindowClass) GetName() string                            { return "wails\\application\\WebviewWindow" }
func (s *WebviewWindowClass) GetExtend() *string                         { return nil }
func (s *WebviewWindowClass) GetImplements() []string                    { return nil }
func (s *WebviewWindowClass) AsString() string                           { return "WebviewWindow{}" }
func (s *WebviewWindowClass) GetSource() any                             { return s.source }
func (s *WebviewWindowClass) GetProperty(_ string) (data.Property, bool) { return nil, false }
func (s *WebviewWindowClass) GetProperties() map[string]data.Property    { return nil }
func (s *WebviewWindowClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "bounds":
		return s.bounds, true
	case "callError":
		return s.callError, true
	case "callResponse":
		return s.callResponse, true
	case "center":
		return s.center, true
	case "close":
		return s.close, true
	case "dialogError":
		return s.dialogError, true
	case "dialogResponse":
		return s.dialogResponse, true
	case "disableSizeConstraints":
		return s.disableSizeConstraints, true
	case "dispatchWailsEvent":
		return s.dispatchWailsEvent, true
	case "emitEvent":
		return s.emitEvent, true
	case "enableSizeConstraints":
		return s.enableSizeConstraints, true
	case "error":
		return s.error, true
	case "execJS":
		return s.execJS, true
	case "flash":
		return s.flash, true
	case "focus":
		return s.focus, true
	case "forceReload":
		return s.forceReload, true
	case "fullscreen":
		return s.fullscreen, true
	case "getBorderSizes":
		return s.getBorderSizes, true
	case "getScreen":
		return s.getScreen, true
	case "getZoom":
		return s.getZoom, true
	case "handleDragAndDropMessage":
		return s.handleDragAndDropMessage, true
	case "handleKeyEvent":
		return s.handleKeyEvent, true
	case "handleMessage":
		return s.handleMessage, true
	case "handleWindowEvent":
		return s.handleWindowEvent, true
	case "height":
		return s.height, true
	case "hide":
		return s.hide, true
	case "hideMenuBar":
		return s.hideMenuBar, true
	case "iD":
		return s.iD, true
	case "info":
		return s.info, true
	case "initiateFrontendDropProcessing":
		return s.initiateFrontendDropProcessing, true
	case "isFocused":
		return s.isFocused, true
	case "isFullscreen":
		return s.isFullscreen, true
	case "isIgnoreMouseEvents":
		return s.isIgnoreMouseEvents, true
	case "isMaximised":
		return s.isMaximised, true
	case "isMinimised":
		return s.isMinimised, true
	case "isVisible":
		return s.isVisible, true
	case "maximise":
		return s.maximise, true
	case "minimise":
		return s.minimise, true
	case "name":
		return s.name, true
	case "nativeWindow":
		return s.nativeWindow, true
	case "onWindowEvent":
		return s.onWindowEvent, true
	case "openContextMenu":
		return s.openContextMenu, true
	case "openDevTools":
		return s.openDevTools, true
	case "physicalBounds":
		return s.physicalBounds, true
	case "position":
		return s.position, true
	case "print":
		return s.print, true
	case "registerHook":
		return s.registerHook, true
	case "relativePosition":
		return s.relativePosition, true
	case "reload":
		return s.reload, true
	case "resizable":
		return s.resizable, true
	case "restore":
		return s.restore, true
	case "run":
		return s.run, true
	case "setAlwaysOnTop":
		return s.setAlwaysOnTop, true
	case "setBackgroundColour":
		return s.setBackgroundColour, true
	case "setBounds":
		return s.setBounds, true
	case "setCloseButtonState":
		return s.setCloseButtonState, true
	case "setContentProtection":
		return s.setContentProtection, true
	case "setEnabled":
		return s.setEnabled, true
	case "setFrameless":
		return s.setFrameless, true
	case "setHTML":
		return s.setHTML, true
	case "setIgnoreMouseEvents":
		return s.setIgnoreMouseEvents, true
	case "setMaxSize":
		return s.setMaxSize, true
	case "setMaximiseButtonState":
		return s.setMaximiseButtonState, true
	case "setMenu":
		return s.setMenu, true
	case "setMinSize":
		return s.setMinSize, true
	case "setMinimiseButtonState":
		return s.setMinimiseButtonState, true
	case "setPhysicalBounds":
		return s.setPhysicalBounds, true
	case "setPosition":
		return s.setPosition, true
	case "setRelativePosition":
		return s.setRelativePosition, true
	case "setResizable":
		return s.setResizable, true
	case "setSize":
		return s.setSize, true
	case "setTitle":
		return s.setTitle, true
	case "setURL":
		return s.setURL, true
	case "setZoom":
		return s.setZoom, true
	case "show":
		return s.show, true
	case "showMenuBar":
		return s.showMenuBar, true
	case "size":
		return s.size, true
	case "snapAssist":
		return s.snapAssist, true
	case "toggleFrameless":
		return s.toggleFrameless, true
	case "toggleFullscreen":
		return s.toggleFullscreen, true
	case "toggleMaximise":
		return s.toggleMaximise, true
	case "toggleMenuBar":
		return s.toggleMenuBar, true
	case "unFullscreen":
		return s.unFullscreen, true
	case "unMaximise":
		return s.unMaximise, true
	case "unMinimise":
		return s.unMinimise, true
	case "width":
		return s.width, true
	case "zoom":
		return s.zoom, true
	case "zoomIn":
		return s.zoomIn, true
	case "zoomOut":
		return s.zoomOut, true
	case "zoomReset":
		return s.zoomReset, true
	}
	return nil, false
}

func (s *WebviewWindowClass) GetMethods() []data.Method {
	return []data.Method{
		s.bounds,
		s.callError,
		s.callResponse,
		s.center,
		s.close,
		s.dialogError,
		s.dialogResponse,
		s.disableSizeConstraints,
		s.dispatchWailsEvent,
		s.emitEvent,
		s.enableSizeConstraints,
		s.error,
		s.execJS,
		s.flash,
		s.focus,
		s.forceReload,
		s.fullscreen,
		s.getBorderSizes,
		s.getScreen,
		s.getZoom,
		s.handleDragAndDropMessage,
		s.handleKeyEvent,
		s.handleMessage,
		s.handleWindowEvent,
		s.height,
		s.hide,
		s.hideMenuBar,
		s.iD,
		s.info,
		s.initiateFrontendDropProcessing,
		s.isFocused,
		s.isFullscreen,
		s.isIgnoreMouseEvents,
		s.isMaximised,
		s.isMinimised,
		s.isVisible,
		s.maximise,
		s.minimise,
		s.name,
		s.nativeWindow,
		s.onWindowEvent,
		s.openContextMenu,
		s.openDevTools,
		s.physicalBounds,
		s.position,
		s.print,
		s.registerHook,
		s.relativePosition,
		s.reload,
		s.resizable,
		s.restore,
		s.run,
		s.setAlwaysOnTop,
		s.setBackgroundColour,
		s.setBounds,
		s.setCloseButtonState,
		s.setContentProtection,
		s.setEnabled,
		s.setFrameless,
		s.setHTML,
		s.setIgnoreMouseEvents,
		s.setMaxSize,
		s.setMaximiseButtonState,
		s.setMenu,
		s.setMinSize,
		s.setMinimiseButtonState,
		s.setPhysicalBounds,
		s.setPosition,
		s.setRelativePosition,
		s.setResizable,
		s.setSize,
		s.setTitle,
		s.setURL,
		s.setZoom,
		s.show,
		s.showMenuBar,
		s.size,
		s.snapAssist,
		s.toggleFrameless,
		s.toggleFullscreen,
		s.toggleMaximise,
		s.toggleMenuBar,
		s.unFullscreen,
		s.unMaximise,
		s.unMinimise,
		s.width,
		s.zoom,
		s.zoomIn,
		s.zoomOut,
		s.zoomReset,
	}
}

func (s *WebviewWindowClass) GetConstruct() data.Method { return s.construct }

type WebviewWindowConstructMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *WebviewWindowConstructMethod) GetName() string               { return "construct" }
func (h *WebviewWindowConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *WebviewWindowConstructMethod) GetIsStatic() bool             { return false }
func (h *WebviewWindowConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *WebviewWindowConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *WebviewWindowConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
