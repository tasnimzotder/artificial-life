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

func GameWindow(gs *utils.GameSettings) *fyne.Container {
	fpsWidget := widget.NewLabel("0")
	fpsWidget.SetText("FPS: 60")

	startStopButton := widget.NewButton("Start/Pause", func() {})
	resetButton := widget.NewButton("Reset", func() {})

	fpsSlider := widget.NewSlider(1, 100)
	fpsSlider.SetValue(2)

	gameTypeSelector := widget.NewSelect(gs.GameTypes, func(s string) {
		gs.GameType = s
	})
	gameTypeSelector.SetSelected("GoL")
	gameTypeSelectorLabel := widget.NewLabel("Game Type")

	gameTypeSelectorContainer := container.NewVBox(gameTypeSelectorLabel, gameTypeSelector)

	presetSelector := widget.NewSelect(gs.Presets[gs.GameType], func(s string) {
		gs.Preset = s
	})
	presetSelector.SetSelected("Random")
	presetSelectorLabel := widget.NewLabel("Preset")

	presetSelectorContainer := container.NewVBox(presetSelectorLabel, presetSelector)

	fpsSliderContainer := container.NewVBox(widget.NewLabel("FPS"), fpsSlider)

	wrapAroundCheckbox := widget.NewCheck("Wrap Around", func(b bool) {
		gs.WrapAround = b
	})
	wrapAroundCheckbox.SetChecked(true)

	//beginButton := widget.NewButton("Begin", func() {
	//	GameRunner(gameSettings, startStopButton, resetButton, fps, fpsSlider, presetSelector, gameTypeSelector, wrapAroundCheckbox)
	//})

	ctrlContainer := container.NewVBox(fpsWidget, container.NewHBox(startStopButton, resetButton, wrapAroundCheckbox))

	buttons := container.NewHBox(gameTypeSelectorContainer, presetSelectorContainer, ctrlContainer)
	content := container.NewVBox(fpsSliderContainer, buttons)

	//go routine to update updating user input
	go func() {
		for {
			startStopButton.OnTapped = func() {
				fmt.Println("startStopButton.OnTapped")
				gs.IsPaused = !gs.IsPaused

				if !gs.IsPaused && gs.IsReset {
					gs.IsReset = false
				}
			}

			resetButton.OnTapped = func() {
				fmt.Println("resetButton.OnTapped")
				gs.IsReset = true
				gs.FPS = 2
				fpsSlider.SetValue(2)

				//utils.HandleReset(gs)
				services.GameRunner(gs)
			}

			fpsSlider.OnChanged = func(value float64) {
				fmt.Println("fpsSlider.OnChanged")
				gs.FPS = int(value)
			}

			gameTypeSelector.OnChanged = func(s string) {
				fmt.Printf("gameSelector.OnChanged: %s\n", s)
				gs.GameType = s

				presetSelector.Options = gs.Presets[gs.GameType]
				presetSelector.Refresh()
			}

			presetSelector.OnChanged = func(s string) {
				fmt.Printf("presetSelector.OnChanged: %s\n", s)
				gs.Preset = s

				services.GameRunner(gs)
			}

			wrapAroundCheckbox.OnChanged = func(b bool) {
				fmt.Printf("wrapAroundCheckbox.OnChanged: %t\n", b)
				gs.WrapAround = b
			}

			//	update FPS in evert 500ms
			for range time.Tick(500 * time.Millisecond) {
				fpsWidget.SetText(fmt.Sprintf("FPS: %d", gs.CurrentFPS))
			}
		}
	}()

	return content

}
