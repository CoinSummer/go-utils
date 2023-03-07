package string

import (
	"encoding/json"
	"gotest.tools/assert"
	"testing"
)

func TestSplitNum(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{
			name:  "normal: remix",
			value: "0-2,2-3,233,244",
			want:  true,
		},
		{
			name:  "normal: single dash",
			value: "0-9,25-35",
			want:  true,
		},
		{
			name:  "normal: comma",
			value: "233,244",
			want:  true,
		},
		{
			name:  "error: dash",
			value: "-",
			want:  true,
		},
		{
			name:  "error: multi dash",
			value: "0--25",
			want:  true,
		},
		{
			name:  "error: not exists end index",
			value: "0-",
			want:  true,
		},
		{
			name:  "error: not exists start index",
			value: "-25",
			want:  true,
		},
		{
			name:  "normal: redundant commas.",
			value: "0-25,,",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numMap, err := SplitNum(tt.value)
			if (err != nil) == tt.want {
				assert.Error(t, err, "split num error")
				return
			}
			marshal, _ := json.Marshal(numMap)
			t.Logf(string(marshal))
		})
	}
}
