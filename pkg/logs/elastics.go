package logs

import (
	"bytes"
	"encoding/json"
	"fmt"

	domain "github.com/billowdev/go-fiber-e-commerce/internal/core/domain/logging"
	"github.com/elastic/go-elasticsearch/v8"
)

func logToElasticsearch(logEntry domain.LogEntry) error {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return err
	}

	// Convert logEntry to JSON
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(logEntry); err != nil {
		return err
	}

	// Index the log entry in Elasticsearch
	res, err := es.Index("logs", &buf)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error indexing log: %s", res.Status())
	}

	return nil
}
