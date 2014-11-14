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
			}

			if context.State == STATE_DECLARE_BLOCK {
				if token.Value == "important" {
					context.NowImportant = 1
				} else if token.Value == "inherit" {
					context.NowValue = token.Value
				} else {
					if context.NowValue != "" {
						csd := &CSSStyleDeclaration{
							Value: strings.Trim(context.NowValue, " "),

							Important: context.NowImportant,
						}

						context.CurrentRule.Style.Styles[context.NowProperty] = csd

						context.NowProperty = ""
						context.NowValue = ""
						context.NowImportant = 0
					}

					context.NowProperty = strings.Trim(token.Value, " \t\n")
				}
			}

		case scanner.TokenS:
			println("scanner.TokenS:" + token.Value)
			if context.State == STATE_SELECTOR {
				context.NowSelectorText += token.Value
			}

			if context.State == STATE_DECLARE_BLOCK {
				context.NowProperty += strings.Trim(token.Value, " \t\n")
			}

		case scanner.TokenChar:
			if string(':') == token.Value {
				break
			}
			if string('!') == token.Value {
				break
			}
			println("scanner.TokenChar:" + token.Value)

			if string('{') == token.Value {
				context.State = STATE_DECLARE_BLOCK
				context.CurrentRule = NewStyleRule()
				context.CurrentRule.Style.SelectorText = strings.Trim(context.NowSelectorText, " ")

			} else if string('}') == token.Value {
				csd := &CSSStyleDeclaration{
					Value:     strings.Trim(context.NowValue, " "),
					Important: context.NowImportant,
				}
				context.CurrentRule.Style.Styles[strings.Trim(context.NowProperty, " ")] = csd
				css.CssRuleList = append(css.CssRuleList, context.CurrentRule)

				context.NowSelectorText = ""
				context.NowProperty = ""
				context.NowValue = ""
				context.NowImportant = 0

				context.State = STATE_NONE
			} else if context.State == STATE_DECLARE_BLOCK {
				context.NowValue += token.Value
			} else {
				if context.State == STATE_SELECTOR {
					context.NowSelectorText += token.Value
				}
			}
		case scanner.TokenPercentage:
			fallthrough
		case scanner.TokenDimension:
			fallthrough
		case scanner.TokenHash:
			fallthrough
		case scanner.TokenNumber:
			fallthrough
		case scanner.TokenFunction:
			fallthrough
		case scanner.TokenIncludes:
			fallthrough
		case scanner.TokenDashMatch:
			fallthrough
		case scanner.TokenPrefixMatch:
			fallthrough
		case scanner.TokenSuffixMatch:
			fallthrough
		case scanner.TokenSubstringMatch:

			if context.State == STATE_DECLARE_BLOCK {
				context.NowValue += strings.Trim(token.Value, " ")

			}
		}
	}
	return css
}
