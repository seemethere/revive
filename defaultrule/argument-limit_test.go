package defaultrule

import (
	"testing"

	"github.com/mgechev/revive/rule"
	"github.com/mgechev/revive/testutil"
)

func TestArgumentLimit(t *testing.T) {
	t.Parallel()

	program := `
	package foo

	[@1]func foo(a int, b int, c int)[/@1] {
		return a + b + c;
	}
	`
	testutil.AssertFailures(t, program, &ArgumentsLimitRule{}, rule.Arguments{"2"})
}

func TestArgumentLimit2(t *testing.T) {
	t.Parallel()

	program := `
	package foo

	func foo(a int, b int, c int) {
		return a + b + c;
	}
	`
	testutil.AssertSuccess(t, program, &ArgumentsLimitRule{}, rule.Arguments{"3"})
}
