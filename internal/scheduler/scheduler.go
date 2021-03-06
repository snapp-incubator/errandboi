package scheduler

import (
	"fmt"
	"time"

	"github.com/gammazero/workerpool"
	"github.com/snapp-incubator/errandboi/internal/publisher"
	"go.uber.org/zap"
)

type scheduler struct {
	Publisher *publisher.Publisher
	Stop      chan struct{}
	Logger    *zap.Logger
}

var errPublisherNil = fmt.Errorf("publisher cannot be nil")

// nolint:revive
func NewScheduler(pb *publisher.Publisher, logger *zap.Logger) (*scheduler, error) {
	sch := &scheduler{Stop: make(chan struct{}), Logger: logger}
	if pb != nil {
		sch.Publisher = pb

		return sch, nil
	}

	return nil, errPublisherNil
}

func (sch *scheduler) WorkInIntervals(d time.Duration) {
	ticker := time.NewTicker(d)

	go func() {
		for {
			select {
			case <-ticker.C:
				sch.Publisher.GetEvents()
				sch.Publisher.Work()
			case <-sch.Stop:
				sch.Publisher.Cancel()
				sch.Publisher.Wp = workerpool.New(sch.Publisher.WorkerSize)

				ticker.Stop()

				return
			}
		}
	}()
}
