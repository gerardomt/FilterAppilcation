package test

import ("testing"
	"filter"
	"math/rand"
	"image/color"
	_ "image/png"
)

var validImage = "media/fotitus.jpg"
var noValidImage = "media/testpng.png"
var nPixelTest = 300

func TestNewFilter(t *testing.T){
	_, err := filter.NewFilter(noValidImage)
	if err == nil {
		t.Errorf("NewFilter did not return an error")
	}
}

func TestSetImage(t *testing.T){
	ft, _ := filter.NewFilter(validImage)
	err := ft.SetImage(noValidImage)
	if err == nil {
		t.Errorf("SetImage did not return an error")
	}
}

func TestSaveImage(t *testing.T){
	ft, _ := filter.NewFilter(validImage)
	_, err := ft.SaveImage("")
	if err == nil {
		t.Errorf("SaveImage did not return an error")
	}
}

func TestRedFilter(t *testing.T){
	ft, _ := filter.NewFilter(validImage)
	err := ft.RedFilter()
	if err != nil{
		t.Errorf("%s", err)
		return
	}
	var x, y int
	var pixel color.RGBA
	for i:=0; i<nPixelTest; i++{
		x = rand.Intn(ft.Size.X)
		y = rand.Intn(ft.Size.Y)
		pixel = color.RGBAModel.Convert(ft.Buffer.At(x,y)).(color.RGBA)

		if pixel.B!=0 || pixel.G!=0 {
			t.Errorf("Pixel at %d, %d is not a red pixel", x, y)
			break
		}
	}
}

func TestBlueFilter(t *testing.T){
	ft, _ := filter.NewFilter(validImage)
	err := ft.BlueFilter()
	if err != nil{
		t.Errorf("%s", err)
		return
	}
	var x, y int
	var pixel color.RGBA
	for i:=0; i<nPixelTest; i++{
		x = rand.Intn(ft.Size.X)
		y = rand.Intn(ft.Size.Y)
		pixel = color.RGBAModel.Convert(ft.Buffer.At(x,y)).(color.RGBA)

		if pixel.R!=0 || pixel.G!=0 {
			t.Errorf("Pixel at %d, %d is not a blue pixel", x, y)
			break
		}
	}	
}

func TestGreenFilter(t *testing.T){
	ft, _ := filter.NewFilter(validImage)
	err := ft.GreenFilter()
	if err != nil{
		t.Errorf("%s", err)
		return
	}
	var x, y int
	var pixel color.RGBA
	for i:=0; i<nPixelTest; i++{
		x = rand.Intn(ft.Size.X)
		y = rand.Intn(ft.Size.Y)
		pixel = color.RGBAModel.Convert(ft.Buffer.At(x,y)).(color.RGBA)

		if pixel.B!=0 || pixel.R!=0 {
			t.Errorf("Pixel at %d, %d is not a green pixel", x, y)
			break
		}
	}	
}

func TestPixelFilter(t *testing.T){
	bigPxSize := 30
	var testPixel, pixel color.RGBA
	ft, _ := filter.NewFilter(validImage)
	err := ft.PixelFilter(bigPxSize)
		if err != nil{
		t.Errorf("%s", err)
		return
	}

	for x:=0; x<ft.Size.X-bigPxSize; x=x+bigPxSize {
		for y:=0; y<ft.Size.Y-bigPxSize; y=y+bigPxSize {
			testPixel = color.RGBAModel.Convert(
				ft.Buffer.At(x,y)).(color.RGBA)
			for a:=0;a<bigPxSize;a++ {
				for b:=0;b<bigPxSize;b++ {
					pixel = color.RGBAModel.Convert(
						ft.Buffer.At(x+a,y+b)).(color.RGBA)
					if (testPixel.R!=pixel.R || testPixel.B!=pixel.B || testPixel.G!=pixel.G){
						t.Errorf("Pixel at %d, %d is different from test pixel at %d, %d", x+a, y+b, x, y)
						return
					}
				}
			}
		}
	}
}
