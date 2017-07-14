package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320 //размер канвы в пикселях
	cells         = 100  //Количество ячеек в сетки
	xyrange       = 30.0     //Диапозон осей
	xyscale       = width / 2 / xyrange //Пикселей в единице x или y
	zscale        = height * 0.4 //  Пикселей в единице z
	angle         = math.Pi / 6 //Углы осей x, y (=30 градусов)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Printf("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := hill(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale + z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	if r == 0 {
		return 0
	}
 	sin := math.Sin(r) / r
	return sin

}


func hill(x, y float64) float64 {
	r := math.Hypot((math.Pi*x)/100, (math.Pi*y)/100)
	if r == 0{
		return 0
	}
	sin := math.Sin(r)
	return sin
}
