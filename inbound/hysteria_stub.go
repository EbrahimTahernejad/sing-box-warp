//go:build !with_quic

package inbound

import (
	"context"

	"github.com/EbrahimTahernejad/sing-box-warp/adapter"
	C "github.com/EbrahimTahernejad/sing-box-warp/constant"
	"github.com/EbrahimTahernejad/sing-box-warp/log"
	"github.com/EbrahimTahernejad/sing-box-warp/option"
)

func NewHysteria(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.HysteriaInboundOptions) (adapter.Inbound, error) {
	return nil, C.ErrQUICNotIncluded
}

func NewHysteria2(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.Hysteria2InboundOptions) (adapter.Inbound, error) {
	return nil, C.ErrQUICNotIncluded
}
