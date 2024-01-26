package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-kanban-container-creates-rmq-kube/DPFM_API_Input_Reader"
	"encoding/json"
	"time"

	"golang.org/x/xerrors"
)

func ConvertToHeaderCreates(sdc *dpfm_api_input_reader.SDC) (*Header, error) {
	data := sdc.Header

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}
	// header.CreationDate = *getSystemDatePtr()
	// header.CreationTime = *getSystemTimePtr()
	// header.LastChangeDate = getSystemDatePtr()
	// header.LastChangeTime = getSystemTimePtr()

	return header, nil
}

func ConvertToHeaderUpdates(headerData dpfm_api_input_reader.Header) (*Header, error) {
	data := headerData

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}

func getSystemDatePtr() *string {
	// jst, _ := time.LoadLocation("Asia/Tokyo")
	// day := time.Now().In(jst)

	day := time.Now()
	res := day.Format("2006-01-02")
	return &res
}

func getSystemTimePtr() *string {
	// jst, _ := time.LoadLocation("Asia/Tokyo")
	// day := time.Now().In(jst)

	day := time.Now()
	res := day.Format("15:04:05")
	return &res
}
