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

func Test_ValuePx(t *testing.T) {
	css := Parse("div .a { font-size: 45px}")

	crl := css.CssRuleList
	cr := crl[0]

	if cr.Type != STYLE_RULE {
		t.Errorf("cr.Type = %d , want div 1 .", cr.Type)
	}

	csr := cr.Style
	csd := csr.Styles
	v := csd["font-size"]
	if v.Value != "45px" {
		t.Errorf("v.Value = %s , want 45px .", v.Value)
	}
}

func Test_ValueEm(t *testing.T) {
	css := Parse("div .a { a: 45em}")

	crl := css.CssRuleList
	cr := crl[0]

	csr := cr.Style
	if cr.Type != STYLE_RULE {
		t.Errorf("cr.Type = %d , want div 1 .", cr.Type)
	}

	csd := csr.Styles
	v := csd["a"]
	if v.Value != "45em" {
		t.Errorf("v.Value = %s , want 45px .", v.Value)
	}

}

func Test_ValueRRGGBB(t *testing.T) {
	css := Parse("div .a { a: #123456}")

	crl := css.CssRuleList
	cr := crl[0]

	csr := cr.Style
	if cr.Type != STYLE_RULE {
		t.Errorf("cr.Type = %d , want div 1 .", cr.Type)
	}

	csd := csr.Styles
	v := csd["a"]
	if v.Value != "#123456" {
		t.Errorf("v.Value = %s , want 45px .", v.Value)
	}

}

func Test_ValueNumber(t *testing.T) {
	css := Parse("div .a { a: 456}")

	crl := css.CssRuleList
	cr := crl[0]

	csr := cr.Style
	if cr.Type != STYLE_RULE {
		t.Errorf("cr.Type = %d , want div 1 .", cr.Type)
	}

	csd := csr.Styles
	v := csd["a"]
	if v.Value != "456" {
		t.Errorf("v.Value = %s , want 45px .", v.Value)
	}
}

func Test_ValueInherit(t *testing.T) {
	css := Parse("div .a { a: inherit}")

	crl := css.CssRuleList
	cr := crl[0]

	csr := cr.Style
	if cr.Type != STYLE_RULE {
		t.Errorf("cr.Type = %d , want div 1 .", cr.Type)
	}

	csd := csr.Styles
	v := csd["a"]
	if v.Value != "inherit" {
		t.Errorf("v.Value = %s , want 45px .", v.Value)
	}
}

func Test_ValueRGBFunction(t *testing.T) {
	css := Parse(`div .a { 
					a: rgb(1,2,3)
					font-size: 150% !important
		 }`)
	crl := css.CssRuleList
	cr := crl[0]

	csr := cr.Style
	if cr.Type != STYLE_RULE {
		t.Errorf("cr.Type = %d , want div 1 .", cr.Type)
	}
	csd := csr.Styles
	v := csd["a"]
	if v.Value != "rgb(1,2,3)" {
		t.Errorf("v.Value = %s , want 45px .", v.Value)
	}

}
