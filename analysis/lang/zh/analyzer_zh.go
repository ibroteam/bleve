package zh

import (
	"errors"
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

func NewAnalyzer(config map[string]interface{}, cache *registry.Cache) (analysis.Analyzer, error) {
	tokenizerName, ok := config["tokenizer"].(string)
	if !ok {
		return nil, errors.New("must specify tokenizer")
	}
	tokenizer, err := cache.TokenizerNamed(tokenizerName)
	if err != nil {
		return nil, err
	}
	alz := &analysis.DefaultAnalyzer{Tokenizer: tokenizer}
	return alz, nil
}
func init() {
	registry.RegisterAnalyzer(Name, NewAnalyzer)
}
