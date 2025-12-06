package theme

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"

	"gioui.org/font"
	"gioui.org/font/gofont"
	"gioui.org/font/opentype"
	"gioui.org/text"
)

var extentions = []string{".otf", ".otc", ".ttf", ".ttc"}

var builtinFonts []font.FontFace

func init() {
	fonts, err := LoadBuiltinFonts()
	if err != nil {
		log.Println("error loading builtin fonts", err)
		return
	}
	builtinFonts = fonts
}

//go:embed fonts
var fontFiles embed.FS

func LoadBuiltinFonts() ([]font.FontFace, error) {
	fonts := []font.FontFace{}
	for _, font := range gofont.Collection() {
		fonts = append(fonts, font)
	}

	files, err := fontFiles.ReadDir(".")
	if err != nil {
		return nil, fmt.Errorf("failed to read font files: %v", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filename := file.Name()
		if !slices.Contains(extentions, filepath.Ext(filename)) {
			continue
		}

		ttfData, err := os.ReadFile(filepath.Join(".", filename))
		if err != nil {
			log.Printf("read font %s from dir failed: %v", filename, err)
			continue
		}

		font, err := loadFont(ttfData)
		if err != nil {
			log.Println("error loading font", filename, err)
			continue
		}
		fonts = append(fonts, *font)
	}

	return fonts, nil
}

func loadFont(ttf []byte) (*font.FontFace, error) {
	faces, err := opentype.ParseCollection(ttf)
	if err != nil {
		return nil, fmt.Errorf("failed to parse font: %v", err)
	}

	return &text.FontFace{
		Font: faces[0].Font,
		Face: faces[0].Face,
	}, nil
}

func BuiltinFonts() []font.FontFace {
	return builtinFonts
}
