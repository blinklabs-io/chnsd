package indexer

import (
	"fmt"

	"github.com/blinklabs-io/chnsd/internal/config"
	"github.com/blinklabs-io/chnsd/internal/logging"

	"github.com/blinklabs-io/snek/event"
	//filter_chainsync "github.com/blinklabs-io/snek/filter/chainsync"
	filter_event "github.com/blinklabs-io/snek/filter/event"
	input_chainsync "github.com/blinklabs-io/snek/input/chainsync"
	output_embedded "github.com/blinklabs-io/snek/output/embedded"
	"github.com/blinklabs-io/snek/pipeline"
)

type Indexer struct {
	pipeline *pipeline.Pipeline
}

// Singleton indexer instance
var globalIndexer = &Indexer{}

func (i *Indexer) Start() error {
	cfg := config.GetConfig()
	logger := logging.GetLogger()
	// Create pipeline
	i.pipeline = pipeline.New()
	// Configure pipeline input
	inputOpts := []input_chainsync.ChainSyncOptionFunc{
		input_chainsync.WithIntersectTip(true),
	}
	if cfg.Node.NetworkMagic > 0 {
		inputOpts = append(
			inputOpts,
			input_chainsync.WithNetworkMagic(cfg.Node.NetworkMagic),
		)
	} else {
		inputOpts = append(
			inputOpts,
			input_chainsync.WithNetwork(cfg.Node.Network),
		)
	}
	input := input_chainsync.New(
		inputOpts...,
	)
	i.pipeline.AddInput(input)
	// Configure pipeline filter
	filterEvent := filter_event.New(
		filter_event.WithType("chainsync.transaction"),
	)
	i.pipeline.AddFilter(filterEvent)
	// TODO: add chainsync filter with address
	// Configure pipeline output
	output := output_embedded.New(
		output_embedded.WithCallbackFunc(i.handleEvent),
	)
	i.pipeline.AddOutput(output)
	// Start pipeline and wait for error
	if err := i.pipeline.Start(); err != nil {
		logger.Fatalf("failed to start pipeline: %s\n", err)
	}
	err, ok := <-i.pipeline.ErrorChan()
	if ok {
		logger.Fatalf("pipeline failed: %s\n", err)
	}
	return nil
}

func (i *Indexer) handleEvent(evt event.Event) error {
	//fmt.Printf("handleEvent(): evt.Payload = %#v\n", evt.Payload)
	eventTx := evt.Payload.(input_chainsync.TransactionEvent)
	for _, txOutput := range eventTx.Outputs {
		datum := txOutput.Datum()
		if datum != nil {
			if _, err := datum.Decode(); err != nil {
				fmt.Printf("handleEvent(): txid = %s, err = %s\n", eventTx.TransactionHash, err)
				return err
			}
			datumJson, _ := datum.MarshalJSON()
			fmt.Printf("handleEvent(): txid = %s, datum = %s\n", eventTx.TransactionHash, datumJson)
		}
	}
	return nil
}

// GetIndexer returns the global indexer instance
func GetIndexer() *Indexer {
	return globalIndexer
}
