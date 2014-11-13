package cssom

import (
	"testing"
)

func CSSStyleDeclarationTest(t *testing.T, csd map[string]*CSSStyleDeclaration, property string, value string, important int) bool {

	return true
}

func Test_WithoutImpotant(t *testing.T) {
	css := Parse("div .a { font-size: 150%}")

	crl := css.CssRuleList
	cr := crl[0]
	csr := cr.Style
	csd := csr.Styles
	v := csd["font-size"]
	if v.Important != 0 {
		t.Errorf("v.Important = %d , want 1 .", v.Important)
	}
}

func Test_WithImpotant(t *testing.T) {
	css := Parse("div .a { font-size: 150% !important}")

	crl := css.CssRuleList
	cr := crl[0]

	if cr.Type != STYLE_RULE {
		t.Errorf("cr.Type = %d , want div 1 .", cr.Type)
	}

	csr := cr.Style
	if csr.SelectorText != "div .a" {
		t.Errorf("csr.SelectorText() = z%sz , want div .a .", csr.SelectorText)
	}

	csd := csr.Styles
	v := csd["font-size"]
	if v.Value != "150%" {
		t.Errorf("v.Value = %s , want 150% .", v.Value)
	}

	if v.Important != 1 {
		t.Errorf("v.Important = %d , want 1 .", v.Important)
	}

}

func Test_MultipleDeclarations(t *testing.T) {
	css := Parse(`div .a {
				font-size1: 150%
				font-size2: 250%
			}`)
	css.Print()

	crl := css.CssRuleList
	cr := crl[0]

	if cr.Type != STYLE_RULE {
		t.Errorf("cr.Type = %d , want div 1 .", cr.Type)
	}

	csr := cr.Style
	if csr.SelectorText != "div .a" {
		t.Errorf("csr.SelectorText() = %s , want div .a .", csr.SelectorText)
	}

	csd := csr.Styles
	v := csd["font-size1"]
	if v.Value != "150%" {
		t.Errorf("v.Value = %s , want 150% .", v.Value)
	}

	if v.Important != 0 {
		t.Errorf("v.Important = %d , want 0 .", v.Important)
	}

	v = csd["font-size2"]
	if v.Value != "250%" {
		t.Errorf("v.Value = %s , want 150% .", v.Value)
	}

	if v.Important != 0 {
		t.Errorf("v.Important = %d , want 0 .", v.Important)
	}

}

func Test_MultipleSelectors(t *testing.T) {
	css := Parse(`div .a {
				font-size_a: 150%
			}
			p .b {
				font-size_b1: 250%
			}
			`)
	crl := css.CssRuleList
	cr := crl[0]

	if cr.Type != STYLE_RULE {
		t.Errorf("cr.Type = %d , want div 1 .", cr.Type)
	}

	csr := cr.Style
	if csr.SelectorText != "div .a" {
		t.Errorf("csr.SelectorText() = %s , want div .a .", csr.SelectorText)
	}

	cr = crl[1]

	if cr.Type != STYLE_RULE {
		t.Errorf("cr.Type = %d , want div 1 .", cr.Type)
	}

	csr = cr.Style
	if csr.SelectorText != "p .b" {
		t.Errorf("csr.SelectorText() = %s , want div .a .", csr.SelectorText)
	}
}

/*
func Test_Import(t *testing.T) {
	Parse("@import url('common.css');")
}
*/
