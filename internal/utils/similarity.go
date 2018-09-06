package utils

import (
	"math"
	"regexp"
	"strings"
)


// Use from https://play.golang.org/p/72sGLipfE7

// Cosine converts texts to vectors
// associatting each chracter with its frequncy
// and caculates cosine similarities.
// (https://en.wikipedia.org/wiki/Cosine_similarity)
func Cosine(txt1, txt2 []byte) float64 {
	vect1 := make(map[byte]int)
	for _, t := range txt1 {
		vect1[t]++
	}
	vect2 := make(map[byte]int)
	for _, t := range txt2 {
		vect2[t]++
	}
	//
	// dot-product two vectors
	// map[byte]int return 0 for non-existing key
	// and if two texts are equal, product will be highest
	// and if two texts are totally different, it will be 0
	//
	// to calculate A·B
	dotProduct := 0.0
	for k, v := range vect1 {
		dotProduct += float64(v) * float64(vect2[k])
	}
	// to calculate |A|*|B|
	sum1 := 0.0
	for _, v := range vect1 {
		sum1 += math.Pow(float64(v), 2)
	}
	sum2 := 0.0
	for _, v := range vect2 {
		sum2 += math.Pow(float64(v), 2)
	}
	magnitude := math.Sqrt(sum1) * math.Sqrt(sum2)
	if magnitude == 0 {
		return 0.0
	}
	return float64(dotProduct) / float64(magnitude)
}

func wordCount(s string) map[string]int {
	words := strings.Fields(s)
	counts := make(map[string]int, len(words))
	for _, word := range words {
		counts[word]++
	}
	return counts
}

func intersection(vec1, vec2 map[string]int) []string {
	var intersection []string

	for key, _ := range vec1 {
		if _, ok := vec2[key]; ok {
			intersection = append(intersection, key)
		}
	}
	return intersection
}

func TextToVector(sentence string) (map[string]int, error) {
	if reg, err := regexp.Compile("[^a-zA-Z0-9]+"); err != nil {
		return nil, err
	} else {
		processed := reg.ReplaceAllString(sentence, " ")
		wordCount(processed)
		return wordCount(processed), nil
	}
}

func ProcessedCosine(vec1, vec2 map[string]int) float64 {
	intersection := intersection(vec1, vec2)
	numerator := 0
	for _, v := range intersection {
		numerator = numerator + (vec1[v] * vec2[v])
	}

	sum1 := 0.0
	for _, v := range vec1 {
		sum1 += math.Pow(float64(v), 2)
	}
	sum2 := 0.0
	for _, v := range vec2 {
		sum2 += math.Pow(float64(v), 2)
	}
	denominator := math.Sqrt(sum1) * math.Sqrt(sum2)

	if denominator == 0 {
		return 0.0
	}
	return float64(numerator) / float64(denominator)
}
