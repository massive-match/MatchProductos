package models

import (
	"github.com/nicolasvasquez/matchProducts/internal/utils"
)

type Products struct {
	Title string
}

func (product *Products) TitleCosine (similar Products) float64 {
	return utils.Cosine([]byte(product.Title), []byte(similar.Title))
}

func (product *Products) TitleProcessedCosine (similar Products) (float64, error) {
	if vec1, err := utils.TextToVector(product.Title); err != nil {
		return 0.0, err
	} else {
		if vec2, err := utils.TextToVector(similar.Title); err != nil {
			return 0.0, err
		} else {
			return utils.ProcessedCosine(vec1, vec2), nil
		}
	}
}
