package observability

import (
	"context"
	"os"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewRelicApp() (*newrelic.Application, error) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("clean-arch-go"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	if err != nil {
		return nil, err
	}

	if err := app.WaitForConnection(10 * time.Second); err != nil {
		return nil, err
	}

	return app, nil
}

func StartSegment(ctx context.Context, name string) func() {
	txn := newrelic.FromContext(ctx)
	if txn == nil {
		return func() {
			// no-op
		}
	}
	s := txn.StartSegment(name)
	return func() {
		s.End()
	}
}
