package filter

import (
	"os"
	"image"
	_ "image/jpeg"
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

func (filter *Filter) GreyFilter() error{
	return errors.New("Not defined")
}

func (filter *Filter) PixelFilter(pixelSize int) error{
	return errors.New("Not defined")
}

func (filter *Filter) ColorFilter(r, g, b uint8) error{
	return errors.New("Not defined")
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
