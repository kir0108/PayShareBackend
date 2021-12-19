package redis

import "time"

type Config struct {
	RedisUrl             string
	CodeTTL              time.Duration
	SessionTTl           time.Duration
	SessionCleanupPeriod time.Duration
	SessionWindowPeriod  time.Duration
}
