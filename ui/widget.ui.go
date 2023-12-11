package ui

import (
	"fmt"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"github.com/tasnimzotder/artificial-life/data"
	"github.com/tasnimzotder/artificial-life/settings"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"image/color"
	"log"
)

type UiWidget struct {
	RootContainer *widget.Container
}

func NewUiWidget(gs *settings.GameSettings) *UiWidget {
	uw := &UiWidget{}
	uw.RootContainer = getRootContainer(widget.DirectionVertical)

	//buttonImage, _ := loadButtonImage()
	//face, _ := loadFont(16)

	startStopButton := getButton("Start/Stop", gs)
	resetButton := getButton("Reset", gs)

	gameTypetComboBox := getCombobox("GameType", gs)
	presetComboBox := getCombobox("Preset", gs)
	//ruleComboBox := getCombobox("LtLRule", gs)

	// sliders
	speedSlider := getSlider("Speed", gs)

	// inner containers
	innerContainer1 := getInnerContainer(widget.DirectionHorizontal)
	innerContainer2 := getInnerContainer(widget.DirectionVertical)
	innerContainerSliders := getInnerContainer(widget.DirectionVertical)

	// inner container 1
	innerContainer1.AddChild(startStopButton)
	innerContainer1.AddChild(resetButton)

	// inner container 2
	innerContainer2.AddChild(gameTypetComboBox)
	innerContainer2.AddChild(presetComboBox)
	//innerContainer2.AddChild(ruleComboBox)

	innerContainerSliders.AddChild(speedSlider)

	//uw.RootContainer.AddChild(uw.StartStopButton)
	uw.RootContainer.AddChild(innerContainerSliders)
	uw.RootContainer.AddChild(innerContainer2)
	uw.RootContainer.AddChild(innerContainer1)

	return uw
}

func loadButtonImage() (*widget.ButtonImage, error) {
	//idle := image.NewNineSliceColor(color.NRGBA{R: 170, G: 170, B: 180, A: 255})
	idle := image.NewNineSliceColor(color.NRGBA{R: 50, G: 50, B: 50, A: 255})

	hover := image.NewNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255})

	pressed := image.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 120, A: 255})

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}

func loadFont(size float64) (font.Face, error) {
	ttfFont, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ttfFont, &truetype.Options{
		Size:    size,
		DPI:     144,
		Hinting: font.HintingFull,
	}), nil
}

func getButton(name string, gs *settings.GameSettings) *widget.Button {
	buttonImage, _ := loadButtonImage()
	face, _ := loadFont(7)

	button := widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),

		widget.ButtonOpts.Image(buttonImage),
		widget.ButtonOpts.Text(name, face, &widget.ButtonTextColor{
			Idle: color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
		}),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   5,
			Right:  5,
			Top:    5,
			Bottom: 5,
		}),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			// todo: fix it
			if name == "Start/Stop" {
				gs.IsPaused = !gs.IsPaused
				log.Printf("Button clicked: %s", name)
			} else if name == "Reset" {
				gs.IsReset = true
				log.Printf("Button clicked: %s", name)
			}

		}),
	)

	return button
}

func getInnerContainer(direction widget.Direction) *widget.Container {
	innerContainer := widget.NewContainer(
		//widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0, 0, 0, 0xff})),
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(direction),
				widget.RowLayoutOpts.Padding(widget.Insets{
					Top:    5,
					Left:   10,
					Right:  10,
					Bottom: 10,
				}),
				widget.RowLayoutOpts.Spacing(10),
			),
		),

		//widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{255, 0, 0, 255})),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				//Specify where within the row or column this element should be positioned.
				Position: widget.RowLayoutPositionStart,
				//Should this widget be stretched across the row or column
				Stretch: false,
				//How wide can this element grow to (override preferred widget size)
				//MaxWidth: 100,
				//How tall can this element grow to (override preferred widget size)
				//MaxHeight: 100,
			}),
			//widget.WidgetOpts.MinSize(10, 10),
		),
	)

	return innerContainer
}

