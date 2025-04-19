package blume

// import (
// 	"testing"
// 	"github.com/stretchr/testify/assert"
// )
//
// func TestGet(t *testing.T) {
// 	// Test basic regex matching
// 	input := "hello world"
// 	getWord := Get(Rgx[string](`\w+`))
// 	assert.Equal(t, "helloworld", getWord(input), "Should extract all words")
//
// 	// Test multiple selectors
// 	getDigitsOrVowels := Get(Rgx[string](`[\d]+`), Rgx[string](`[aeiou]`))
// 	input = "h3ll0 woerld"
// 	assert.Equal(t, "30oe", getDigitsOrVowels(input), "Should extract all digits and vowels")
//
// 	// Test no matches
// 	getZ := Get(Rgx[string](`z+`))
// 	assert.Equal(t, "", getZ(input), "Should return empty string when no matches")
//
// 	// Test empty string input
// 	assert.Equal(t, "", getWord(""), "Should handle empty input")
// }
//
// func TestNth(t *testing.T) {
// 	input := "apple banana cherry date"
//
// 	// Test positive indices
// 	getFirstWord := Nth[string](0, Rgx[string](`\w+`))
// 	assert.Equal(t, "apple", getFirstWord(input), "Should extract first word")
//
// 	getSecondWord := Nth[string](1, Rgx[string](`\w+`))
// 	assert.Equal(t, "banana", getSecondWord(input), "Should extract second word")
//
// 	getLastWord := Nth[string](3, Rgx[string](`\w+`))
// 	assert.Equal(t, "date", getLastWord(input), "Should extract last word")
//
// 	// Test negative indices
// 	getLastWordNeg := Nth[string](-1, Rgx[string](`\w+`))
// 	assert.Equal(t, "date", getLastWordNeg(input), "Should extract last word with negative index")
//
// 	getSecondToLastWord := Nth[string](-2, Rgx[string](`\w+`))
// 	assert.Equal(t, "cherry", getSecondToLastWord(input), "Should extract second-to-last word")
//
// 	// Test out of bounds indices
// 	getTooLargeIndex := Nth[string](10, Rgx[string](`\w+`))
// 	assert.Equal(t, "", getTooLargeIndex(input), "Should return empty for out of bounds index")
//
// 	getTooNegativeIndex := Nth[string](-10, Rgx[string](`\w+`))
// 	assert.Equal(t, "", getTooNegativeIndex(input), "Should return empty for too negative index")
//
// 	// Test multiple selectors with interleaved matches
// 	input = "a1 b2 c3 d4"
// 	letterSelector := Rgx[string](`[a-d]`)
// 	digitSelector := Rgx[string](`\d`)
//
// 	getFirstMatch := Nth[string](0, letterSelector, digitSelector)
// 	assert.Equal(t, "a", getFirstMatch(input), "Should extract first match (a)")
//
// 	getSecondMatch := Nth[string](1, letterSelector, digitSelector)
// 	assert.Equal(t, "1", getSecondMatch(input), "Should extract second match (1)")
//
// 	getThirdMatch := Nth[string](2, letterSelector, digitSelector)
// 	assert.Equal(t, "b", getThirdMatch(input), "Should extract third match (b)")
//
// 	// Test no matches
// 	getNoMatches := Nth[string](0, Rgx[string](`z+`))
// 	assert.Equal(t, "", getNoMatches(input), "Should return empty when no matches")
// }
