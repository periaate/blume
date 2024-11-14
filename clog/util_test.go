package clog

import "testing"

func TestHumanizeBytes(t *testing.T) {
	tests := []struct {
		name     string
		val      int
		decimals int
		asKiB    bool
		expected string
	}{
		{"Zero value", 0, 2, false, "0.00 B"},
		{"Small byte value no decimals", 512, 0, false, "512 B"},
		{"Exact one KiB", 1024, 2, true, "1.00 KiB"},
		{"Between KiB and MiB", 51200, 2, true, "50.00 KiB"},
		{"One MB", 1000000, 1, false, "1.0 MB"},
		{"Between KB and MB", 51200, 1, false, "51.2 KB"},
		{"Negative value", -1024, 2, true, "-1.00 KiB"},
		{"Large value", 1073741824, 2, true, "1.00 GiB"},
		{"Very large decimal", 1000000000, 3, false, "1.000 GB"},
		{"High precision", 123456789, 5, false, "123.45679 MB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HumanizeBytes(0, tt.val, tt.decimals, tt.asKiB)
			if result != tt.expected {
				t.Errorf("HumanizeBytes(%d, %d, %t) = %s; expected %s", tt.val, tt.decimals, tt.asKiB, result, tt.expected)
			}
		})
	}
}
