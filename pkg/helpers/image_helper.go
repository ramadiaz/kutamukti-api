package helpers

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"kutamukti-api/pkg/exceptions"
	"net/http"
	"strings"

	"golang.org/x/image/draw"
)

func ResizeImage(imageData []byte, width int) ([]byte, *exceptions.Exception) {
	reader := bytes.NewReader(imageData)

	img, format, err := image.Decode(reader)
	if err != nil {
		return nil, exceptions.NewException(http.StatusInternalServerError, exceptions.ErrImageDecode)
	}

	bounds := img.Bounds()
	originalWidth := bounds.Dx()
	originalHeight := bounds.Dy()

	targetWidth := width

	if originalWidth <= targetWidth {
		targetWidth = originalWidth
	}

	aspectRatio := float64(originalHeight) / float64(originalWidth)
	targetHeight := int(float64(targetWidth) * aspectRatio)

	newImg := image.NewRGBA(image.Rect(0, 0, targetWidth, targetHeight))

	draw.CatmullRom.Scale(newImg, newImg.Bounds(), img, bounds, draw.Over, nil)

	var buf bytes.Buffer

	switch strings.ToLower(format) {
	case "jpeg", "jpg":
		err = jpeg.Encode(&buf, newImg, &jpeg.Options{Quality: 85})
	case "png":
		err = png.Encode(&buf, newImg)
	default:

		err = jpeg.Encode(&buf, newImg, &jpeg.Options{Quality: 85})
	}

	if err != nil {
		return nil, exceptions.NewException(http.StatusInternalServerError, exceptions.ErrImageEncode)
	}

	fmt.Printf("Image resized from %dx%d to %dx%d\n", originalWidth, originalHeight, targetWidth, targetHeight)
	return buf.Bytes(), nil
}
