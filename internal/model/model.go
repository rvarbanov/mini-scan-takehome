package model

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/rvarbanov/mini-scan-takehome/pkg/scanning"
)

type Scan struct {
	IP        string    `json:"ip" db:"ip"`
	Port      uint32    `json:"port" db:"port"`
	Service   string    `json:"service" db:"service"`
	Data      string    `json:"data" db:"data"`
	Timestamp int64     `json:"timestamp" db:"timestamp"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func GetDataFromScan(scan scanning.Scan) (string, error) {
	switch scan.DataVersion {
	case scanning.V1:
		// Try to get V1Data directly
		v1Data, ok := scan.Data.(*scanning.V1Data)
		if !ok {
			// If not a V1Data struct, try to get it from the map
			rawData, ok := scan.Data.(map[string]interface{})
			if !ok {
				return "", fmt.Errorf("invalid V1 data format: %T", scan.Data)
			}
			// Get the base64 string from the map
			responseStr, ok := rawData["response_bytes_utf8"].(string)
			if !ok {
				return "", fmt.Errorf("missing or invalid response_bytes_utf8")
			}
			// Decode base64
			decodedData, err := base64.StdEncoding.DecodeString(responseStr)
			if err != nil {
				return "", err
			}
			return string(decodedData), nil
		}
		// If it was a V1Data struct, decode it directly
		decodedData, err := base64.StdEncoding.DecodeString(string(v1Data.ResponseBytesUtf8))
		if err != nil {
			return "", err
		}
		return string(decodedData), nil

	case scanning.V2:
		v2Data, ok := scan.Data.(*scanning.V2Data)
		if !ok {
			rawData := scan.Data.(map[string]interface{})
			v2Data = &scanning.V2Data{
				ResponseStr: rawData["response_str"].(string),
			}
		}
		return v2Data.ResponseStr, nil

	default:
		return "", fmt.Errorf("invalid data version: %d", scan.DataVersion)
	}
}
