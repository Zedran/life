package ui

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/font"
)

/* Returns the new widget.Button given font, text to display, Controller and signal corresponding with the Button. */
func NewButton(font *font.Face, text string, c *Controller, s UISignal) *widget.Button {
	button := widget.NewButton(
		widget.ButtonOpts.Image(loadButtonImage()),
		
		widget.ButtonOpts.Text(
			text, 
			*font, 
			&widget.ButtonTextColor{
				Idle: color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
			},
		),
	
		widget.ButtonOpts.TextPadding(
			widget.Insets{
				Left:   10,
				Right:  10,
				Top:    5,
				Bottom: 5,
			},
		),

		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.GridLayoutData{
					HorizontalPosition: widget.GridLayoutPositionCenter,
					VerticalPosition: widget.GridLayoutPositionStart,
				},
			),
		),
	)

	button.Configure(
		widget.ButtonOpts.PressedHandler(
			func(pressedArgs *widget.ButtonPressedEventArgs) {
				b := pressedArgs.Button.GetWidget()

				if button.GetWidget() == b {
					c.SetLMBDownOn(b)
				}
			},
		),

		widget.ButtonOpts.ReleasedHandler(
			func(releasedArgs *widget.ButtonReleasedEventArgs) {
				b := releasedArgs.Button.GetWidget()

				if releasedArgs.Inside && c.IsLMBDownOn(b) {
					c.SetSignal(s)
					c.ClearLMBPointer()
				}
			},
		),
	)

	return button
}

func loadButtonImage() *widget.ButtonImage {
	return &widget.ButtonImage{
		Idle   : image.NewNineSliceColor(color.NRGBA{0, 110, 0, 255}),
		Hover  : image.NewNineSliceColor(color.NRGBA{0, 170, 0, 255}),
		Pressed: image.NewNineSliceColor(color.NRGBA{0, 80, 0, 255}),
	}
}
