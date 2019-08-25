package gui

import (
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func LoadImage(filename string, image *gtk.Image){
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

	//MENÚS 
	//Menú de archivos
	filemenu := gtk.NewMenuItemWithMnemonic("_File")
	menubar.Append(filemenu)
	subfilemenu := gtk.NewMenu()

	filemenu.SetSubmenu(subfilemenu)

	//Abrir archivo
	openFileItem := gtk.NewMenuItemWithMnemonic("_Open")
	subfilemenu.Append(openFileItem)

	//Guardar archivo
	saveFileItem := gtk.NewMenuItemWithMnemonic("_Save")
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

	//Filtro Azul
	blueFilterItem := gtk.NewMenuItemWithMnemonic("_Blue")
	subfiltermenu.Append(blueFilterItem)

	//Filtro Rojo
	redFilterItem := gtk.NewMenuItemWithMnemonic("_Red")
	subfiltermenu.Append(redFilterItem)

	//Filtro Verde
	greenFilterItem := gtk.NewMenuItemWithMnemonic("_Green")
	subfiltermenu.Append(greenFilterItem)

	//Filtro Gris
	greyFilterItem := gtk.NewMenuItemWithMnemonic("_Grey")
	subfiltermenu.Append(greyFilterItem)

	//Filtro Mosaico
	pixelItem := gtk.NewMenuItemWithMnemonic("_Pixelation")
	subfiltermenu.Append(pixelItem)

	window.Add(vbox)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
}
