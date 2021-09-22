package zh

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
	"github.com/go-ego/gse"
)

const (
	Name = "gse"
)

type GseTokenizer struct {
	segmenter *gse.Segmenter
}

func NewGseTokenizer() *GseTokenizer {
	var segmenter gse.Segmenter
	segmenter.SkipLog = true
	segmenter.LoadDictEmbed()
	segmenter.LoadStopEmbed()
	return &GseTokenizer{&segmenter}
}

func (t *GseTokenizer) Tokenize(sentence []byte) analysis.TokenStream {
	segments := t.segmenter.Segment(sentence)
	sz := len(segments)
	if sz == 0 {
		return analysis.TokenStream{}
	}

	pos := 1
	result := make(analysis.TokenStream, 0, sz)

	for _, seg := range segments {
		token := analysis.Token{
			Term:     []byte(seg.Token().Text()),
			Start:    seg.Start(),
			End:      seg.End(),
			Position: pos,
			Type:     analysis.Ideographic,
		}
		result = append(result, &token)
		pos++
	}
	return result
}

func tokenizerConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.Tokenizer, error) {
	return NewGseTokenizer(), nil
}

func init() {
	registry.RegisterTokenizer(Name, tokenizerConstructor)
}
