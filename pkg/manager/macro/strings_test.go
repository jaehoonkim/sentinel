package macro_test

import (
	"testing"

	. "github.com/jaehoonkim/synapse/pkg/manager/macro"
)

func TestStringJoin(t *testing.T) {

	adder, builder := StringBuilder()

	for i := 0; i < 10; i++ {
		adder(i)
	}

	t.Log(builder(","))

}
