//go:build js && wasm

package main

import (
	"image"
	"image/color"
	"log"
	"os"
	"syscall/js"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/pointer"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

var versionProg = ""

func main() {
	log.Println("ffff")
	var cb js.Func
	cb = js.FuncOf(func(this js.Value, args []js.Value) any {
		log.Println("button clicked")
		js.Global().Get("document").Call("getElementById", "giowindow").Call("setAttribute", "style", "width:503px")
		go startWindow()
		//cb.Release() // release the function if the button will not be clicked again
		return nil
	})
	js.Global().Get("document").Call("getElementById", "myButton").Call("addEventListener", "click", cb)

	// go func() {
	// 	w := app.NewWindow(
	// 		app.Size(unit.Dp(700), unit.Dp(600)), //  x-y  windows
	// 	)
	// 	err := run(w)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	os.Exit(0)
	// }()
	//go startWindow()
	app.Main()
}

func startWindow() {
	w := app.NewWindow(
	//	app.Size(unit.Dp(1006), unit.Dp(600)), //  x-y  windows

	)

	err := run(w)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run(w *app.Window) error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops
	// button := &Button{}

	// button2 := &Button{}
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.StageEvent:
			// f := js.Global().Get("document")
			// f = f.Call("getElementById", "giowindow").Call("getElementsByTagName", "canvas")
			// g := f.Index(0) // Call("item", 0)
			// g.Call("setAttribute", "style", "width:1008px;height: 500px")
			// g.Call("setAttribute", "width", "1007")
			// log.Printf("teg %+v %+v", f, g)
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:

			//
			gtx := layout.NewContext(&ops, e)
			title := material.H3(th, "Превед УУ Е, Gio мы тута")
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 200}
			title.Color = maroon
			title.Alignment = text.Middle
			title.Layout(gtx)

			// e.Frame(gtx.Ops)
			// gtx2 := layout.NewContext(&ops, e)

			title2 := material.H1(th, "Превед УУ Е, Gio ")
			maroon2 := color.NRGBA{R: 0, G: 127, B: 0, A: 200}
			title2.Color = maroon2
			title2.Font.Style = text.Italic
			title2.Alignment = text.Middle
			title2.Layout(gtx)
			//..//m := &ButtonVisual{} // просто рисуем
			//...//m := &Button{} // через нажатие
			// button.Layout(gtx)
			// button2.Layout(gtx)
			// inset(gtx)
			// insetTop(gtx)
			flexed(gtx)

			e.Frame(gtx.Ops)
		}
	}
}

// рисуем кнопку --------------------
type ButtonVisual struct {
	pressed bool
}

func (b *ButtonVisual) Layout(gtx layout.Context) layout.Dimensions {
	col := color.NRGBA{R: 0x80, A: 0xFF}
	if b.pressed {
		col = color.NRGBA{G: 0x80, A: 0xFF}
	}
	return drawSquare(gtx.Ops, col)
}

func drawSquare(ops *op.Ops, color color.NRGBA) layout.Dimensions {
	defer clip.Rect{Max: image.Pt(100, 100)}.Push(ops).Pop()
	paint.ColorOp{Color: color}.Add(ops)
	paint.PaintOp{}.Add(ops)
	return layout.Dimensions{Size: image.Pt(100, 100)}
}

// Здесь нажатие кнопки.
type Button struct {
	pressed bool
}

func (b *Button) Layout(gtx layout.Context) layout.Dimensions {
	// here we loop through all the events associated with this button.
	for _, e := range gtx.Events(b) {
		if e, ok := e.(pointer.Event); ok {
			switch e.Type {
			case pointer.Press:
				b.pressed = true
			case pointer.Release:
				b.pressed = false
			}
		}
	}

	// Confine the area for pointer events.
	area := clip.Rect(image.Rect(0, 0, 100, 100)).Push(gtx.Ops)
	pointer.InputOp{
		Tag:   b,
		Types: pointer.Press | pointer.Release,
	}.Add(gtx.Ops)
	area.Pop()

	// Draw the button.
	col := color.NRGBA{B: 0x80, A: 0xFF}
	if b.pressed {
		log.Println("Press почему два?")
		col = color.NRGBA{G: 0x80, A: 0xFF}
	}
	return drawSquare(gtx.Ops, col)
}

// Test colors. -------------------------------------
var (
	background = color.NRGBA{R: 0xC0, G: 0xC0, B: 0xC0, A: 0xFF}
	red        = color.NRGBA{R: 0xC0, G: 0x40, B: 0x40, A: 0xFF}
	green      = color.NRGBA{R: 0x40, G: 0xC0, B: 0x40, A: 0xFF}
	blue       = color.NRGBA{R: 0x40, G: 0x40, B: 0xC0, A: 0xFF}
)

// ColorBox creates a widget with the specified dimensions and color.
func ColorBox(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}

//layout.Insetдобавляет пространство вокруг виджета.

func inset(gtx layout.Context) layout.Dimensions {
	// Draw rectangles inside of each other, with 30dp padding.
	return layout.UniformInset(unit.Dp(30)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		//return ColorBox(gtx, gtx.Constraints.Max, red)
		return ColorBox(gtx, image.Pt(100, 130), red)
	})
}

func insetTop(gtx layout.Context) layout.Dimensions {
	// Draw rectangles inside of each other, with 30dp padding.
	return layout.Inset{Top: unit.Dp(30)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		//return ColorBox(gtx, gtx.Constraints.Max, red)
		return ColorBox(gtx, image.Pt(100, 130), blue)
	})
}

// а вот и Flex !!
func flexed(gtx layout.Context) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(100, 100), red)
		}),
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(gtx.Constraints.Min.X, 100), blue)
		}),
		// layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		// 	return ColorBox(gtx, image.Pt(100, 100), red)
		// }),
		// Цифра указывает долю, из всех размеров оставшихся для Flex (не понятно но 0.5 это пополам)
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			//return ColorBox(gtx, gtx.Constraints.Min, green)
			return ColorBox(gtx, image.Pt(gtx.Constraints.Min.X, 100), green)
		}),
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			//return ColorBox(gtx, gtx.Constraints.Min, green)
			return ColorBox(gtx, image.Pt(gtx.Constraints.Max.X, 100), red)
		}),
	)
}
