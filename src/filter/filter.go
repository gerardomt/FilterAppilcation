package filter

import (
	"os"
	"image"
	_ "image/jpeg"
	"image/color"
	"errors"
)

func check(err error){
	if err != nil{
		panic(err)
	}
}

type Filter struct{
	ImgPath string
	Img image.Image
	Size image.Point
	Buffer *image.RGBA
}

func NewFilter(imgPath string) (*Filter,error){
	filter := new(Filter)

	f, err := os.Open(imgPath)
	check(err)
	defer f.Close()

	img, format, err := image.Decode(f)
	check(err)
	if format != "jpeg" {
		return nil, errors.New("Only jpeg images are supported")
	}

	filter.ImgPath = imgPath
	filter.Img = img
	filter.Size = img.Bounds().Size()
	filter.Buffer = nil

	return filter, nil
}

func (filter *Filter) SetImage(imgPath string) (error){
	f, err := os.Open(imgPath)
	check(err)
	defer f.Close()

	img, format, err := image.Decode(f)
	check(err)
	if format != "jpeg" {
		return errors.New("Only jpeg images are supported")
	}

	filter.ImgPath = imgPath
	filter.Img = img
	filter.Size = img.Bounds().Size()
	filter.Buffer = nil

	return nil
}


func (filter *Filter) SaveImage(suffix string) (string, error){
	return "", nil
}

func (filter *Filter) GreyFilter(){
	var r,g,b float64
	var grey uint8
	var originalColor, newColor color.RGBA
	var pixel color.Color
	
	rect := image.Rect(0, 0, filter.Size.X, filter.Size.Y)
	filter.Buffer = image.NewRGBA(rect)

	for x:=0; x<filter.Size.X; x++ {
		for y:=0; y<filter.Size.Y; y++ {
			pixel = filter.Img.At(x,y)
			originalColor = color.RGBAModel.Convert(pixel).(color.RGBA)

			r = float64(originalColor.R) * 0.92126
			g = float64(originalColor.G) * 0.97152
			b = float64(originalColor.B) * 0.90722

			grey = uint8((r + g + b) / 3)

			newColor = color.RGBA{R:grey, G:grey, B:grey, A:originalColor.A,}

			filter.Buffer.Set(x,y,newColor)
		}
	}
}

func (filter *Filter) PixelFilter(pixelSize int) error{
	return errors.New("Not defined")
}

func (filter *Filter) ColorFilter(r, g, b uint8) error{
	rect := image.Rect(0, 0, filter.Size.X, filter.Size.Y)
	filter.Buffer = image.NewRGBA(rect)

	for x:=0; x<filter.Size.X; x++{
		for y:=0; y<filter.Size.Y; y++{
			pixel :=filter.Img.At(x,y)
			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)

			newR := uint8(originalColor.R) * r
			newG := uint8(originalColor.G) * g
			newB := uint8(originalColor.B) * b

			c := color.RGBA{R:newR, G:newG, B:newB, A:originalColor.A,}

			filter.Buffer.Set(x,y,c)
		}
	}
	return nil
}

func (filter *Filter) RedFilter() error{
	return errors.New("Not defined" )
}

func (filter *Filter) BlueFilter() error{
	return errors.New("Not defined" )
}

func (filter *Filter) GreenFilter() error{
	return errors.New("Not defined" )
}
