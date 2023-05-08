package ui

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/font"
)

/*
	Returns new widget.TextInput. Accepts font, placeholder (text displayed when empty), text to display,
	Controller and UISignal associated with the TextInput.
*/
func NewTextInput(font *font.Face, placeholder, text string, c *Controller, s UISignal) *widget.TextInput {
	ti := widget.NewTextInput(
		widget.TextInputOpts.Image(loadTextInputImage()),

		widget.TextInputOpts.Face(*font),

		widget.TextInputOpts.Color(loadTextInputFontColor()),

		widget.TextInputOpts.Padding(widget.NewInsetsSimple(5)),

		widget.TextInputOpts.CaretOpts(
			widget.CaretOpts.Size(*font, 2),
		),

		widget.TextInputOpts.Placeholder(placeholder),

		widget.TextInputOpts.SubmitHandler(func(args *widget.TextInputChangedEventArgs) {
			c.SetSignal(s)
			c.SetRules(args.InputText)
		}),

		widget.TextInputOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.GridLayoutData{
					HorizontalPosition: widget.GridLayoutPositionCenter,
					VerticalPosition: widget.GridLayoutPositionStart,
				},
			),
			widget.WidgetOpts.MinSize(140, 10),
		),
	)

	ti.InputText = text

	return ti
}

func loadTextInputImage() *widget.TextInputImage {
	return &widget.TextInputImage{
		Idle:     image.NewNineSliceColor(color.RGBA{R: 100, G: 100, B: 100, A: 255}),
	}
}

func loadTextInputFontColor() *widget.TextInputColor {
	return &widget.TextInputColor{
		Idle:          color.White,
		Disabled:      color.RGBA{R: 200, G: 200, B: 200, A: 255},
		Caret:         color.White,
		DisabledCaret: color.RGBA{R: 200, G: 200, B: 200, A: 255},
	}
}
