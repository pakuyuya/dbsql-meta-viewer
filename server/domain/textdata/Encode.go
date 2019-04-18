package textdata

import (
	"encoding/csv"
	"io"
)

func EncodeCSV(w io.Writer, datas []Textdata) error {
	wcsv := csv.NewWriter(w)
	for _, data := range datas {
		line := []string{data.Namespace, data.Caption, data.Body}
		if err := wcsv.Write(line); err != nil {
			return err
		}
	}
	wcsv.Flush()
	return nil
}
