//go:build !with_quic

package outbound

import (
	"context"

	"github.com/ebrahimtahernejad/sing-box-warp/adapter"
	C "github.com/ebrahimtahernejad/sing-box-warp/constant"
	"github.com/ebrahimtahernejad/sing-box-warp/log"
	"github.com/ebrahimtahernejad/sing-box-warp/option"
)

func NewTUIC(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.TUICOutboundOptions) (adapter.Outbound, error) {
	return nil, C.ErrQUICNotIncluded
}
