//go:build !with_shadowsocksr

package outbound

import (
	"context"

	"github.com/EbrahimTahernejad/sing-box-warp/adapter"
	"github.com/EbrahimTahernejad/sing-box-warp/log"
	"github.com/EbrahimTahernejad/sing-box-warp/option"
	E "github.com/sagernet/sing/common/exceptions"
)

func NewShadowsocksR(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.ShadowsocksROutboundOptions) (adapter.Outbound, error) {
	return nil, E.New("ShadowsocksR is deprecated and removed in sing-box 1.6.0")
}
