package ethutils

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gotest.tools/assert"
	"testing"
)

func TestKeccak256(t *testing.T) {
	tests := []struct {
		name   string
		params [2]string
		expect string
	}{
		{
			name: "normal: threeandwo's address",
			params: [2]string{
				"0x0000000000000000000000007D0aAbA94724320A5155ba8670465b96d33f2696",
				"0x0000000000000000000000000000000000000000000000000000000000000000",
			},
			expect: "0x5bbc208f7564fe53a7299642bf4f36df6a7f10c7c4943d886276889c01bda6cc",
		},
		{
			name: "normal: e2 addr",
			params: [2]string{
				"0x000000000000000000000000e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2",
				"0x0000000000000000000000000000000000000000000000000000000000000000",
			},
			expect: "0xedd7d04419e9c48ceb6055956cbb4e2091ae310313a4d1fa7cbcfe7561616e03",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p1, _ := hexutil.Decode(tt.params[0])
			p2, _ := hexutil.Decode(tt.params[1])

			keccak256 := Keccak256(p1, p2)
			assert.Equal(t, "0x"+keccak256, tt.expect)
		})
	}

}
