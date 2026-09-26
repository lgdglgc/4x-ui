package sub

import (
	"regexp"
	"strings"
	"unicode"
)

type flagRule struct {
	flag     string
	keywords []string
}

var flagRules = []flagRule{
	// 中国香港
	{flag: "🇭🇰", keywords: []string{"香港", "港", "Hong Kong", "HongKong", "HKG", "HK"}},
	// 中国台湾
	{flag: "🇹🇼", keywords: []string{"台湾", "台北", "台中", "Taiwan", "Taipei", "TW"}},
	// 日本
	{flag: "🇯🇵", keywords: []string{"日本", "东京", "大阪", "Japan", "Tokyo", "Osaka", "JP"}},
	// 美国
	{flag: "🇺🇸", keywords: []string{"美西", "美东", "美国", "洛杉矶", "圣何塞", "西雅图", "芝加哥", "纽约", "硅谷", "波特兰", "达拉斯", "凤凰城", "USA", "United States", "America", "Los Angeles", "San Jose", "Seattle", "Chicago", "New York", "Silicon Valley", "Portland", "Dallas", "Phoenix", "US"}},
	// 新加坡
	{flag: "🇸🇬", keywords: []string{"新加坡", "狮城", "Singapore", "SG"}},
	// 韩国
	{flag: "🇰🇷", keywords: []string{"韩国", "首尔", "Korea", "Seoul", "KR"}},
	// 英国
	{flag: "🇬🇧", keywords: []string{"英国", "伦敦", "Britain", "United Kingdom", "London", "UK", "GB"}},
	// 德国
	{flag: "🇩🇪", keywords: []string{"德国", "法兰克福", "Germany", "Frankfurt", "DE"}},
	// 法国
	{flag: "🇫🇷", keywords: []string{"法国", "巴黎", "France", "Paris", "FR"}},
	// 荷兰
	{flag: "🇳🇱", keywords: []string{"荷兰", "阿姆斯特丹", "Netherlands", "Amsterdam", "NL"}},
	// 俄罗斯
	{flag: "🇷🇺", keywords: []string{"俄罗斯", "莫斯科", "Russia", "Moscow", "RU"}},
	// 加拿大
	{flag: "🇨🇦", keywords: []string{"加拿大", "多伦多", "温哥华", "Canada", "Toronto", "Vancouver", "CA"}},
	// 澳大利亚
	{flag: "🇦🇺", keywords: []string{"澳大利亚", "澳洲", "悉尼", "墨尔本", "Australia", "Sydney", "Melbourne", "AU"}},
	// 土耳其
	{flag: "🇹🇷", keywords: []string{"土耳其", "Turkey", "Istanbul", "TR"}},
	// 阿联酋 / 迪拜
	{flag: "🇦🇪", keywords: []string{"阿联酋", "迪拜", "Dubai", "AE"}},
	// 印度
	{flag: "🇮🇳", keywords: []string{"印度", "孟买", "India", "Mumbai", "IN"}},
	// 巴西
	{flag: "🇧🇷", keywords: []string{"巴西", "Brazil", "BR"}},
	// 马来西亚
	{flag: "🇲🇾", keywords: []string{"马来西亚", "大马", "吉隆坡", "Malaysia", "MY"}},
	// 泰国
	{flag: "🇹🇭", keywords: []string{"泰国", "曼谷", "Thailand", "Bangkok", "TH"}},
	// 越南
	{flag: "🇻🇳", keywords: []string{"越南", "Vietnam", "VN"}},
	// 菲律宾
	{flag: "🇵🇭", keywords: []string{"菲律宾", "Philippines", "PH"}},
	// 印度尼西亚
	{flag: "🇮🇩", keywords: []string{"印尼", "印度尼西亚", "Indonesia", "ID"}},
	// 中国大陆
	{flag: "🇨🇳", keywords: []string{"中国", "国内", "北京", "上海", "广州", "深圳", "China", "CN"}},
}

// reRandomSuffix matches random hashes appended by 3x-ui multi-user / default remark (e.g. -k7sc0bq19p)
var reRandomSuffix = regexp.MustCompile(`(?i)-[a-z0-9]{8,16}$`)

// knownProtocols are keywords that should not be stripped even if they end in a hyphen-tag
var knownProtocols = map[string]bool{
	"reality":     true,
	"trojan":      true,
	"shadowsocks": true,
	"wireguard":   true,
	"hysteria":    true,
	"hysteria2":   true,
	"vless":       true,
	"vmess":       true,
}

// BeautifyNodeName removes random user/client identifier suffixes and prepends the regional flag emoji.
func BeautifyNodeName(raw string) string {
	name := strings.TrimSpace(raw)
	if name == "" {
		return name
	}

	// 1. Remove random identifier suffix (e.g. -k7sc0bq19p)
	name = stripRandomIdentifier(name)

	// 2. Prepend country/region flag if not already present
	name = AddCountryFlag(name)

	return name
}

func stripRandomIdentifier(s string) string {
	loc := reRandomSuffix.FindStringIndex(s)
	if loc != nil {
		suffix := strings.ToLower(s[loc[0]+1:])
		// Avoid stripping protocol names or simple numbered segments (e.g. -01)
		if !knownProtocols[suffix] && hasBothLetterAndDigit(suffix) {
			return s[:loc[0]]
		}
	}
	return s
}

func hasBothLetterAndDigit(s string) bool {
	hasLetter := false
	hasDigit := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			hasLetter = true
		} else if unicode.IsDigit(r) {
			hasDigit = true
		}
	}
	return hasLetter && hasDigit
}

// AddCountryFlag prepends the corresponding flag emoji based on region keywords.
func AddCountryFlag(name string) string {
	name = strings.TrimSpace(name)
	if hasFlagEmoji(name) {
		return name
	}

	for _, rule := range flagRules {
		for _, kw := range rule.keywords {
			if containsKeyword(name, kw) {
				return rule.flag + " " + name
			}
		}
	}
	return name
}

func containsKeyword(name, kw string) bool {
	// If keyword contains non-ASCII (e.g. Chinese characters), substring match is safe
	for _, r := range kw {
		if r > 127 {
			return strings.Contains(name, kw)
		}
	}

	// For ASCII keywords (e.g. US, HK), use case-insensitive word-boundary matching
	pattern := "(?i)(^|[^a-zA-Z])" + regexp.QuoteMeta(kw) + "([^a-zA-Z]|$)"
	matched, err := regexp.MatchString(pattern, name)
	if err == nil && matched {
		return true
	}
	return false
}

func hasFlagEmoji(s string) bool {
	runes := []rune(strings.TrimSpace(s))
	if len(runes) >= 2 {
		if runes[0] >= 0x1F1E6 && runes[0] <= 0x1F1FF && runes[1] >= 0x1F1E6 && runes[1] <= 0x1F1FF {
			return true
		}
	}
	return false
}
