package db

import (
	"encoding/base64"
	"strconv"
)

func EncodeCursor(id int) string {
	return base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(id)))
}

func DecodeCursor(encoded string) (int, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(decoded))
}
