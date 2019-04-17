package textdata

import (
	"os"

	"github.com/pakuyuya/gopathmatch"
)

func ListupFilePathes(pathpattern string) []string {
	return gopathmatch.Listup(pathpattern, gopathmatch.FlgFileOnly)
}

func LoadAllCsv(pathpattern string) ([]Textdata, error) {
	files := gopathmatch.Listup(pathpattern, gopathmatch.FlgFileOnly)
	datas := make([]Textdata, 0)

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		dats, err := DecodeCSV(f)
		if err != nil {
			return nil, err
		}
		datas = append(datas, dats...)
	}

	return datas, nil
}

func WriteCsv(path string, datas []Textdata) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = EncodeCSV(f, datas)
	if err != nil {
		return err
	}

	return nil
}
