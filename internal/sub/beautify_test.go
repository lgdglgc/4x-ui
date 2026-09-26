package sub

import (
	"testing"
)

func TestBeautifyNodeName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "美西CN2GIA-Reality+Tcp-01-k7sc0bq19p",
			expected: "🇺🇸 美西CN2GIA-Reality+Tcp-01",
		},
		{
			input:    "美西CN2GIA-Reality+Tcp-01",
			expected: "🇺🇸 美西CN2GIA-Reality+Tcp-01",
		},
		{
			input:    "香港CN2-02",
			expected: "🇭🇰 香港CN2-02",
		},
		{
			input:    "日本软银-03",
			expected: "🇯🇵 日本软银-03",
		},
		{
			input:    "新加坡原生-04",
			expected: "🇸🇬 新加坡原生-04",
		},
		{
			input:    "台湾Hinet-05",
			expected: "🇹🇼 台湾Hinet-05",
		},
		{
			input:    "德国法兰克福-06",
			expected: "🇩🇪 德国法兰克福-06",
		},
		{
			input:    "英国伦敦-07",
			expected: "🇬🇧 英国伦敦-07",
		},
		{
			input:    "🇺🇸 已有国旗节点-08",
			expected: "🇺🇸 已有国旗节点-08",
		},
	}

	for _, tc := range tests {
		actual := BeautifyNodeName(tc.input)
		if actual != tc.expected {
			t.Errorf("BeautifyNodeName(%q) = %q; want %q", tc.input, actual, tc.expected)
		}
	}
}
