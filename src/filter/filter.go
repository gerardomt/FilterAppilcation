package filter

import (
	"os"
	"image"
	"image/jpeg"
	"path/filepath"
	"fmt"
	"strings"
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

func (filter *Filter) SaveImageAt(path string) (string,error){
	if filter.Buffer == nil{
		return "", errors.New("No filter applied")
	}
	file, err := os.Create(path)
	defer file.Close()
	check(err)
	err = jpeg.Encode(file, filter.Buffer, nil)
	check(err)

	return path, nil
}

func (filter *Filter) SaveImage(suffix string) (string, error){
	ext := filepath.Ext(filter.ImgPath)
	name := strings.TrimSuffix(filepath.Base(filter.ImgPath), ext)
	newImagePath := fmt.Sprintf("%s/%s_%s%s", filepath.Dir(filter.ImgPath), name, suffix, ext)

	_, err := filter.SaveImageAt(newImagePath)
	if err != nil {
		return newImagePath, err
	}

	return newImagePath, nil
}

func (filter *Filter) GreyFilter() error{
	return ErrNotDefined
}

func (filter *Filter) PixelFilter(pixelSize int) error{
	return ErrNotDefined
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
