package cssom

import (
	"github.com/gorilla/css/scanner"
	"log"
	"strings"
)

type State int

const (
	STATE_NONE State = iota
	STATE_SELECTOR
	STATE_DECLARE_BLOCK
	STATE_PROPERTY
	STATE_VALUE
)

type ParserContext struct {
	State           State
	NowSelectorText string
	NowProperty     string
	NowValue        string
	NowImportant    int

	CurrentRule *CSSRule
}

func Parse(input string) *CSSStyleSheet {
	context := &ParserContext{
		State:           STATE_NONE,
		NowSelectorText: "",
		NowProperty:     "",
		NowValue:        "",
		NowImportant:    0,
	}

	css := &CSSStyleSheet{}
	css.CssRuleList = make([]*CSSRule, 0)
	s := scanner.New(input)

	for {
		token := s.Next()
		//log.Println(token)

		if token.Type == scanner.TokenEOF || token.Type == scanner.TokenError {
			break
		}
		switch token.Type {
		case scanner.TokenIdent:
			if context.State == STATE_NONE || context.State == STATE_SELECTOR {
				context.State = STATE_SELECTOR
				context.NowSelectorText += token.Value
			}

			if context.State == STATE_DECLARE_BLOCK {
				log.Println("NowProperty:" + token.Value)

				if token.Value == "important" {
					context.NowImportant = 1
				} else {
					context.NowProperty += token.Value
				}

			}

		case scanner.TokenS:
			if context.State == STATE_SELECTOR {
				context.NowSelectorText += token.Value
			}

			if context.State == STATE_DECLARE_BLOCK {
				context.NowProperty += token.Value
			}

		case scanner.TokenChar:

			if string('{') == token.Value {
				context.State = STATE_DECLARE_BLOCK
				context.CurrentRule = NewStyleRule()
				context.CurrentRule.Style.SelectorText = context.NowSelectorText
				log.Println(context.CurrentRule.Style.SelectorText + "{")
			} else if string('}') == token.Value {
				context.State = STATE_NONE

				log.Println(strings.Trim(context.NowProperty, " "))
				log.Println(context.NowImportant)
				log.Println(strings.Trim(context.NowValue, " "))

				crl := css.GetCSSRuleList()

				csr := NewStyleRule()
				csr.Style.SelectorText = strings.Trim(context.NowSelectorText, " ")

				csd := &CSSStyleDeclaration{
					Value:     strings.Trim(context.NowValue, " "),
					Important: context.NowImportant,
				}

				csr.Style.Styles[strings.Trim(context.NowProperty, " ")] = csd
				css.CssRuleList = append(crl, csr)

				log.Println("}")

			} else {
				if context.State == STATE_SELECTOR {
					context.NowSelectorText += token.Value
				}
			}
		case scanner.TokenPercentage:
			if context.State == STATE_DECLARE_BLOCK {
				context.NowValue = token.Value
			}
		}

		//css.Print()

		//log.Println(rulelist[0].CssType)
	}
	return css
}
