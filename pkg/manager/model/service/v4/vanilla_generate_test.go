package service_test

import (
	"os"
	"testing"

	"github.com/jaehoonkim/synapse/pkg/manager/database/vanilla/ice_cream_maker"
	"github.com/jaehoonkim/synapse/pkg/manager/model/service/v4"
)

var objs = []interface{}{
	service.Service{},

	service.ServiceResult{},

	service.ServiceStatus{},

	service.Service_polling{},
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
