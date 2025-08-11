package protocol

import (
	"github.com/caramel-juni/onionscan-v3links/config"
	"github.com/caramel-juni/onionscan-v3links/report"
)

type Scanner interface {
	ScanProtocol(hiddenService string, onionscanConfig *config.OnionScanConfig, report *report.OnionScanReport)
}
