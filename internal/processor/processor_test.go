package processor

import (
	"context"
	"testing"

	"cloud.google.com/go/pubsub"
	"github.com/golang/mock/gomock"
	mock_db "github.com/rvarbanov/mini-scan-takehome/internal/db/mock"
	"github.com/stretchr/testify/suite"
)

type ProcessorTestSuite struct {
	suite.Suite
	ctrl      *gomock.Controller
	ctx       context.Context
	db        *mock_db.MockDBInterface
	processor *Processor
}

func TestProcessorTestSuite(t *testing.T) {
	suite.Run(t, new(ProcessorTestSuite))
}

func (suite *ProcessorTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.db = mock_db.NewMockDBInterface(suite.ctrl)
	suite.processor = New(suite.db)
}

func (suite *ProcessorTestSuite) TestProcessMessage_v1() {
	suite.db.EXPECT().StoreScan(gomock.Any(), gomock.Any()).Return(nil)
	pubsubMessage := &pubsub.Message{Data: temp1()}

	suite.processor.ProcessMessage(suite.ctx, pubsubMessage)
}

func (suite *ProcessorTestSuite) TestProcessMessage_v2() {
	suite.db.EXPECT().StoreScan(gomock.Any(), gomock.Any()).Return(nil)
	pubsubMessage := &pubsub.Message{Data: temp2()}

	suite.processor.ProcessMessage(suite.ctx, pubsubMessage)
}

func temp1() []byte {
	return []byte(`{"ip":"127.0.0.1","port":80,"service":"http","timestamp":1733858499, "data_version": 1, "data": {"response_bytes_utf8": "aGVsbG8gd29ybGQ="}}`)
}

func temp2() []byte {
	return []byte(`{"ip":"127.0.0.1","port":80,"service":"http","timestamp":1733858499, "data_version": 2, "data": {"response_str": "service response: 61"}}`)
}
