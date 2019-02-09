package textdata

import (
	"encoding/csv"
	"io"
	"time"
)

func DecodeCSV(r io.Reader) ([]Textdata, error) {
	rc := csv.NewReader(r)
	t := time.Now()
	createat := t.Format("2006-01-02T15:04:05")

	datas := make([]Textdata, 0)

	for {
		row, err := rc.Read()
		if err != nil {
			if err.Error() != "EOF" {
				return nil, err
			}
			break
		}
		data := Textdata{}
		data.Caption = ""
		if len(row) >= 1 {
			data.Caption = row[0]
		}
		data.Body = ""
		if len(row) >= 2 {
			data.Body = row[1]
		}
		data.CreateAt = createat
		datas = append(datas, data)
	}

	return datas, nil
}
