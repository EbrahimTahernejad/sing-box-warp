//go:build with_quic

package include

import (
	_ "github.com/ebrahimtahernejad/sing-box-warp/transport/v2rayquic"
	_ "github.com/sagernet/sing-dns/quic"
)

const WithQUIC = true
