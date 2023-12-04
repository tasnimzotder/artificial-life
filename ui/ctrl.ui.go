package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/tasnimzotder/artificial-life/settings"
	"github.com/tasnimzotder/artificial-life/utils"
)

type UIController struct {
	gs      *settings.GameSettings
	Content *fyne.Container
	Buttons struct {
		LoadButton      *widget.Button
		StartStopButton *widget.Button
		ResetButton     *widget.Button
	}
	Selectors struct {
		GameTypeSelector *widget.Select
		PresetSelector   *widget.Select
		RuleSelector     *widget.Select
	}
	Sliders struct {
		SpeedSlider *widget.Slider
		//ZoomSlider  *widget.Slider
	}
}

func (uic *UIController) NewUIController(gs *settings.GameSettings) {
	uic.gs = gs

	buttons := uic.ButtonsContainer()
	selectors := uic.SelectorsContainer()
	sliders := uic.SlidersContainer()

	uic.Content = container.NewHBox(
		container.NewVBox(
			selectors,
			buttons,
		),
		sliders,
	)

	//uic.Content.Resize(fyne.NewSize(900, 800))
}

func (uic *UIController) ButtonsContainer() *fyne.Container {
	// load button
	loadButton := widget.NewButton("Load", func() {})
	loadButton.OnTapped = func() {
		uic.gs.States.Paused = true
		uic.gs.States.Loaded = false
		uic.Buttons.StartStopButton.SetText("Start")

		//	todo: implement
	}

	uic.Buttons.LoadButton = loadButton

	// start/stop button
	startStopButton := widget.NewButton("Start", func() {})
	startStopButton.OnTapped = func() {
		//uic.gs.States.Running = !uic.gs.States.Running
		uic.gs.States.Paused = !uic.gs.States.Paused

		if uic.gs.States.Paused {
			startStopButton.SetText("Start")
		} else {
			startStopButton.SetText("Stop")
		}
	}

	//resetButton := widget.NewButton("Reset", func() {})
	//resetButton.OnTapped = func() {
	//	uic.gs.States.Paused = true
	//	startStopButton.SetText("Start")
	//}

	uic.Buttons.StartStopButton = startStopButton

	buttonContainer := container.NewHBox(
		loadButton,
		startStopButton,
	)

	return buttonContainer
}

func (uic *UIController) SelectorsContainer() *fyne.Container {
	// game type selector
	gameTypeSelector := widget.NewSelect(uic.gs.GameTypes, func(s string) {})
	gameTypeSelector.SetSelected(uic.gs.GameType)
	gameTypeSelector.OnChanged = func(s string) {
		uic.gs.GameType = s
		uic.Selectors.PresetSelector.Options = uic.gs.GoLPresets[s]
		uic.Selectors.PresetSelector.SetSelected(uic.gs.GoLPresets[s][0])

		if s == "LtL" {
			uic.Selectors.RuleSelector.Show()
		} else {
			uic.Selectors.RuleSelector.Hide()
		}

		uic.gs.States.Loaded = false
	}

	uic.Selectors.GameTypeSelector = gameTypeSelector
	gameTypeSelectorLabel := widget.NewLabel("Game Type: ")

	gameTypeSelectorContainer := container.NewVBox(
		gameTypeSelectorLabel,
		gameTypeSelector,
	)

	// preset selector
	presetSelector := widget.NewSelect(uic.gs.GoLPresets[uic.gs.GameType], func(s string) {})
	presetSelector.SetSelected(uic.gs.Preset)
	presetSelector.OnChanged = func(s string) {
		uic.gs.Preset = s
		uic.gs.States.Loaded = false
	}

	uic.Selectors.PresetSelector = presetSelector
	presetSelectorLabel := widget.NewLabel("Preset: ")

	presetSelectorContainer := container.NewVBox(
		presetSelectorLabel,
		presetSelector,
	)

	// rule selector

	ruleSelector := widget.NewSelect(settings.GetLtLRuleNames(), func(s string) {})
	ruleSelector.OnChanged = func(s string) {
		uic.gs.LtLRule = s
		//	todo: implement

		uic.gs.States.Paused = true
		uic.gs.States.Loaded = false
	}

	if uic.gs.GameType == "LtL" {
		ruleSelector.Show()
	} else {
		ruleSelector.Hide()
	}

	uic.Selectors.RuleSelector = ruleSelector

	ruleSelectorLabel := widget.NewLabel("Rule: ")

	ruleSelectorContainer := container.NewVBox(
		ruleSelectorLabel,
		ruleSelector,
	)

	// final container

	selectorsContainer := container.NewHBox(
		gameTypeSelectorContainer,
		presetSelectorContainer,
		ruleSelectorContainer,
	)

	return selectorsContainer
}

