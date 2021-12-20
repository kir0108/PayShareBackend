package redis

import "time"

type Config struct {
	RedisUrl             string
	CodeHideTTL          time.Duration
	CodeExpiredTTL       time.Duration
	SessionTTl           time.Duration
	SessionCleanupPeriod time.Duration
	SessionWindowPeriod  time.Duration
}
