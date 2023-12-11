package ui

import (
	"fmt"
	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/tasnimzotder/artificial-life/constants"
	"github.com/tasnimzotder/artificial-life/settings"
	"log"
	"math"
	"time"
)

var (
	ScreenWidth         int = 640 * 1.75
	ScreenHeight        int = 480 * 1.75
	gridWidthPercentage     = 0.75
)

var (
	prevUpdateMillis int64 = 0
)

type Game struct {
	World        *settings.World
	pixels       []byte
	Settings     *settings.GameSettings
	UI           *ebitenui.UI
	ActualSpeeds []int64
	ActualSpeed  float64
}

func (g *Game) createInitialPixels() {
	//maxInitialLiveCells := int(float32(g.Settings.WorldWidth*g.Settings.WorldHeight) / 1.5)
	//maxInitialLiveCells := int(float32(g.Settings.WorldWidth*g.Settings.WorldHeight) / 2)
	//
	if g.Settings.GameType == constants.GAME_TYPE_GOL {
		if g.Settings.Preset == "Random" {
			g.World.InitRandom(g.Settings.GameType)
		} else {
			g.World.InitPreset(g.Settings.GameType, g.Settings.Preset)
		}
	} else if g.Settings.GameType == constants.GAME_TYPE_LTL || g.Settings.GameType == constants.GAME_TYPE_SMOOTH_LIFE {
		g.World.InitRandom(g.Settings.GameType)
	}

	//g.World.InitRandom(g.Settings.GameType)
}

func (g *Game) Update() error {
	if ebiten.IsFullscreen() {
		ScreenWidth, ScreenHeight = ebiten.ScreenSizeInFullscreen()
		gridWidthPercentage = 0.85
	} else {
		ScreenWidth, ScreenHeight = 640*1.5, 480*1.5
		gridWidthPercentage = 0.75
	}

	if g.isKeyJustPressed() {
		log.Printf("Key pressed")
	}

	if g.Settings.IsReset {
		log.Printf("Preset: %s", g.Settings.Preset)
		//g.World.ClearCells()
		g.createInitialPixels()
		g.Settings.ResetSettings()

		if g.Settings.Replay {
			g.Settings.IsPaused = false
			g.Settings.Replay = false
		}
	}

	tps := g.Settings.DesiredTPS

	if tps < 60 {
		tps = 60
	} else {
		tps = int(float32(tps) * 2)
	}

	//ebiten.SetTPS(g.Settings.DesiredTPS)
	ebiten.SetTPS(tps)

	if !g.Settings.IsPaused {
		currentMillis := time.Now().UnixNano() / 1000000

		if (currentMillis - prevUpdateMillis) >= 1000/int64(g.Settings.DesiredTPS) {
			g.nextGeneration()
			//delay := time.Since(time.Unix(0, prevUpdateMillis*1000000)).Milliseconds()
			//fmt.Printf("Delay: %d\n", delay)
			duration := time.Since(time.Unix(0, prevUpdateMillis*1000000))
			g.ActualSpeeds = append(g.ActualSpeeds, duration.Milliseconds())

			prevUpdateMillis = currentMillis
		}

	}

	g.UI.Update()
	return nil
}

func (g *Game) nextGeneration() {
	g.Settings.T += 1
	g.World.NextGeneration()
}

func (g *Game) isKeyJustPressed() bool {
	// zoom in
	if inpututil.IsKeyJustPressed(ebiten.KeyPeriod) {
		g.Settings.ZoomLevel += 1
		return true
	}

	// zoom out
	if inpututil.IsKeyJustPressed(ebiten.KeyComma) {
		g.Settings.ZoomLevel -= 1

		if g.Settings.ZoomLevel < 1 {
			g.Settings.ZoomLevel = 1
		}

		return true
	}

	// play/pause
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.Settings.IsPaused = !g.Settings.IsPaused
		return true
	}

	// reset
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		//g.Settings = settings.NewGameSettings()
		g.Settings.IsReset = true
		return true
	}

	// toggle fullscreen
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
		return true
	}

	// reset and play
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.Settings.IsReset = true
		g.Settings.Replay = true
		return true
	}

	// next generation - single step
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.nextGeneration()
		return true
	}

	return false
}

