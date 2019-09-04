package filter

import (
	"os"
	"image"
	"image/color"
	_ "image/jpeg"
	"errors"
)

var ErrNotDefined error = errors.New("Not Defined")

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

func (filter *Filter) GreyFilter() error{
	return ErrNotDefined
}

func (filter *Filter) PixelFilter(pixelSize int) error{
	if pixelSize<=0 || pixelSize>filter.Size.X || pixelSize>filter.Size.Y {
		return errors.New("pixelSize must be greater than zero and less than"+
						  "the size of the image")
	}
	var sumR, sumG, sumB, myA uint8
	var pixel color.Color
	var originalColor, newColor color.RGBA
	
	rect := image.Rect(0, 0, filter.Size.X, filter.Size.Y)
	filter.Buffer = image.NewRGBA(rect)
	var square uint8 = uint8(pixelSize*pixelSize)

	for x:=0; x<filter.Size.X; x=x+pixelSize{
		for y:=0; y<filter.Size.Y; y=y+pixelSize{
			sumR = 0; sumG = 0; sumB = 0; myA = 0;

			for a:=0; a<pixelSize && x+a<filter.Size.X ;a++ {
				for b:=0; b<pixelSize && y+b<filter.Size.Y; b++ {
					pixel = filter.Img.At(x+a,y+b)
					originalColor = color.RGBAModel.Convert(pixel).(color.RGBA)
					sumR += originalColor.R
					sumG += originalColor.G
					sumB += originalColor.B
					if x==x+a && y==y+b {
						myA = originalColor.A
					}
				}
			}

			newColor = color.RGBA{R:sumR/square, G:sumG/square,B:sumB/square, A:myA,}
			for a:=0;a<pixelSize;a++{
				for b:=0;b<pixelSize;b++{
					filter.Buffer.Set(x+a, y+b, newColor)
				}
			}
		}
	}
	return nil
}

func (filter *Filter) ColorFilter(r, g, b uint8) error{
	return ErrNotDefined
}

func (filter *Filter) RedFilter() error{
	return ErrNotDefined
}

func (filter *Filter) BlueFilter() error{
	return ErrNotDefined
}

func (filter *Filter) GreenFilter() error{
	return ErrNotDefined
}
