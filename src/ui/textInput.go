package ui

import (
	"github.com/Zedran/life/src/config/theme"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/font"
)

/*
Returns new widget.TextInput. Accepts font, placeholder (text displayed when empty), text to display,
Controller and UISignal associated with the TextInput.
*/
func NewTextInput(tit *theme.TextInputTheme, font *font.Face, placeholder, text string, c *Controller, s UISignal) *widget.TextInput {
	ti := widget.NewTextInput(
		widget.TextInputOpts.Image(loadTextInputImage(tit)),

		widget.TextInputOpts.Face(*font),

		widget.TextInputOpts.Color(loadTextInputFontColor(tit)),

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
					VerticalPosition:   widget.GridLayoutPositionStart,
				},
			),
			widget.WidgetOpts.MinSize(140, 10),
		),
	)

	ti.InputText = text

	return ti
}

func loadTextInputImage(tit *theme.TextInputTheme) *widget.TextInputImage {
	return &widget.TextInputImage{
		Idle: image.NewNineSliceColor(tit.Text),
	}
}

func loadTextInputFontColor(tit *theme.TextInputTheme) *widget.TextInputColor {
	return &widget.TextInputColor{
		Idle:          tit.Idle,
		Disabled:      tit.Disabled,
		Caret:         tit.Caret,
		DisabledCaret: tit.DisabledCaret,
	}
}
