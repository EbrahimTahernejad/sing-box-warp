package sniff

import (
	"context"
	"crypto/tls"
	"io"

	"github.com/ebrahimtahernejad/sing-box-warp/adapter"
	C "github.com/ebrahimtahernejad/sing-box-warp/constant"
	"github.com/sagernet/sing/common/bufio"
)

func TLSClientHello(ctx context.Context, reader io.Reader) (*adapter.InboundContext, error) {
	var clientHello *tls.ClientHelloInfo
	err := tls.Server(bufio.NewReadOnlyConn(reader), &tls.Config{
		GetConfigForClient: func(argHello *tls.ClientHelloInfo) (*tls.Config, error) {
			clientHello = argHello
			return nil, nil
		},
	}).HandshakeContext(ctx)
	if clientHello != nil {
		return &adapter.InboundContext{Protocol: C.ProtocolTLS, Domain: clientHello.ServerName}, nil
	}
	return nil, err
}
