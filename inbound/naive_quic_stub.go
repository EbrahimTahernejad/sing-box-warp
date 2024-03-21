//go:build !with_quic

package inbound

import (
	C "github.com/ebrahimtahernejad/sing-box-warp/constant"
)

func (n *Naive) configureHTTP3Listener() error {
	return C.ErrQUICNotIncluded
}
