package redis

import (
	"math"
	"strconv"
	"testing"
)

func TestParseInt(t *testing.T) {
	for _, v := range []int64{0, -1, 1, math.MinInt64, math.MaxInt64} {
		got := ParseInt([]byte(strconv.FormatInt(v, 10)))
		if got != v {
			t.Errorf("got %d, want %d", got, v)
		}
	}
	if got := ParseInt(nil); got != 0 {
		t.Errorf("got %d for the empty string, want 0", got)
	}
}

func TestNormalizeAddr(t *testing.T) {
	golden := []struct{ Addr, Normal string }{
		{"", "localhost:6379"},
		{":", "localhost:6379"},
		{"test.host", "test.host:6379"},
		{"test.host:", "test.host:6379"},
		{":99", "localhost:99"},
		{"/var/redis/../run/redis.sock", "/var/run/redis.sock"},
	}
	for _, gold := range golden {
		if got := normalizeAddr(gold.Addr); got != gold.Normal {
			t.Errorf("got %q for %q, want %q", got, gold.Addr, gold.Normal)
		}
	}
}
