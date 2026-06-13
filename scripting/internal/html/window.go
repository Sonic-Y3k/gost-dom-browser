package html

import (
	"errors"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/entity"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func Window_window[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return cbCtx.This(), nil
}

// windowName backs the Window.name attribute. The spec models name as a string
// that defaults to "" and persists across reads; it is stored as a per-window
// component because the html.Window model has no name field of its own.
type windowName struct{ value string }

// Window_name implements the Window.name getter, returning the name previously
// assigned to this window, or "" if none has been set.
func Window_name[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	win, err := js.As[html.Window](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	if n, ok := entity.ComponentType[*windowName](win); ok {
		return codec.EncodeString(cbCtx, n.value)
	}
	return codec.EncodeString(cbCtx, "")
}

// Window_setName implements the Window.name setter, storing the assigned value
// on the window so a subsequent read returns it.
func Window_setName[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	win, err := js.As[html.Window](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	value, err := js.ParseSetterArg(cbCtx, codec.DecodeString)
	if err != nil {
		return nil, err
	}
	n, ok := entity.ComponentType[*windowName](win)
	if !ok {
		n = &windowName{}
		entity.SetComponentType(win, n)
	}
	n.value = value
	return nil, nil
}

func Window_history[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	win, err := js.As[html.Window](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return cbCtx.Constructor("History").NewInstance(win.History())
}

func Window_self[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return cbCtx.This(), nil
}

func Window_parent[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return cbCtx.This(), nil
}

func Window_opener[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return cbCtx.Null(), nil
}

func Window_setOpener[T any](_ js.CallbackContext[T]) (js.Value[T], error) {
	return nil, errors.New("Not implemented")
}

func encodeNavigator[T any](s js.Scope[T], n *html.Navigator) (js.Value[T], error) {
	return codec.EncodeEntityScopedWithPrototype(s, n, "Navigator")
}
