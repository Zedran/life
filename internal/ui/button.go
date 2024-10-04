package ui

import (
	"github.com/Zedran/life/internal/config/theme"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/font"
)

/* Returns the new widget.Button given font, text to display, Controller and signal corresponding with the Button. */
func NewButton(bt *theme.ButtonTheme, font *font.Face, text string, c *Controller, s UISignal) *widget.Button {
	button := widget.NewButton(
		widget.ButtonOpts.Image(loadButtonImage(bt)),

		widget.ButtonOpts.Text(
			text,
			*font,
			&widget.ButtonTextColor{
				Idle: bt.Text,
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

/* Returns a pointer to the container designed to hold buttons that serve similar functions. */
func NewButtonCluster() *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Spacing(5),
			),
		),
	)
}

func loadButtonImage(bt *theme.ButtonTheme) *widget.ButtonImage {
	return &widget.ButtonImage{
		Idle:    image.NewNineSliceColor(bt.Idle),
		Hover:   image.NewNineSliceColor(bt.Hover),
		Pressed: image.NewNineSliceColor(bt.Pressed),
	}
}
