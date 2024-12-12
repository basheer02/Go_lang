package conversion

import (
	"encoding/json"
	"errors"
	"strconv"
)

func StringToFloat(strings []string) ([]float64, error) {

	var floats []float64

	for _, val := range strings {
		floatPrice, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return nil, errors.New("error in converting price to float")
		}
		floats = append(floats, floatPrice)
	}
	return floats, nil
}

func StructToJson(data *struct{}) ([]byte, error) {
	json, err := json.Marshal(data)

	if err != nil {
		return nil, errors.New(" failed to convert to json ")
	}

	return json, nil
}
