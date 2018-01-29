package utils

import (
	"testing"
	"time"
)

// Test Function TimeTrack
func TestTimeTrack(t *testing.T) {
	TimeTrack(time.Now(), "Test")
}