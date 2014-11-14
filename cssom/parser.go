package cssom

import (
	"github.com/gorilla/css/scanner"
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
		if token.Type == scanner.TokenEOF || token.Type == scanner.TokenError {
			break
		}
		switch token.Type {
		case scanner.TokenAtKeyword:
			println("scanner.TokenAtKeyword:" + token.Value)
		case scanner.TokenString:
			println("scanner.TokenString:" + token.Value)
		case scanner.TokenURI:
			println("scanner.TokenURI:" + token.Value)
		case scanner.TokenUnicodeRange:
			println("scanner.TokenUnicodeRange:" + token.Value)
		case scanner.TokenCDO:
			println("scanner.TokenCDO:" + token.Value)
		case scanner.TokenCDC:
			println("scanner.TokenCDC:" + token.Value)
		case scanner.TokenComment:
			println("scanner.TokenComment:" + token.Value)
		case scanner.TokenIdent:
			println("scanner.TokenIdent:" + token.Value)

			if context.State == STATE_NONE || context.State == STATE_SELECTOR {

				context.State = STATE_SELECTOR
				context.NowSelectorText += strings.Trim(token.Value, " ")
				break
			}

			if context.State == STATE_DECLARE_BLOCK {
				context.State = STATE_PROPERTY
				context.NowProperty = strings.Trim(token.Value, " \t\n")
				break
			}

			if context.State == STATE_VALUE {
				if token.Value == "important" {
					context.NowImportant = 1
				} else {
					context.NowValue = token.Value
				}
				break
			}

		case scanner.TokenS:
			println("scanner.TokenS:" + token.Value)
			if context.State == STATE_SELECTOR {
				context.NowSelectorText += token.Value
			}

		case scanner.TokenChar:
			println("scanner.TokenChar:" + token.Value)

			if context.State == STATE_SELECTOR {
				if string('{') == token.Value {

					context.State = STATE_DECLARE_BLOCK

					context.CurrentRule = NewStyleRule()
					context.CurrentRule.Style.SelectorText = strings.Trim(context.NowSelectorText, " ")
					break
				} else {
					context.NowSelectorText += token.Value
				}
				break
			}

			if context.State == STATE_PROPERTY {
				if token.Value == ":" {
					context.State = STATE_VALUE
				}
				break
			}

			if context.State == STATE_DECLARE_BLOCK {
				if token.Value == "}" {
					css.CssRuleList = append(css.CssRuleList, context.CurrentRule)

					context.NowSelectorText = ""
					context.NowProperty = ""
					context.NowValue = ""
					context.NowImportant = 0

					context.State = STATE_NONE

				}
				break
			}

			if context.State == STATE_VALUE {
				if token.Value == ";" {
					csd := &CSSStyleDeclaration{
						Value:     strings.Trim(context.NowValue, " "),
						Important: context.NowImportant,
					}

					context.CurrentRule.Style.Styles[context.NowProperty] = csd

					context.NowProperty = ""
					context.NowValue = ""
					context.NowImportant = 0

					context.State = STATE_DECLARE_BLOCK
				} else {
					if token.Value != "!" {
						context.NowValue += token.Value
					}

				}
				break

			}
		case scanner.TokenPercentage:
			println("scanner.TokenPercentage:" + token.Value)
			fallthrough
		case scanner.TokenDimension:
			println("scanner.TokenDimension:" + token.Value)
			fallthrough
		case scanner.TokenHash:
			println("scanner.TokenHash:" + token.Value)
			fallthrough
		case scanner.TokenNumber:
			println("scanner.TokenNumber:" + token.Value)
			fallthrough
		case scanner.TokenFunction:
			println("scanner.TokenFunction:" + token.Value)
			fallthrough
		case scanner.TokenIncludes:
			println("scanner.TokenIncludes:" + token.Value)
			fallthrough
		case scanner.TokenDashMatch:
			println("scanner.TokenDashMatch:" + token.Value)
			fallthrough
		case scanner.TokenPrefixMatch:
			println("scanner.TokenPrefixMatch:" + token.Value)
			fallthrough
		case scanner.TokenSuffixMatch:
			println("scanner.TokenSuffixMatch:" + token.Value)
			fallthrough
		case scanner.TokenSubstringMatch:
			println("scanner.TokenSubstringMatch:" + token.Value)
			if context.State == STATE_VALUE {
				context.NowValue += strings.Trim(token.Value, " ")

			}
		}
	}
	return css
}
