package sniff

import (
	std_bufio "bufio"
	"context"
	"io"

	"github.com/EbrahimTahernejad/sing-box-warp/adapter"
	C "github.com/EbrahimTahernejad/sing-box-warp/constant"
	M "github.com/sagernet/sing/common/metadata"
	"github.com/sagernet/sing/protocol/http"
)

func HTTPHost(ctx context.Context, reader io.Reader) (*adapter.InboundContext, error) {
	request, err := http.ReadRequest(std_bufio.NewReader(reader))
	if err != nil {
		return nil, err
	}
	return &adapter.InboundContext{Protocol: C.ProtocolHTTP, Domain: M.ParseSocksaddr(request.Host).AddrString()}, nil
}
