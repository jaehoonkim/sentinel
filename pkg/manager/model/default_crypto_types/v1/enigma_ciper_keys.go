//go:generate go run github.com/abice/go-enum --file=enigma_ciper_keys.go --names --nocase=false
package v1

/*
	ENUM(

morpheus.default.crypto
)
*/
type CiperKey int
