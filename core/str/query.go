package str

/*
These functions, although functional, feel a bit unpolished to me.
TODO: flesh out fuzzy string search functions.
*/

// func GetScoringFunction(N int, queries ...string) func(string) float32 {
// 	queryGrams, _ := GenNgrams(queries, N)
// 	n := N
// 	return func(str string) (score float32) {
// 		str = strings.ToLower(str)
//
// 		ngramLen := len(str) - n + 1
// 		ngrams := make(map[string]int, ngramLen)
// 		for i := 0; i < len(str)-n+1; i++ {
// 			gram := str[i : i+n]
// 			ngrams[gram]++
// 		}
//
// 		// check for ngrams
// 		for gram, weight := range queryGrams {
// 			if v, ok := ngrams[gram]; ok {
// 				score += (float32(v) * float32(weight)) / float32(len(queryGrams))
// 			}
// 		}
//
// 		return score
// 	}
// }
//
// func GenNgrams(sar []string, n int) (map[string]int, int) {
// 	var qlen int
// 	l := len(sar) - n + 1
// 	ngrams := make(map[string]int, l)
// 	for _, query := range sar {
// 		qlen += len(query)
// 		query = strings.ToLower(query)
// 		for i := 0; i < len(query)-n+1; i++ {
// 			gram := query[i : i+n]
// 			ngrams[gram]++
// 		}
// 	}
//
// 	return ngrams, qlen
// }
