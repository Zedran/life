package ui

import (
	"github.com/Zedran/life/src/config/theme"
	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/font"
)

/* Creates a label with specified text and returns a pointer to it. */
func NewLabel(lt *theme.LabelTheme, font *font.Face, text string) *widget.Label {
	label := widget.NewLabel(
		widget.LabelOpts.Text(text, *font, loadLabelColor(lt)),
	)

	return label
}

/*
	Creates a special kind of label. It consists of 2 widget.Label instances. One of them serves as a title, 
	denoting a type of information displayed and another is holding value that can be changed. 
	It is used to provide insight to game variables such as current generation number, current speed or 
	zoom level. This function returns a container holding both labels and the label serving as a value holder. 
	Accepts the title text.
*/
func NewLabeledDisplay(lt *theme.LabelTheme, font *font.Face, labelText string) (*widget.Container, *widget.Label) {
	label := NewLabel(lt, font, labelText)
	val   := NewLabel(lt, font, "")

	display := widget.NewContainer(
		widget.ContainerOpts.Layout(	
			widget.NewRowLayout(
				widget.RowLayoutOpts.Padding(
					widget.Insets{
						Left:   5,
						Right:  5,
						Top:    0,
						Bottom: 0,
					},
				),
				widget.RowLayoutOpts.Spacing(10),
			),
		),
	)

	display.AddChild(label)
	display.AddChild(val)

	return display, val
}

func loadLabelColor(lt *theme.LabelTheme) *widget.LabelColor {
	return &widget.LabelColor{
		Idle: lt.Text,
	}
}
