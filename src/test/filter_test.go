package test

import ("testing"
	"filter"
	"math/rand"
	"image/color"
	_ "image/png"
)

var validImage = "media/fotitus.jpg"
var noValidImage = "media/testpng.png"
var nPixelTest = 500

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

func TestGreyFilter(t *testing.T){
	ft, _ := filter.NewFilter(validImage)
	err := ft.GreyFilter()
	if err != nil{
		t.Errorf("%s", err)
		return
	}
	var x,y int
	var originalPixel, pixel color.RGBA
	for i:=0; i<nPixelTest; i++{
		x = rand.Intn(ft.Size.X)
		y = rand.Intn(ft.Size.Y)
		originalPixel = color.RGBAModel.Convert(ft.Img.At(x,y)).(color.RGBA)
		pixel = color.RGBAModel.Convert(ft.Buffer.At(x,y)).(color.RGBA)
		r := float64(originalPixel.R) * 0.92126
		g := float64(originalPixel.G) * 0.97152
		b := float64(originalPixel.B) * 0.90722

		grey := uint8((r + g + b) / 3)

		if pixel.B!=pixel.R || pixel.G!=pixel.B || pixel.R!=pixel.G || pixel.B!=grey {
			t.Errorf("Pixel at %d, %d is not a valid grey pixel", x, y)
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

func TestNegativePixelSize(t *testing.T){
	negative := rand.Int() * -1
	ft, _ := filter.NewFilter(validImage)
	err := ft.PixelFilter(negative)
	
	//if PixelFilter is not defined
	if err == filter.ErrNotDefined {
		t.Errorf("PixelFilter is not defined")
		return
	}
	if err == nil{
		t.Errorf("PixelFilter do not return an error with negative parameter")
		return
	}
}

func TestGreaterThanImagePixelSize(t *testing.T){
	plus := rand.Int()
	ft, _ := filter.NewFilter(validImage)
	err := ft.PixelFilter(ft.Size.X+plus)
	
	//if PixelFilter is not defined
	if err == filter.ErrNotDefined {
		t.Errorf("PixelFilter is not defined" )
		return
	}
	
	if err == nil{
		t.Errorf("PixelFilter do not return an error with greater than image"+
				 "size parameter")
		return
	}
	
	err = ft.PixelFilter(ft.Size.Y+plus)
	if err == nil{
		t.Errorf("PixelFilter do not return an error with greater than image"+
				 "size parameter")
		return
	}
}
//funic 