func (uic *UIController) SlidersContainer() *fyne.Container {
	// speed slider - generations per second
	speedSliderLabel := widget.NewLabel("Speed: ")
	speedSliderValueLabel := widget.NewLabel(fmt.Sprintf("x%d", uic.gs.GPS))

	speedSlider := widget.NewSlider(1, 64)
	speedSlider.SetValue(float64(uic.gs.GPS))
	speedSlider.OnChanged = func(f float64) {
		uic.gs.GPS = uint8(f)
		uic.gs.LogicLoopDelay = utils.FPSToMilliseconds(int64(uic.gs.GPS))

		speedSliderValueLabel.SetText(fmt.Sprintf("x%d", uic.gs.GPS))
	}

	speedSlider.Step = 1
	speedSlider.Resize(fyne.NewSize(100, 20))

	speedLessButton := widget.NewButton("-", func() {})
	speedLessButton.OnTapped = func() {
		if speedSlider.Value > 1 {
			speedSlider.SetValue(speedSlider.Value - 1)
		}
	}

	speedMoreButton := widget.NewButton("+", func() {})
	speedMoreButton.OnTapped = func() {
		if speedSlider.Value < 64 {
			speedSlider.SetValue(speedSlider.Value + 1)
		}
	}

	uic.Sliders.SpeedSlider = speedSlider

	speedSliderContainer := container.NewHBox(
		speedSliderLabel,
		speedSliderValueLabel,
		speedLessButton,
		speedSlider,
		speedMoreButton,
	)

	// zoom slider
	//zoomSliderLabel := widget.NewLabel("Zoom: ")
	//zoomSliderValueLabel := widget.NewLabel(fmt.Sprintf("x%d", uic.gs.ZoomLevel))
	//
	//zoomSlider := widget.NewSlider(10, 100)
	//zoomSlider.SetValue(float64(uic.gs.ZoomLevel))
	//zoomSlider.OnChanged = func(f float64) {
	//	uic.gs.ZoomLevel = int(f)
	//	zoomSliderValueLabel.SetText(fmt.Sprintf("x%d", uic.gs.ZoomLevel))
	//
	//}
	//
	//zoomSlider.Step = 1
	//zoomSlider.Resize(fyne.NewSize(100, 20))
	//
	//zoomLessButton := widget.NewButton("-", func() {
	//	if zoomSlider.Value > 10 {
	//		zoomSlider.SetValue(zoomSlider.Value - 1)
	//	}
	//})
	//
	//zoomMoreButton := widget.NewButton("+", func() {
	//	if zoomSlider.Value < 100 {
	//		zoomSlider.SetValue(zoomSlider.Value + 1)
	//	}
	//})
	//
	//uic.Sliders.ZoomSlider = zoomSlider
	//
	//zoomSliderContainer := container.NewHBox(
	//	zoomSliderLabel,
	//	zoomSliderValueLabel,
	//	zoomLessButton,
	//	zoomSlider,
	//	zoomMoreButton,
	//)

	//	 final container
	slidersContainer := container.NewVBox(
		speedSliderContainer,
		//zoomSliderContainer,
	)

	return slidersContainer
}
