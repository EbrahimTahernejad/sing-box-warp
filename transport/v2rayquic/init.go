//go:build with_quic

package v2rayquic

import "github.com/ebrahimtahernejad/sing-box-warp/transport/v2ray"

func init() {
	v2ray.RegisterQUICConstructor(NewServer, NewClient)
}
