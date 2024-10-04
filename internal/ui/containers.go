package ui

import (
	"github.com/Zedran/life/internal/config/theme"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

/* Creates the info panel holding information about the current state of the game. */
func createInfo(uit *theme.UITheme) *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Padding(
					widget.Insets{
						Left:   10,
						Right:  10,
						Top:    5,
						Bottom: 0,
					},
				),
				widget.GridLayoutOpts.Columns(3),
				widget.GridLayoutOpts.Spacing(50, 5),
				widget.GridLayoutOpts.Stretch([]bool{true, true, true}, []bool{false}),
			),
		),
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(uit.InfoPanel.Background)),
	)
}

/* Creates the panel container that allows the user to control the game. */
func createPanel(uit *theme.UITheme) *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Padding(
					widget.Insets{
						Left:   10,
						Right:  10,
						Top:    5,
						Bottom: 5,
					},
				),
				widget.GridLayoutOpts.Columns(4),
				widget.GridLayoutOpts.Spacing(15, 15),
			),
		),
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(uit.CtrlPanel.Background)),
	)
}

/* Creates a spacer container that pushes info and panel to the bottom of the screen. */
func createSpacer() *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(1),
			),
		),
	)
}

/* Creates the root container of the UI. */
func createRoot() *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(1),
				widget.GridLayoutOpts.Stretch([]bool{true}, []bool{true, false, false}),
			),
		),
	)
}
