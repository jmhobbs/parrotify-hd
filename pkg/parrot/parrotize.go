package parrot

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/png"
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"github.com/disintegration/imaging"
)

const DEFAULT_SCALE int = 70

//go:embed parrot.gif
var GifBytes []byte

var centers []image.Point

func init() {
	centers = []image.Point{
		image.Pt(80, 54),
		image.Pt(70, 52),
		image.Pt(55, 52),
		image.Pt(45, 62),
		image.Pt(42, 62),
		image.Pt(45, 64),
		image.Pt(60, 66),
		image.Pt(76, 68),
		image.Pt(82, 62),
		image.Pt(86, 58),
	}
}

func Overlay(overlay image.Image, scale int, shiftX, shiftY int, flip bool, rotate float64) ([]byte, error) {
	buf := bytes.NewBuffer(GifBytes)

	var err error
	frames, err := gif.DecodeAll(buf)
	if err != nil {
		panic(err)
	}

	overlayScaled := imaging.Resize(overlay, int(DEFAULT_SCALE+scale), 0, imaging.Lanczos)
	if flip {
		overlayScaled = imaging.FlipH(overlayScaled)
	}
	overlayScaled = imaging.Rotate(overlayScaled, rotate*-1, color.Transparent)

	halfOverlayWidth := overlayScaled.Bounds().Dx() / 2
	halfOverlayHeight := overlayScaled.Bounds().Dy() / 2

	dir, err := os.MkdirTemp("", "parrot")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(dir)

	for i, center := range centers {
		frame := frames.Image[i]
		dst := image.NewRGBA(frame.Bounds())
		draw.Draw(dst, frame.Bounds(), frame, image.Point{0, 0}, draw.Over)

		position := offsetRectangle(overlayScaled.Bounds(), center.X-halfOverlayWidth+shiftX, center.Y-halfOverlayHeight+shiftY)
		draw.Draw(dst, position, overlayScaled, image.Point{0, 0}, draw.Over)

		out, err := os.Create(path.Join(dir, fmt.Sprintf("out.%d.png", i)))
		if err != nil {
			panic(err)
		}

		err = png.Encode(out, dst)
		if err != nil {
			panic(err)
		}
	}

	// rather than build our own palettes we write out
	// to disk and let imagemagick do the hard work
	cmd := exec.Command(
		"convert",
		"-delay", "5",
		"-loop", "0",
		"-dispose", "previous",
		path.Join(dir, "out.0.png"),
		path.Join(dir, "out.1.png"),
		path.Join(dir, "out.2.png"),
		path.Join(dir, "out.3.png"),
		path.Join(dir, "out.4.png"),
		path.Join(dir, "out.5.png"),
		path.Join(dir, "out.6.png"),
		path.Join(dir, "out.7.png"),
		path.Join(dir, "out.8.png"),
		path.Join(dir, "out.9.png"),
		path.Join(dir, "out.gif"))

	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	return ioutil.ReadFile(path.Join(dir, "out.gif"))
}

func offsetRectangle(rect image.Rectangle, x, y int) image.Rectangle {
	return image.Rect(x, y, rect.Dx()+x, rect.Dy()+y)
}
