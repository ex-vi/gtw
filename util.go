package gtw

import "time"

func StringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func BoolValue(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

func IntValue(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

func Float64Value(f *float64) float64 {
	if f == nil {
		return float64(0)
	}
	return *f
}

func TimeValue(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}
