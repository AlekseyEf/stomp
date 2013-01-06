package frame

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	// Regexp for heart-beat header value
	heartBeatRegexp = regexp.MustCompile("^[0-9]+,[0-9]+$")

	invalidHeartBeat = errors.New("invalid heart-beat")
)

const (
	// Maximum number of milliseconds that can be represented
	// in a time.Duration.
	maxMilliseconds = math.MaxInt64 / int64(time.Millisecond)
)

// ParseHeartBeat parses the value of a STOMP heart-beat entry and
// returns two time durations. Returns an error if the heart-beat
// value is not in the correct format, or if the time durations are
// too big to be represented by an unsigned, 64 bit integer.
func ParseHeartBeat(heartBeat string) (time.Duration, time.Duration, error) {
	if !heartBeatRegexp.MatchString(heartBeat) {
		return 0, 0, invalidHeartBeat
	}
	slice := strings.Split(heartBeat, ",")
	value1, err := strconv.ParseInt(slice[0], 10, 64)
	if err != nil {
		return 0, 0, invalidHeartBeat
	}
	value2, err := strconv.ParseInt(slice[1], 10, 64)
	if err != nil {
		return 0, 0, invalidHeartBeat
	}
	if value1 > maxMilliseconds || value2 > maxMilliseconds {
		return 0, 0, invalidHeartBeat
	}
	return time.Duration(value1) * time.Millisecond,
		time.Duration(value2) * time.Millisecond, nil
}
