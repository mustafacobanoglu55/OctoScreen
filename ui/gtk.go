package ui

import (
	"fmt"
	"path/filepath"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

// MustGrid returns a new gtk.Grid, if error panics.
func MustGrid() *gtk.Grid {
	grid, err := gtk.GridNew()
	if err != nil {
		panic(err)
	}

	return grid
}

// MustBox returns a new gtk.Box, with the given configuration, if err panics.
func MustBox(o gtk.Orientation, spacing int) *gtk.Box {
	box, err := gtk.BoxNew(o, spacing)
	if err != nil {
		panic(err)
	}

	return box
}

// MustProgressBar returns a new gtk.ProgressBar, if err panics.
func MustProgressBar() *gtk.ProgressBar {
	p, err := gtk.ProgressBarNew()
	if err != nil {
		panic(err)
	}

	return p
}

// MustLabel returns a new gtk.Label, if err panics.
func MustLabel(label string, args ...interface{}) *gtk.Label {
	l, err := gtk.LabelNew(fmt.Sprintf(label, args...))
	if err != nil {
		panic(err)
	}

	return l
}

// LabelWithImage represents a gtk.Label with a image to the right.
type LabelWithImage struct {
	Label *gtk.Label
	*gtk.Box
}

// LabelImageSize default width and height of the image for a LabelWithImage
const LabelImageSize = 20

// MustLabelWithImage returns a new LabelWithImage based on a gtk.Box containing
// a gtk.Label with a gtk.Image, the image is scaled at LabelImageSize.
func MustLabelWithImage(img, label string, args ...interface{}) *LabelWithImage {
	l := MustLabel(label, args...)
	b := MustBox(gtk.ORIENTATION_HORIZONTAL, 5)
	b.Add(MustImageFromFileWithSize(img, LabelImageSize, LabelImageSize))
	b.Add(l)

	return &LabelWithImage{Label: l, Box: b}
}

// MustButtonImage returns a new gtk.Button with the given label, image and
// clicked callback. If error panics.
func MustButtonImage(label, img string, clicked func()) *gtk.Button {
	b, err := gtk.ButtonNewWithLabel(label)
	if err != nil {
		panic(err)
	}

	b.SetImage(MustImageFromFile(img))
	b.SetAlwaysShowImage(true)
	b.SetImagePosition(gtk.POS_TOP)
	b.SetVExpand(true)
	b.SetHExpand(true)

	//c, _ := b.GetStyleContext()
	//c.AddClass("flat")

	if clicked != nil {
		b.Connect("clicked", clicked)
	}

	return b
}

// MustImageFromFileWithSize returns a new gtk.Image based on rescaled version
// of the given file.
func MustImageFromFileWithSize(img string, w, h int) *gtk.Image {
	p, err := gdk.PixbufNewFromFileAtScale(
		filepath.Join(ImagesFolder, img), w, h, true)

	if err != nil {
		panic(err)
	}

	i, err := gtk.ImageNewFromPixbuf(p)
	if err != nil {
		panic(err)
	}

	return i
}

// MustImageFromFile returns a new gtk.Image based on the given file.
func MustImageFromFile(img string) *gtk.Image {
	i, err := gtk.ImageNewFromFile(filepath.Join(ImagesFolder, img))
	if err != nil {
		panic(err)
	}

	return i
}