package filter

import (
	"os"
	"image"
	"image/jpeg"
	"path/filepath"
	"fmt"
	"strings"
	"image/color"
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
	return nil
}

func (filter *Filter) PixelFilter(pixelSize int) error{
	if pixelSize<=0 || pixelSize>filter.Size.X || pixelSize>filter.Size.Y {
		return errors.New("pixelSize must be greater than zero and less than"+
						  "the size of the image")
	}
	var sumR, sumG, sumB uint
	var myA uint8
	var pixel color.Color
	var originalColor, newColor color.RGBA
	
	rect := image.Rect(0, 0, filter.Size.X, filter.Size.Y)
	filter.Buffer = image.NewRGBA(rect)
	var square uint = uint(pixelSize*pixelSize)
	var a,b int//uint para operarlo con las sumas

	for x:=0; x<filter.Size.X; x+=pixelSize{
		for y:=0; y<filter.Size.Y; y+=pixelSize{
			sumR = 0; sumG = 0; sumB = 0; myA = 0;

			for a=0; a<pixelSize && x+a<filter.Size.X ;a++{
				for b=0; b<pixelSize && y+b<filter.Size.Y; b++{
					pixel = filter.Img.At(x+a,y+b)
					originalColor = color.RGBAModel.Convert(pixel).(color.RGBA)
					sumR += uint(originalColor.R)
					sumG += uint(originalColor.G)
					sumB += uint(originalColor.B)
				}
			}
			myA = originalColor.A
			if a<pixelSize || b<pixelSize {
				newColor = color.RGBA{R:uint8(sumR/uint(a*b)), G:uint8(sumG/uint(a*b)),B:uint8(sumB/uint(a*b)), A:myA,}
			} else {
				newColor = color.RGBA{R:uint8(sumR/square), G:uint8(sumG/square),B:uint8(sumB/square), A:myA,}
			}
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
	//Como r,g,b son uint8 no es necesario revisar que sean mayores a cero
	
	var newR, newG, newB uint8
	var pixel color.Color
	var originalColor, newColor color.RGBA
	
	rect := image.Rect(0, 0, filter.Size.X, filter.Size.Y)
	filter.Buffer = image.NewRGBA(rect)

	for x:=0; x<filter.Size.X; x++ {
		for y:=0; y<filter.Size.Y; y++ {
			pixel = filter.Img.At(x,y)
			originalColor = color.RGBAModel.Convert(pixel).(color.RGBA)

			newR = uint8(originalColor.R) * r
			newG = uint8(originalColor.G) * g
			newB = uint8(originalColor.B) * b

			newColor = color.RGBA{R:newR, G:newG, B:newB, A:originalColor.A,}

			filter.Buffer.Set(x,y,newColor)
		}
	}
	return nil
}

func (filter *Filter) RedFilter() error{
	return filter.ColorFilter(1, 0, 0)
}

func (filter *Filter) BlueFilter() error{
	return filter.ColorFilter(0, 0, 1)
}

func (filter *Filter) GreenFilter() error{
	return filter.ColorFilter(0, 1, 0)
}
