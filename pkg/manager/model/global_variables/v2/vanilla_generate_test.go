package v2_test

import (
	"os"
	"testing"

	"github.com/jaehoonkim/synapse/pkg/manager/database/vanilla/ice_cream_maker"
	v2 "github.com/jaehoonkim/synapse/pkg/manager/model/global_variables/v2"
)

var objs = []interface{}{
	v2.GlobalVariables{},
}

func TestNoXormColumns(t *testing.T) {
	s, err := ice_cream_maker.GenerateParts(objs, ice_cream_maker.Ingredients)
	if err != nil {
		t.Fatal(err)
	}

	println(s)

	if true {
		filename := "vanilla_generated.go"
		fd, err := os.Create(filename)
		if err != nil {
			t.Fatal(err)
		}

		if _, err = fd.WriteString(s); err != nil {
			t.Fatal(err)
		}
	}
}
