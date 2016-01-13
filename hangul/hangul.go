// Package hangul provides basic functions for Hangul processing.
package hangul

var (
	start = rune(44032) // "가"의 유니코드 포인트
	end   = rune(55204) // "힣" 다음 글자의 유니코드 포인트
)

// HasConsonantSuffix returns true if s has Hangul consonant jamo at the end.
func HasConsonantSuffix(s string) bool {
	numEnds := 28
	result := false
	for _, r := range s {
		if start <= r && r < end {
			index := int(r - start)
			result = index%numEnds != 0
		}
	}
	return result
}
