package processor

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/rvarbanov/mini-scan-takehome/internal/db"
	"github.com/rvarbanov/mini-scan-takehome/internal/model"
	"github.com/rvarbanov/mini-scan-takehome/pkg/scanning"
)

type Processor struct {
	db db.DBInterface
}

func New(db db.DBInterface) *Processor {
	return &Processor{db: db}
}

func (p *Processor) ProcessMessage(ctx context.Context, msg *pubsub.Message) error {
	fmt.Printf("Processing message: %s\n", string(msg.Data))

	var inboundData scanning.Scan
	err := json.Unmarshal(msg.Data, &inboundData)
	if err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}

	scanData, err := model.GetDataFromScan(inboundData)
	if err != nil {
		return fmt.Errorf("failed to get data from scan: %w", err)
	}

	s := model.Scan{
		IP:        inboundData.Ip,
		Port:      inboundData.Port,
		Service:   inboundData.Service,
		Timestamp: inboundData.Timestamp,
		Data:      scanData,
	}

	return p.db.StoreScan(ctx, s)
}
