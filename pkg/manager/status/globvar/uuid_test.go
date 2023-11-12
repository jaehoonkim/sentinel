package globvar_test

import (
	"testing"

	"github.com/NexClipper/sudory/pkg/manager/macro"
)

func TestGenerateUuid(t *testing.T) {

	for i := 0; i < 100; i++ {

		println(macro.NewUuidString())
	}

}
