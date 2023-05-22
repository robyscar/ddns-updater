package update

import (
	"context"
	"net"
	"time"

	"github.com/qdm12/ddns-updater/internal/records"
)

type PublicIPFetcher interface {
	IP(ctx context.Context) (net.IP, error)
	IP4(ctx context.Context) (net.IP, error)
	IP6(ctx context.Context) (net.IP, error)
}

type UpdaterInterface interface {
	Update(ctx context.Context, recordID uint, ip net.IP, now time.Time) (err error)
}

type Database interface {
	Select(recordID uint) (record records.Record, err error)
	SelectAll() (records []records.Record)
	Update(recordID uint, record records.Record) (err error)
}

type LookupIPer interface {
	LookupIP(ctx context.Context, network, host string) (ips []net.IP, err error)
}

type Logger interface {
	DebugLogger
	Info(s string)
	Warn(s string)
	Error(s string)
}
