package textdata

import (
	"encoding/csv"
	"fmt"
	"io"
	"time"
)

func DecodeCSV(r io.Reader) ([]Textdata, error) {
	rc := csv.NewReader(r)
	t := time.Now()
	createat := t.Format("2006-01-02T15:04:05")

	datas := make([]Textdata, 0)

	i := 0
	for {
		row, err := rc.Read()
		i = i + 1
		if err != nil {
			if err.Error() != "EOF" {
				return nil, err
			}
			break
		}
		if len(row) < 2 {
			return nil, fmt.Errorf("invalid row. too less columns num. [position]: row %d [data]: %s", i, row)
		}
		data := Textdata{}
		data.Caption = row[0]
		data.Body = row[1]
		data.CreateAt = createat
		datas = append(datas, data)
	}

	return datas, nil
}
