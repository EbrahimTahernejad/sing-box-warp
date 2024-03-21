//go:build !with_quic

package outbound

import (
	"context"

	"github.com/ebrahimtahernejad/sing-box-warp/adapter"
	C "github.com/ebrahimtahernejad/sing-box-warp/constant"
	"github.com/ebrahimtahernejad/sing-box-warp/log"
	"github.com/ebrahimtahernejad/sing-box-warp/option"
)

func NewHysteria(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.HysteriaOutboundOptions) (adapter.Outbound, error) {
	return nil, C.ErrQUICNotIncluded
}

func NewHysteria2(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.Hysteria2OutboundOptions) (adapter.Outbound, error) {
	return nil, C.ErrQUICNotIncluded
}
