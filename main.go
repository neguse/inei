package main

import (
	_ "embed"
	"log"
	"syscall/js"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed public/shader.kage
var shaderSrc string

var (
	startTime = time.Now()
)

type Game struct {
	shader *ebiten.Shader
}

func (g *Game) ApplyShader(code string) error {
	newShader, err := ebiten.NewShader([]byte(code))
	if err != nil {
		return err
	}
	g.shader = newShader
	return nil
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawRectShaderOptions{}
	cx, cy := ebiten.CursorPosition()

	op.Uniforms = map[string]interface{}{
		"Time":   float32(time.Since(startTime).Seconds()),
		"Cursor": []float32{float32(cx), float32(cy)},
	}

	screen.DrawRectShader(screen.Bounds().Dx(), screen.Bounds().Dy(), g.shader, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	game := &Game{}
	game.ApplyShader(shaderSrc)

	applyShader := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			log.Println("Error: no shader code provided")
			return nil
		}
		code := args[0].String()
		log.Println("Applying shader from editor...")
		if err := game.ApplyShader(code); err != nil {
			log.Printf("Shader compile error: %v\n", err)
			return err.Error()
		}
		return nil
	})
	js.Global().Set("applyShader", applyShader)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
