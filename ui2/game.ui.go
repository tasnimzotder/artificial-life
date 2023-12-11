package ui2

import (
	"fmt"
	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/tasnimzotder/artificial-life/constants"
	"github.com/tasnimzotder/artificial-life/settings"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"
)

var (
	ScreenWidth         int = 640 * 1.5
	ScreenHeight        int = 480 * 1.5
	gridWidthPercentage     = 0.75
)

var (
	prevUpdateMillis int64 = 0
)

type Game struct {
	World    *settings.World
	pixels   []byte
	Settings *settings.GameSettings2
	UI       *ebitenui.UI
}

func (g *Game) createInitialPixels() {
	maxInitialLiveCells := int(float32(g.Settings.WorldWidth*g.Settings.WorldHeight) / 1.5)

	if g.Settings.GameType == constants.GAME_TYPE_GOL {
		if g.Settings.Preset == "Random" {
			g.World.InitRandom(maxInitialLiveCells)
		} else {
			g.World.InitPreset(g.Settings.GameType, g.Settings.Preset)
		}
	} else if g.Settings.GameType == constants.GAME_TYPE_LTL {
		g.World.InitRandom(maxInitialLiveCells)
	}
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
		g.World.ClearCells()
		g.createInitialPixels()
		g.Settings.ResetSettings()
	}

	tps := g.Settings.DesiredTPS

	if tps < 60 {
		tps = 60
	} else {
		tps = int(float32(tps) * 1.25)
	}

	//ebiten.SetTPS(g.Settings.DesiredTPS)
	//ebiten.SetTPS(64)

	if !g.Settings.IsPaused {
		currentMillis := time.Now().UnixNano() / 1000000

		if (currentMillis - prevUpdateMillis) >= 1000/int64(g.Settings.DesiredTPS) {
			g.World.NextGeneration(g.Settings)

			//delay := time.Since(time.Unix(0, prevUpdateMillis*1000000)).Milliseconds()
			//fmt.Printf("Delay: %d\n", delay)

			prevUpdateMillis = currentMillis
		}

	}

	g.UI.Update()
	return nil
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

	return false
}

func (g *Game) Draw(screen *ebiten.Image) {
	// todo: shift the image/component variable to the struct
	screen.Fill(&color.RGBA{A: 0xff})

	newScreenWidth := float64(ScreenWidth) * gridWidthPercentage
	newScreenHeight := float64(ScreenHeight)

	widthRatio := newScreenWidth / float64(g.Settings.WorldWidth)
	heightRatio := float64(newScreenHeight) / float64(g.Settings.WorldHeight)

	widthRatioWithZoom := widthRatio * float64(g.Settings.ZoomLevel)
	heightRatioWithZoom := heightRatio * float64(g.Settings.ZoomLevel)

	// only draw the visible area of the world
	visibleWidth := math.Round(newScreenWidth / widthRatioWithZoom)
	visibleHeight := math.Round(float64(newScreenHeight) / heightRatioWithZoom)

	if visibleWidth > float64(g.Settings.WorldWidth) {
		visibleWidth = float64(g.Settings.WorldWidth)
	}

	if visibleHeight > float64(g.Settings.WorldHeight) {
		visibleHeight = float64(g.Settings.WorldHeight)
	}

	//grid := ebiten.NewImage(ScreenWidth, newScreenHeight)
	grid := ebiten.NewImage(int(visibleWidth), int(visibleHeight))

	cX := int(math.Round(float64(g.Settings.WorldWidth) / 2))
	cY := int(math.Round(float64(g.Settings.WorldHeight) / 2))

	for y := cY - int(math.Round(float64(visibleHeight)/2)); y < cY+int(math.Floor(float64(visibleHeight)/2)); y++ {
		for x := cX - int(math.Round(float64(visibleWidth)/2)); x < cX+int(math.Floor(float64(visibleWidth)/2)); x++ {
			yGrid := y - (cY - int(math.Round(float64(visibleHeight)/2)))
			xGrid := x - (cX - int(math.Round(float64(visibleWidth)/2)))

			if x >= g.Settings.WorldWidth || y >= g.Settings.WorldHeight {
				//grid.Set(x, y, color.RGBA{100, 0, 0, 0xff})
				grid.Set(xGrid, yGrid, color.RGBA{100, 0, 0, 0xff})

				continue
			}

			if x < 0 || y < 0 {
				//grid.Set(x, y, color.RGBA{0, 0, 100, 0xff})
				grid.Set(xGrid, yGrid, color.RGBA{0, 0, 100, 0xff})
				continue
			}

			if g.World.GetCell(x, y) != 0 {
				bc := uint(rand.Intn(50))
				//grid.Set(x, y, color.RGBA{uint8(255 - bc), 255, uint8(255 - bc), 0xff})
				grid.Set(xGrid, yGrid, color.RGBA{uint8(255 - bc), 255, uint8(255 - bc), 0xff})
			} else {
				//grid.Set(x, y, color.RGBA{100, 130, 160, 0xff})
				grid.Set(xGrid, yGrid, color.RGBA{50, 60, 70, 0xff})
			}
		}
	}

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(widthRatioWithZoom, widthRatioWithZoom)
	// translate the image to the center of the screen
	//op.GeoM.Translate(ScreenWidth/2-visibleWidth*widthRatioWithZoom/2, newScreenHeight/2-visibleHeight*widthRatioWithZoom/2)
	op.Filter = ebiten.FilterNearest

	ctrlPanelWidth := 200
	ctrlPanelHeight := 200
	controlPanel := ebiten.NewImage(ctrlPanelWidth, ctrlPanelHeight)
	//controlPanel.Fill(color.RGBA{A: 0xff})

	opCtrlPanel := &ebiten.DrawImageOptions{}

	opCtrlPanel.GeoM.Scale(1, 1)
	opCtrlPanel.Filter = ebiten.FilterNearest
	opCtrlPanel.GeoM.Translate(newScreenWidth, 50)
	opCtrlPanel.Filter = ebiten.FilterNearest

	// write text on control panel
	text := fmt.Sprintf(
		"TPS: %0.2f\nFPS: %0.2f\nTotal Cells: %d\nCurrent Cells: %d\nZoom Level: %d",
		ebiten.ActualTPS(),
		ebiten.ActualFPS(),
		len(g.World.Area[0])*len(g.World.Area),
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
