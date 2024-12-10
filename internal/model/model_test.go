package model

import (
	"encoding/base64"
	"testing"

	"github.com/rvarbanov/mini-scan-takehome/pkg/scanning"
	"github.com/stretchr/testify/suite"
)

type ModelTestSuite struct {
	suite.Suite
}

func TestModelTestSuite(t *testing.T) {
	suite.Run(t, new(ModelTestSuite))
}

func (suite *ModelTestSuite) TestGetDataFromScan_V1() {
	encodedData := base64.StdEncoding.EncodeToString([]byte("hello world"))

	scan := scanning.Scan{
		Ip:          "127.0.0.1",
		Port:        80,
		Service:     "http",
		Timestamp:   1733858499,
		DataVersion: scanning.V1,
		Data:        &scanning.V1Data{ResponseBytesUtf8: []byte(encodedData)},
	}

	scanData, err := GetDataFromScan(scan)
	suite.Require().NoError(err)
	suite.Require().Equal("hello world", scanData)
}

func (suite *ModelTestSuite) TestGetDataFromScan_V2() {
	scan := scanning.Scan{
		Ip:          "127.0.0.1",
		Port:        80,
		Service:     "http",
		Timestamp:   1733858499,
		DataVersion: scanning.V2,
		Data:        &scanning.V2Data{ResponseStr: "service response: 61"},
	}

	scanData, err := GetDataFromScan(scan)
	suite.Require().NoError(err)
	suite.Require().Equal("service response: 61", scanData)
}
