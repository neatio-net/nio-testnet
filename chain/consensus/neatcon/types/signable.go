package types

import (
	"bytes"
	"io"

	. "github.com/nio-net/common"
	"github.com/nio-net/merkle"
)

type Signable interface {
	WriteSignBytes(chainID string, w io.Writer, n *int, err *error)
}

func SignBytes(chainID string, o Signable) []byte {
	buf, n, err := new(bytes.Buffer), new(int), new(error)
	o.WriteSignBytes(chainID, buf, n, err)
	if *err != nil {
		PanicCrisis(err)
	}
	return buf.Bytes()
}

func HashSignBytes(chainID string, o Signable) []byte {
	return merkle.SimpleHashFromBinary(SignBytes(chainID, o))
}
