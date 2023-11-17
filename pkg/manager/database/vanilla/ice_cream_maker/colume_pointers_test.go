package ice_cream_maker_test

import (
	"testing"

	"github.com/jaehoonkim/synapse/pkg/manager/database/vanilla/ice_cream_maker"
)

func TestColumnPointers(t *testing.T) {

	objs := []interface{}{
		ServiceStep_essential{},
		ServiceStep{},
	}

	s, err := ice_cream_maker.ColumnPtrs(objs...)
	if err != nil {
		t.Fatal(err)
	}

	println(s)
}
