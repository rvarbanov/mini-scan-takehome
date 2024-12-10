package processor

import (
	"context"
	"testing"

	"cloud.google.com/go/pubsub"
	"github.com/stretchr/testify/suite"

	"github.com/rvarbanov/mini-scan-takehome/internal/db"
)

// in this test we are going to test inserting into the db

type ProcessorE2ETestSuite struct {
	suite.Suite

	ctx       context.Context
	db        db.DBInterface
	processor *Processor
}

func TestProcessorE2ETestSuite(t *testing.T) {
	suite.Run(t, new(ProcessorE2ETestSuite))
}

func (suite *ProcessorE2ETestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.db = db.NewDB("localhost", "5432", "postgres", "postgres", "postgres")
	suite.processor = New(suite.db)
}

func (suite *ProcessorE2ETestSuite) TestProcessMessage_v1() {
	suite.T().Skip()
	pubsubMessage := &pubsub.Message{Data: temp1()}

	suite.processor.ProcessMessage(suite.ctx, pubsubMessage)
}

func (suite *ProcessorE2ETestSuite) TestProcessMessage_v2() {
	suite.T().Skip()
	pubsubMessage := &pubsub.Message{Data: temp2()}

	suite.processor.ProcessMessage(suite.ctx, pubsubMessage)
}
