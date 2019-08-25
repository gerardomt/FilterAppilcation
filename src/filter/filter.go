package filter

import ("image"
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
	return nil, errors.New("")
}

func (filter *Filter) SetImage(imgPath string) (error){
	return errors.New("")
}


func (filter *Filter) SaveImage(suffix string) (string, error){
	return "", errors.New("")
}

func (filter *Filter) GreyFilter(){
}

func (filter *Filter) PixelFilter(pixelSize int){
}

func (filter *Filter) ColorFilter(r, g, b uint8){
}

func (filter *Filter) RedFilter(){
}

func (filter *Filter) BlueFilter(){
}

func (filter *Filter) GreenFilter(){
}
