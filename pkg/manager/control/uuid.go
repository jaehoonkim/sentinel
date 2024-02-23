package control

import "github.com/jaehoonkim/morpheus/pkg/manager/macro"

func genUuidString(uuid string) string {
	if 0 < len(uuid) {
		return uuid
	}
	return macro.NewUuidString()
}
