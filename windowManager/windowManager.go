package windowManager

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/tasnimzotder/artificial-life/services"
	"github.com/tasnimzotder/artificial-life/utils"
	"time"
)

type GameWindow struct {
	Content *fyne.Container
	Widgets struct {
		FPSWidget *widget.Label
		T0Widget  *widget.Label
	}
	Buttons struct {
		StartStopButton *widget.Button
		ResetButton     *widget.Button
	}
	Sliders struct {
		FPS *widget.Slider
	}
	Selectors struct {
		GameTypeSelector *widget.Select
		PresetSelector   *widget.Select
	}
	CheckBoxes struct {
		WrapAroundCheckbox *widget.Check
	}
}

func (gw *GameWindow) New(gs *utils.GameSettings) *fyne.Container {
	gw.Widgets.FPSWidget = widget.NewLabel("0")
	gw.Widgets.FPSWidget.SetText("FPS: 60")

	gw.Buttons.StartStopButton = widget.NewButton("Start/Pause", func() {})
	gw.Buttons.ResetButton = widget.NewButton("Reset", func() {})

	gw.Sliders.FPS = widget.NewSlider(1, 150)
	gw.Sliders.FPS.SetValue(2)

	gw.Selectors.GameTypeSelector = widget.NewSelect(gs.GameTypes, func(s string) {
		gs.GameType = s
	})
	gw.Selectors.GameTypeSelector.SetSelected("GoL")
	gameTypeSelectorLabel := widget.NewLabel("Game Type")

	gameTypeSelectorContainer := container.NewVBox(gameTypeSelectorLabel, gw.Selectors.GameTypeSelector)

	gw.Selectors.PresetSelector = widget.NewSelect(gs.Presets[gs.GameType], func(s string) {
		gs.Preset = s
	})
	gw.Selectors.PresetSelector.SetSelected("Random")
	presetSelectorLabel := widget.NewLabel("Preset")

	presetSelectorContainer := container.NewVBox(presetSelectorLabel, gw.Selectors.PresetSelector)

	gw.CheckBoxes.WrapAroundCheckbox = widget.NewCheck("Wrap Around", func(b bool) {
		gs.WrapAround = b
	})
	gw.CheckBoxes.WrapAroundCheckbox.SetChecked(true)

	//beginButton := widget.NewButton("Begin", func() {
	//	GameRunner(gameSettings, startStopButton, resetButton, fps, fpsSlider, presetSelector, gameTypeSelector, wrapAroundCheckbox)
	//})

	gw.Widgets.T0Widget = widget.NewLabel("0")
	gw.Widgets.T0Widget.SetText("Generation: 0")

	T0Container := container.NewHBox(gw.Widgets.T0Widget)

	ctrlContainer := container.NewVBox(gw.Widgets.FPSWidget, container.NewHBox(gw.Buttons.StartStopButton, gw.Buttons.ResetButton, gw.CheckBoxes.WrapAroundCheckbox, T0Container))
	buttons := container.NewHBox(gameTypeSelectorContainer, presetSelectorContainer, ctrlContainer)
	fpsContainer := container.NewVBox(widget.NewLabel("FPS"), gw.Sliders.FPS)

	gw.Content = container.NewVBox(fpsContainer, buttons)

	//go routine to update updating user input
	gw.UserInputListener(gs)

	return gw.Content
}

func (gw *GameWindow) UserInputListener(gs *utils.GameSettings) {
	go func() {
		prevTime := time.Now()

		for {
			gw.Buttons.StartStopButton.OnTapped = func() {
				fmt.Println("startStopButton.OnTapped")
				gs.IsPaused = !gs.IsPaused

				if !gs.IsPaused && gs.IsReset {
					gs.IsReset = false
				}
			}

			gw.Buttons.ResetButton.OnTapped = func() {
				fmt.Println("resetButton.OnTapped")
				gs.IsReset = true
				gs.FPS = 2
				gw.Sliders.FPS.SetValue(2)

				//utils.HandleReset(gs)
				services.GameRunner(gs)
			}

			gw.Sliders.FPS.OnChanged = func(value float64) {
				fmt.Println("fpsSlider.OnChanged")
				gs.FPS = int(value)
			}

			gw.Selectors.GameTypeSelector.OnChanged = func(s string) {
				fmt.Printf("gameSelector.OnChanged: %s\n", s)
				gs.GameType = s

				gw.Selectors.PresetSelector.Options = gs.Presets[gs.GameType]
				gw.Selectors.PresetSelector.Refresh()
			}

			gw.Selectors.PresetSelector.OnChanged = func(s string) {
				fmt.Printf("presetSelector.OnChanged: %s\n", s)
				gs.Preset = s

				services.GameRunner(gs)
			}

			gw.CheckBoxes.WrapAroundCheckbox.OnChanged = func(b bool) {
				fmt.Printf("wrapAroundCheckbox.OnChanged: %t\n", b)
				gs.WrapAround = b
			}

			currTime := time.Now()
			elapsed := currTime.Sub(prevTime).Milliseconds()

			if elapsed > 250 {
				// update T0
				t0Str := fmt.Sprintf("Generation: %d", gs.Parameters.T)
				gw.Widgets.T0Widget.SetText(t0Str)

				//	update FPS
				fpsStr := fmt.Sprintf("Desired FPS: %d\tCurrent FPS: %.2f", gs.FPS, gs.CurrentFPS)
				gw.Widgets.FPSWidget.SetText(fpsStr)

				prevTime = currTime
			}
		}
	}()
}
