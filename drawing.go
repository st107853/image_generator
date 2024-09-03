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
	"maze": arts.NewMaze(10),
	"julia": arts.NewJulia(func(z complex128) complex128 {
		return z*z +
			complex(-0.1, 0.651)
	}, 40, 1.5, 1.5),
	"randcicle": arts.NewRandCicle(15, 20, 1, 5, 10, 40, true),
	"blackhole": arts.NewBlackHole(500, 600, 0.01),
	"janus":     arts.NewJanus(4, 10),
	"random":    arts.NewRandomShape(150),
	"silksky":   arts.NewSilkSky(5, 10),
	"circles":   arts.NewColorCircle2(30),
	"oceanfish": arts.NewOceanFish(100, 150),
	"pixelhole": arts.NewPixelHole(200),
	"silksmoke": arts.NewSilkSmoke(10, 20, 3, 20, 10, 80, true),
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
	c.SetLineWidth(1.0)
	c.SetLineColor(common.Moccasin)
	c.Draw(DRAWINGS[art])
	fileName := fmt.Sprintf("./results/%s_%v.png", art, rand.Float64())
	c.ToPNG(fileName)
	return fileName
}