func getRootContainer(direction widget.Direction) *widget.Container {
	rootContainer := widget.NewContainer(
		//widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),
		widget.ContainerOpts.Layout(
			//widget.NewAnchorLayout(
			//	widget.AnchorLayoutOpts.Padding(widget.Insets{
			//		Top: 500,
			//	}),
			//),
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(direction),
				widget.RowLayoutOpts.Padding(widget.Insets{
					Top:    180,
					Left:   int(float64(ScreenWidth) * 0.65),
					Right:  50,
					Bottom: 50,
				}),
				widget.RowLayoutOpts.Spacing(10),
			),
		),
	)

	return rootContainer
}

type ListEntry struct {
	id   int
	name string
}

func getCombobox(comboType string, gs *settings.GameSettings) *widget.ListComboButton {
	buttonImage, _ := loadButtonImage()
	face, _ := loadFont(7)

	numEntries := 0
	var entries []any
	labelName := ""

	if comboType == "GameType" {
		numEntries = len(gs.GameTypes)
		entries = make([]any, 0, numEntries)
		labelName = "GameType: "

		for i, v := range gs.GameTypes {
			entries = append(entries, ListEntry{
				id:   i,
				name: v,
			})
		}
	} else if comboType == "Preset" {
		numEntries = len(gs.Presets)
		entries = make([]any, 0, numEntries)
		labelName = "Preset: "

		for i, v := range gs.Presets {
			entries = append(entries, ListEntry{
				id:   i,
				name: v.Name,
			})
		}
	} else if comboType == "LtLRule" {
		numEntries = len(data.GetLtLRuleNames())
		entries = make([]any, 0, numEntries)
		labelName = "LtL Rule: "

		for i, v := range data.GetLtLRuleNames() {
			entries = append(entries, ListEntry{
				id:   i,
				name: v,
			})
		}
	}

	combobox := widget.NewListComboButton(
		widget.ListComboButtonOpts.SelectComboButtonOpts(
			widget.SelectComboButtonOpts.ComboButtonOpts(
				//Set the max height of the dropdown list
				widget.ComboButtonOpts.MaxContentHeight(300),
				//Set the parameters for the primary displayed button
				widget.ComboButtonOpts.ButtonOpts(
					widget.ButtonOpts.Image(buttonImage),
					widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
					widget.ButtonOpts.Text("", face, &widget.ButtonTextColor{
						Idle:     color.White,
						Disabled: color.White,
					}),
					//widget.ButtonOpts.WidgetOpts(
					//	//Set how wide the button should be
					//	widget.WidgetOpts.MinSize(150, 0),
					//	//Set the combobox's position
					//	widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
					//		HorizontalPosition: widget.AnchorLayoutPositionCenter,
					//		VerticalPosition:   widget.AnchorLayoutPositionCenter,
					//	}),
					//),
				),
			),
		),
		widget.ListComboButtonOpts.ListOpts(
			//Set how wide the dropdown list should be
			widget.ListOpts.ContainerOpts(widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.MinSize(150, 0))),
			//Set the entries in the list
			widget.ListOpts.Entries(entries),
			widget.ListOpts.ScrollContainerOpts(
				//Set the background images/color for the dropdown list
				widget.ScrollContainerOpts.Image(&widget.ScrollContainerImage{
					Idle:     image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
					Disabled: image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
					Mask:     image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
				}),
			),
			widget.ListOpts.SliderOpts(
				//Set the background images/color for the background of the slider track
				widget.SliderOpts.Images(&widget.SliderTrackImage{
					Idle:  image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
					Hover: image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
				}, buttonImage),
				widget.SliderOpts.MinHandleSize(5),
				//Set how wide the track should be
				widget.SliderOpts.TrackPadding(widget.NewInsetsSimple(2))),
			//Set the font for the list options
			widget.ListOpts.EntryFontFace(face),
			//Set the colors for the list
			widget.ListOpts.EntryColor(&widget.ListEntryColor{
				Selected:                   color.NRGBA{254, 255, 255, 255},             //Foreground color for the unfocused selected entry
				Unselected:                 color.NRGBA{254, 255, 255, 255},             //Foreground color for the unfocused unselected entry
				SelectedBackground:         color.NRGBA{R: 130, G: 130, B: 200, A: 255}, //Background color for the unfocused selected entry
				SelectedFocusedBackground:  color.NRGBA{R: 130, G: 130, B: 170, A: 255}, //Background color for the focused selected entry
				FocusedBackground:          color.NRGBA{R: 170, G: 170, B: 180, A: 255}, //Background color for the focused unselected entry
				DisabledUnselected:         color.NRGBA{100, 100, 100, 255},             //Foreground color for the disabled unselected entry
				DisabledSelected:           color.NRGBA{100, 100, 100, 255},             //Foreground color for the disabled selected entry
				DisabledSelectedBackground: color.NRGBA{100, 100, 100, 255},             //Background color for the disabled selected entry
			}),
			//Padding for each entry
			widget.ListOpts.EntryTextPadding(widget.NewInsetsSimple(5)),
		),
		//Define how the entry is displayed
		widget.ListComboButtonOpts.EntryLabelFunc(
			func(e any) string {
				//Button Label function
				return labelName + e.(ListEntry).name
			},
			func(e any) string {
				//List Label function
				return "" + e.(ListEntry).name
			}),
		//Callback when a new entry is selected
		widget.ListComboButtonOpts.EntrySelectedHandler(func(args *widget.ListComboButtonEntrySelectedEventArgs) {
			fmt.Println("Selected Entry: ", args.Entry)

			if comboType == "GameType" {
				gs.GameType = args.Entry.(ListEntry).name
			} else if comboType == "Preset" {
				gs.Preset = args.Entry.(ListEntry).name
			} else if comboType == "LtLRule" {
				gs.LtLRule = args.Entry.(ListEntry).name
			}

			gs.IsReset = true
		}),
	)

	initialEntry := ListEntry{}

	if comboType == "GameType" {
		initialEntry = ListEntry{
			id:   0,
			name: gs.GameType,
		}
	} else if comboType == "Preset" {
		initialEntry = ListEntry{
			id:   0,
			name: gs.Preset,
		}
	} else if comboType == "LtLRule" {
		initialEntry = ListEntry{
			id:   0,
			name: gs.LtLRule,
		}
	}

	combobox.SetSelectedEntry(initialEntry)

	return combobox
}

