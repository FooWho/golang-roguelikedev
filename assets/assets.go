package assets

import (
	"embed"
	"errors"
	"io/fs"
	"log"
)

//go:embed *.png assets.json
var AssetsFS embed.FS

func GetSpriteSheet() []byte {
	data, err := AssetsFS.ReadFile("Full_2.png")
	if err == nil {
		return data
	}

	data, err = AssetsFS.ReadFile("dejavu10x10_gs_tc.png")
	if err != nil {
		log.Fatal("Could not load any spite sheets!")
	}
	return data
}

func FileExists(fsys embed.FS, filename string) bool {
	_, err := fs.Stat(fsys, filename)
	return !errors.Is(err, fs.ErrNotExist)
}
