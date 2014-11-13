package cssom

import (
	//"log"
	"testing"
)

func Test_NewStyleRule(t *testing.T) {
	csl := NewStyleRule()

	if STYLE_RULE != csl.GetType() {
		t.Errorf("GetType() = %d , want STYLE_RULE.", csl.GetType())

	}
}
