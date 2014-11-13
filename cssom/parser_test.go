package cssom

import (
	"testing"
)

func Test_WithoutImpotant(t *testing.T) {
	css := Parse("div .a { font-size: 150%}")

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
	v := csd["font-size"]
	if v.Value != "150%" {
		t.Errorf("v.Value = %s , want 150% .", v.Value)
	}

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
		t.Errorf("csr.SelectorText() = %s , want div .a .", csr.SelectorText)
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

func Test_Import(t *testing.T) {
	Parse("@import url('common.css');")
}
