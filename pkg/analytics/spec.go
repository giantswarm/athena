package analytics

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type Event struct {
	AppID                AppID                  `firestore:"app_id"`
	SessionID            string                 `firestore:"session_id"`
	PayloadType          string                 `firestore:"payload_type"`
	PayloadSchemaVersion int                    `firestore:"payload_schema_version"`
	Payload              map[string]interface{} `firestore:"payload"`
	URI                  string                 `firestore:"uri_path"`
	Timestamp            time.Time              `firestore:"timestamp"`
	InstallationID       string                 `firestore:"installation_id"`
	EnvironmentClass     string                 `firestore:"env_class"`
}

type Reporter interface {
	Report(ctx context.Context, event Event) (Event, error)
}

func (e *Event) CollectionName() string {
	envClass := strings.ToLower(e.EnvironmentClass)

	return fmt.Sprintf("%s-%s", e.Timestamp.Format("2006-01"), envClass)
}