func (g *Game) Draw(screen *ebiten.Image) {
	// todo: shift the image/component variable to the struct
	//screen.Fill(&color.RGBA{A: 0xff})

	newScreenWidth := float64(ScreenWidth) * gridWidthPercentage
	newScreenHeight := float64(ScreenHeight)

	widthRatio := newScreenWidth / float64(g.Settings.WorldWidth)
	heightRatio := float64(newScreenHeight) / float64(g.Settings.WorldHeight)

	widthRatioWithZoom := widthRatio * float64(g.Settings.ZoomLevel)
	heightRatioWithZoom := heightRatio * float64(g.Settings.ZoomLevel)

	visibleWidth := math.Round(newScreenWidth / widthRatioWithZoom)
	visibleHeight := math.Round(float64(newScreenHeight) / heightRatioWithZoom)

	if visibleWidth > float64(ScreenWidth) {
		visibleWidth = float64(ScreenWidth)
	}

	if visibleHeight > float64(ScreenHeight) {
		visibleHeight = float64(ScreenHeight)
	}

	grid := ebiten.NewImage(int(visibleWidth), int(visibleHeight))

	g.World.Draw(grid, int(visibleWidth), int(visibleHeight))

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(widthRatioWithZoom, heightRatioWithZoom)
	// translate the image to the center of the screen
	//op.GeoM.Translate(ScreenWidth/2-visibleWidth*widthRatioWithZoom/2, newScreenHeight/2-visibleHeight*widthRatioWithZoom/2)
	op.Filter = ebiten.FilterNearest

	ctrlPanelWidth := 200
	ctrlPanelHeight := 200
	controlPanel := ebiten.NewImage(ctrlPanelWidth, ctrlPanelHeight)

	opCtrlPanel := &ebiten.DrawImageOptions{}

	opCtrlPanel.GeoM.Scale(1, 1)
	opCtrlPanel.Filter = ebiten.FilterNearest
	opCtrlPanel.GeoM.Translate(newScreenWidth, 50)
	opCtrlPanel.Filter = ebiten.FilterNearest

	// write text on control panel
	text := fmt.Sprintf(
		"T: %d\nUPS: %0.2f\nTPS: %0.2f\nFPS: %0.2f\nTotal Cells: %d\nCurrent Cells: %d\nZoom Level: %d",
		g.Settings.T,
		g.ActualSpeed,
		ebiten.ActualTPS(),
		ebiten.ActualFPS(),
		g.World.GoLWorld.GetLength(),
		grid.Bounds().Dx()*grid.Bounds().Dy(),
		g.Settings.ZoomLevel,
	)
	ebitenutil.DebugPrintAt(controlPanel, text, 10, 5)

	screen.DrawImage(grid, op)
	screen.DrawImage(controlPanel, opCtrlPanel)

	g.UI.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {

	if ebiten.IsFullscreen() {
		return ebiten.ScreenSizeInFullscreen()
	}

	return ScreenWidth, ScreenHeight
}

func (g *Game) UpdateActualSpeed() {
	go func() {
		for range time.Tick(1 * time.Second) {
			speeds := g.ActualSpeeds

			g.ActualSpeeds = make([]int64, 0)
			var val int64 = 0

			for _, v := range speeds {
				val += v
			}

			if len(speeds) == 0 {
				g.ActualSpeed = 0
				continue
			}

			fps := 1000.0 / float64(val) * float64(len(speeds))
			g.ActualSpeed = fps

			//g.ActualSpeed = float64(val) / float64(len(speeds))
		}
	}()
}
