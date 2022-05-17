// Code generated by rlpgen. DO NOT EDIT.

//go:build !norlpgen
// +build !norlpgen

package types

import "github.com/CortexFoundation/CortexTheseus/rlp"
import "io"

func (obj *Meta) EncodeRLP(_w io.Writer) error {
	w := rlp.NewEncoderBuffer(_w)
	_tmp0 := w.List()
	w.WriteString(obj.Comment)
	w.WriteBytes(obj.Hash[:])
	w.WriteUint64(obj.RawSize)
	_tmp1 := w.List()
	for _, _tmp2 := range obj.Shape {
		w.WriteUint64(_tmp2)
	}
	w.ListEnd(_tmp1)
	if obj.BlockNum.Sign() == -1 {
		return rlp.ErrNegativeBigInt
	}
	w.WriteBigInt(&obj.BlockNum)
	w.ListEnd(_tmp0)
	return w.Flush()
}