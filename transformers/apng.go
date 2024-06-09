package transformers

import (
	"fmt"
	"os"
	"pixelizer/renderer"
	"pixelizer/util"

	"github.com/kettek/apng"
)

func Apng(filePath string, pixels int) error {

	nameFile := util.GetNameFromPath(filePath)

	// Abrimos la imagen seleccionada por el usuario
	inFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	// Decodificamos la imagen animada con apng
	a, err := apng.DecodeAll(inFile)
	if err != nil {
		panic(err)
	}

	var pixelatedFrames []apng.Frame

	for _, frame := range a.Frames {
		pixelatedFrame := apng.Frame{
			Image:            renderer.Pixelate(frame.Image, pixels),
			XOffset:          frame.XOffset,
			YOffset:          frame.YOffset,
			DelayNumerator:   frame.DelayNumerator,
			DelayDenominator: frame.DelayDenominator,
			DisposeOp:        frame.DisposeOp,
			BlendOp:          frame.BlendOp,
			IsDefault:        frame.IsDefault,
		}
		pixelatedFrames = append(pixelatedFrames, pixelatedFrame)
	}

	outFile, err := util.CreateFile(fmt.Sprintf("./out/%s.png", nameFile))
	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	err = apng.Encode(outFile, apng.APNG{
		Frames:    pixelatedFrames,
		LoopCount: a.LoopCount,
	})
	if err != nil {
		panic(err)
	}

	return nil
}
