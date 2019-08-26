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
	return nil, nil
}

func (filter *Filter) SetImage(imgPath string) (error){
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
