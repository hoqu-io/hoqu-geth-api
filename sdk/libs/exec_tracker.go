package libs

import (
    "time"
    log "github.com/sirupsen/logrus"
)

type ExecTracker struct {
    name string
    start time.Time
}

func NewExecTracker(name string) *ExecTracker {
    log.Infof("%s started", name)
    return &ExecTracker{
        name: name,
        start: time.Now(),
    }
}

func (et *ExecTracker) Stop() {
    log.Infof("%s finished in %v", et.name, time.Since(et.start))
}
