package cssom

import (
	"fmt"
)

const (
	STYLE_RULE   = 1
	CHARSET_RULE = 2
)

type CSSStyleSheet struct {
	Type        string
	CssRuleList []*CSSRule
}

type CSSRule struct {
	Type  int
	Style CSSStyleRule
}

type CSSStyleRule struct {
	SelectorText string
	Styles       map[string]*CSSStyleDeclaration
}

type CSSStyleDeclaration struct {
	Property  string
	Value     string
	Important int
}

func (rcv *CSSRule) SetType(Type int) {
	rcv.Type = Type
}

func (rcv *CSSRule) GetType() int {
	return rcv.Type
}

func NewStyleRule() *CSSRule {
	crl := &CSSRule{
		Type: STYLE_RULE,
	}
	crl.Style.Styles = make(map[string]*CSSStyleDeclaration)

	return crl
}

func (s *CSSStyleSheet) GetCSSRuleList() []*CSSRule {
	return s.CssRuleList
}

func (rcv *CSSStyleRule) Print() {
	fmt.Printf("SelectorText:%s\n", rcv.SelectorText)
	for k, v := range rcv.Styles {
		fmt.Printf("Property:%s Value:%s important:%d\n", k, v.Value, v.Important)
	}
}

func (s *CSSStyleSheet) Print() {
	for _, cr := range s.GetCSSRuleList() {
		fmt.Printf("type:%d\n", STYLE_RULE)

		if cr.Type == STYLE_RULE {
			cr.Style.Print()
		}
	}
}
