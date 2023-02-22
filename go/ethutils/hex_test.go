package ethutils

import (
	"gotest.tools/assert"
	"testing"
)

func TestHexZeroPad(t *testing.T) {
	tests := []struct {
		name   string
		value  string
		length int
		expect string
	}{
		{
			name:   "threeandwo's value",
			value:  "0x7D0aAbA94724320A5155ba8670465b96d33f2696",
			length: 32,
			expect: "0x0000000000000000000000007D0aAbA94724320A5155ba8670465b96d33f2696",
		},
		{
			name:   "normal value",
			value:  "0xe2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2",
			length: 32,
			expect: "0x000000000000000000000000e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2",
		},
		{
			name:   "zero value",
			value:  "0x0",
			length: 32,
			expect: "0x0000000000000000000000000000000000000000000000000000000000000000",
		},
		{
			name:   "zero value",
			value:  "0x0",
			length: 64,
			expect: "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pad, err := HexZeroPad(tt.value, tt.length)
			if err != nil {
				t.Errorf(err.Error())
				return
			}
			assert.Equal(t, pad, tt.expect)
		})
	}
}
