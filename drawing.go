package main

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"github.com/jdxyw/generativeart/common"
)

var DRAWINGS = map[string]generativeart.Engine{
	"colorline":  arts.NewContourLine(400),
	"circlegrid": arts.NewCircleGrid(4, 6),
	"ciclemove":  arts.NewCircleMove(1000),
	"blackhole":  arts.NewBlackHole(500, 600, 0.01),
	"janus":      arts.NewJanus(4, 10),
	"random":     arts.NewRandomShape(150),
	"colorcanve": arts.NewColorCanve(5),
	"circles":    arts.NewColorCircle2(7),
	"oceanfish":  arts.NewOceanFish(100, 6),
	"yarn":       arts.NewYarn(1000),
	"silksmoke":  arts.NewSilkSmoke(10, 20, 3, 20, 10, 80, true),
	"dotswave":   arts.NewDotsWave(400),
}

func DrawMany(drawings map[string]generativeart.Engine) {
	for k := range drawings {
		DrawOne(k)
	}
}

func DrawOne(art string) string {
	c := generativeart.NewCanva(400, 400)
	c.SetColorSchema([]color.RGBA{
		{22, 50, 91, 0xFF},
		{34, 123, 148, 0xFF},
		{120, 183, 208, 0xFF},
		{255, 220, 127, 0xFF},
	})

	c.SetBackground(common.LightGray)
	c.FillBackground()
	c.SetLineWidth(0.5)
	c.SetLineColor(common.Moccasin)
	c.Draw(DRAWINGS[art])
	fileName := fmt.Sprintf("./results/%s_%v.png", art, rand.Float64())
	c.ToPNG(fileName)
	return fileName
}