func getSlider(name string, gs *settings.GameSettings) *widget.Slider {
	slider := widget.NewSlider(
		// Set the slider orientation - n/s vs e/w
		widget.SliderOpts.Direction(widget.DirectionHorizontal),
		// Set the minimum and maximum value for the slider
		// todo: fix it
		widget.SliderOpts.MinMax(1, 150),

		widget.SliderOpts.WidgetOpts(
			// Set the Widget to layout in the center on the screen
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
			// Set the widget's dimensions
			widget.WidgetOpts.MinSize(200, 6),
		),
		widget.SliderOpts.Images(
			// Set the track images
			&widget.SliderTrackImage{
				Idle:  image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
				Hover: image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
			},
			// Set the handle images
			&widget.ButtonImage{
				Idle:    image.NewNineSliceColor(color.NRGBA{255, 100, 100, 255}),
				Hover:   image.NewNineSliceColor(color.NRGBA{255, 100, 100, 255}),
				Pressed: image.NewNineSliceColor(color.NRGBA{255, 100, 100, 255}),
			},
		),
		// Set the size of the handle
		widget.SliderOpts.FixedHandleSize(6),
		// Set the offset to display the track
		widget.SliderOpts.TrackOffset(0),
		// Set the size to move the handle
		widget.SliderOpts.PageSizeFunc(func() int {
			return 1
		}),
		// Set the callback to call when the slider value is changed
		widget.SliderOpts.ChangedHandler(func(args *widget.SliderChangedEventArgs) {
			speed := int(args.Current)
			gs.DesiredTPS = speed

			fmt.Println(args.Current)
		}),
	)

	// Set the current value of the slider
	slider.Current = gs.DesiredTPS

	return slider
}
