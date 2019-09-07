package gui

import (
	"os"
	"path"
	"filter"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"github.com/mattn/go-gtk/gdkpixbuf"
)

func LoadImage(filename string, image *gtk.Image){
	pixx, _ := gdkpixbuf.NewPixbufFromFileAtScale(filename, 600, 600, true)
    image.SetFromPixbuf(pixx)
}

func CreateDirectory(path string){
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err){
			os.Mkdir(path, 0755)
		}
	}
}

func InitWindow() {
	gtk.Init(nil)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("Filtros")

	//Botón de salida
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		gtk.MainQuit()
	})

	vbox :=gtk.NewVBox(false,0)
	menubar := gtk.NewMenuBar()
	vbox.PackStart(menubar, false, false, 0)

	//Carga una imagen por default
	dir, _ := path.Split(os.Args[0])
	imagefile := path.Join(dir, "../media/fotitus.jpg")
		
	pixx, _ := gdkpixbuf.NewPixbufFromFileAtScale(imagefile, 600, 600, true)
	image := gtk.NewImageFromPixbuf(pixx)

	imageFilter, _ := filter.NewFilter(imagefile)
	
	vbox.Add(image)
	CreateDirectory("../Cache")

	//MENÚ
	//Menú de archivos
	filemenu := gtk.NewMenuItemWithMnemonic("_File")
	menubar.Append(filemenu)
	subfilemenu := gtk.NewMenu()

	filemenu.SetSubmenu(subfilemenu)

	//Abrir archivo
	openFileItem := gtk.NewMenuItemWithMnemonic("_Open")
		openFileItem.Connect("activate", func(){
		filechooserdialog := gtk.NewFileChooserDialog(
			"Choose File...",
			openFileItem.GetTopLevelAsWindow(),
		       	gtk.FILE_CHOOSER_ACTION_OPEN,
	       		gtk.STOCK_OK,
	       		gtk.RESPONSE_ACCEPT)
 		filter := gtk.NewFileFilter()
	   	filter.AddPattern("*.jpg")
	   	filechooserdialog.AddFilter(filter)
		filechooserdialog.Response(func() {
			filePath := filechooserdialog.GetFilename()
			LoadImage(filePath, image)
		   	imageFilter.SetImage(filePath)
	       		filechooserdialog.Destroy()
		})
		filechooserdialog.Run()
	})
	subfilemenu.Append(openFileItem)

	//Guardar archivo
	saveFileItem := gtk.NewMenuItemWithMnemonic("_Save")
	saveFileItem.Connect("activate", func(){
		filechooserdialog := gtk.NewFileChooserDialog(
			"Save File...",
			saveFileItem.GetTopLevelAsWindow(),
			gtk.FILE_CHOOSER_ACTION_SAVE,
			gtk.STOCK_SAVE,
			gtk.RESPONSE_ACCEPT)
		filechooserdialog.Response(func() {
			filePath := filechooserdialog.GetFilename()
			_, err := imageFilter.SaveImageAt(filePath)
			if err != nil {
				messagedialog := gtk.NewMessageDialog(
					filechooserdialog.GetTopLevelAsWindow(),
					gtk.DIALOG_MODAL,
					gtk.MESSAGE_INFO,
					gtk.BUTTONS_OK,
					"No filter applied",
				)
				messagedialog.Response(func() {
					messagedialog.Destroy()
				})
				messagedialog.Run()
			}
			filechooserdialog.Destroy()
		})
		filechooserdialog.Run()
	})
	subfilemenu.Append(saveFileItem)

	//Salir
	exititem := gtk.NewMenuItemWithMnemonic("E_xit")
	exititem.Connect("activate", func() {
	    gtk.MainQuit()
	})
	subfilemenu.Append(exititem)

	//Menú de Filtros
	filtermenu := gtk.NewMenuItemWithMnemonic("Fi_lters")
	menubar.Append(filtermenu)
	subfiltermenu := gtk.NewMenu()

	filtermenu.SetSubmenu(subfiltermenu)
	
	originalItem := gtk.NewMenuItemWithMnemonic("_Original")
	originalItem.Connect("activate", func(){
		LoadImage(imageFilter.ImgPath, image)
	})
	subfiltermenu.Append(originalItem)

	//Filtro Azul
	blueFilterItem := gtk.NewMenuItemWithMnemonic("_Blue")
	blueFilterItem.Connect("activate", func(){
		imageFilter.BlueFilter()
		path, _ := imageFilter.SaveImageAt("../Cache/cache")
		LoadImage(path, image)
	})
	subfiltermenu.Append(blueFilterItem)

	//Filtro Rojo
	redFilterItem := gtk.NewMenuItemWithMnemonic("_Red")
	redFilterItem.Connect("activate", func(){
		imageFilter.RedFilter()
		path, _ := imageFilter.SaveImageAt("../Cache/cache")
		LoadImage(path, image)
	})
	subfiltermenu.Append(redFilterItem)

	//Filtro Verde
	greenFilterItem := gtk.NewMenuItemWithMnemonic("_Green")
	greenFilterItem.Connect("activate", func(){
		imageFilter.GreenFilter()
		path, _ := imageFilter.SaveImageAt("../Cache/cache")
		LoadImage(path, image)
	})
	subfiltermenu.Append(greenFilterItem)

	//Filtro Gris
	greyFilterItem := gtk.NewMenuItemWithMnemonic("_Grey")
	greyFilterItem.Connect("activate", func(){		
		imageFilter.GreyFilter()
		path, _ := imageFilter.SaveImageAt("../Cache/cache")
		LoadImage(path, image)
	})
	subfiltermenu.Append(greyFilterItem)

	//Filtro Mosaico
	pixelItem := gtk.NewMenuItemWithMnemonic("_Pixelation")
	pixelItem.Connect("activate", func(){
		var size int
		subwindow := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
		subwindow.SetPosition(gtk.WIN_POS_CENTER)
		subwindow.SetTitle("Big Pixel Size")

		subbox :=gtk.NewVBox(false, 1)
		label :=gtk.NewLabel("Select the size of the pixel")
		subbox.Add(label)
		scale := gtk.NewHScaleWithRange(10, 100, 10)
		subbox.Add(scale)
		button := gtk.NewButtonWithLabel("Create")
		button.Clicked(func(){
			size = int(scale.GetValue())
			subwindow.Destroy()
			imageFilter.PixelFilter(size)
			path, _ := imageFilter.SaveImageAt("../Cache/cache")
			LoadImage(path, image)
		})
		subbox.Add(button)
		subwindow.Add(subbox)
		subwindow.SetSizeRequest(300, 150)
		subwindow.ShowAll()
	})
	subfiltermenu.Append(pixelItem)

	window.Add(vbox)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
}
