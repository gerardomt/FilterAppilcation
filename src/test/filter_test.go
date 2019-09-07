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

// Prueba que se devuelva un error si el constructor recibe una imagen que no
// sea jpeg
func TestNewFilter(t *testing.T){
	_, err := filter.NewFilter(noValidImage)
	if err == nil {
		t.Errorf("NewFilter did not return an error")
	}
}

// Prueba que se devuelva un error si SetImage recibe una image que no sea jpeg
func TestSetImage(t *testing.T){
	ft, _ := filter.NewFilter(validImage)
	err := ft.SetImage(noValidImage)
	if err == nil {
		t.Errorf("SetImage did not return an error")
	}
}

// Prueba que SaveImage devuelva un error si no se ha aplicado un filtro antes
// de llamarla
func TestSaveImage(t *testing.T){
	ft, _ := filter.NewFilter(validImage)
	_, err := ft.SaveImage("")
	if err == nil {
		t.Errorf("SaveImage did not return an error")
	}
}

// Prueba que ColorFilter devueva un error si se le pasa algún parámetro menor
// que cero
func TestColorFilterParameterLessThanZero(t *testing.T){
	test := [3]float32{-1, 1, 0}
	rand.Shuffle(len(test), func(i,j int){
		test[i], test[j] = test[j], test[i]
	})
	ft, _ := filter.NewFilter(validImage)
	err := ft.ColorFilter(test[0], test[1], test[2])
	if err == nil {
		t.Errorf("ColorFilter dn not return an error with parameter less than zero")
	}
}

// Prueba que ColorFilter devuelva un error si se le pasa algún parámetro mayor
// que uno
func TestColorFilterParameterGreaterThanOne(t *testing.T){
	test := [3]float32{2.0, 1.0, 0.0}
	rand.Shuffle(len(test), func(i,j int){
		test[i], test[j] = test[j], test[i]
	})
	ft, _ := filter.NewFilter(validImage)
	err := ft.ColorFilter(test[0], test[1], test[2])
	if err == nil {
		t.Errorf("ColorFilter do not return an error with parameter greater than 1")
	}
}

// Prueba que después de aplicar RedFilter y tomando pixeles al azar todos estos
// tengas valores b,g iguales a cero 
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

// Prueba que después de aplicar BlueFilter y tomando pixeles al azar todos estos
// tengas valores r,g iguales a cero 
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

// Prueba que después de aplicar GreenFilter y tomando pixeles al azar todos estos
// tengas valores r,b iguales a cero 
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

// Prueba que después de aplicar GreyFilter y tomando pixeles al azar todos estos
// sean el resultado de aplicar la fórmula LUMA a los valores en la imagen original 
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
		r := float32(originalPixel.R) * 0.92126
		g := float32(originalPixel.G) * 0.97152
		b := float32(originalPixel.B) * 0.90722

		grey := uint8((r + g + b) / 3)

		if pixel.B!=pixel.R || pixel.G!=pixel.B || pixel.R!=pixel.G || pixel.B!=grey {
			t.Errorf("Pixel at %d, %d is not a valid grey pixel", x, y)
			break
		}
	}
}

// Prueba que después de aplicar PixelFilter todos los pixeles en regiones de
// 30x30(arbitrario) tengan los mismos valores de pixeles
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

// Prueba que PixelFilter devuelva un error si le pasa un parámetro menor que
// cero
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

// Prueba que PixelFilter devuelva un error si se le pasa un parámetro mayor al
// tamaño de la imagen
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