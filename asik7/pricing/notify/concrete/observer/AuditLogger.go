package observer

import (
	"fmt"

	"asik7/pricing/notify"
)

type AuditLogger struct {
	source string
}

func NewAuditLogger(source string) AuditLogger {
	return AuditLogger{source: source}
}

func (a AuditLogger) OnNotify(e notify.Event) {
	fmt.Printf("[AUDIT:%s] type=%s msg=%s data=%v\n", a.source, e.Type, e.Message, e.Data)
}
