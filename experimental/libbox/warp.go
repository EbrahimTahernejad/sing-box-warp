package libbox

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ebrahimtahernejad/sing-box-warp/experimental/libbox/internal/warp"
	"github.com/ebrahimtahernejad/sing-box-warp/experimental/libbox/internal/ws"
	"github.com/ebrahimtahernejad/sing-box-warp/option"
)

func WarpSetupFree() error {
	// make primary identity
	license := "notset"
	_license := ""
	warp.UpdatePath("./warp-primary")
	if !warp.CheckProfileExists(license) {
		err := warp.LoadOrCreateIdentity(_license)
		if err != nil {
			log.Printf("error: %v", err)
			return fmt.Errorf("error: %v", err)
		}
	}
	// make secondary identity
	warp.UpdatePath("./warp-secondary")
	if !warp.CheckProfileExists(license) {
		err := warp.LoadOrCreateIdentity(_license)
		if err != nil {
			log.Printf("error: %v", err)
			return fmt.Errorf("error: %v", err)
		}
	}
	return nil
}

func convertConfig(device *ws.DeviceConfig) (*option.WireGuardOutboundOptions, error) {
	peers := []option.WireGuardPeer{}
	for _, peer := range device.Peers {
		address, port, found := strings.Cut(*peer.Endpoint, ":")
		if !found {
			return nil, fmt.Errorf("endpoint has no port")
		}
		portUInt64, err := strconv.ParseUint(port, 10, 16)
		if err != nil {
			return nil, err
		}
		ips := []string{}
		for _, allowedIP := range peer.AllowedIPs {
			ips = append(ips, allowedIP.String())
		}
		peers = append(peers, option.WireGuardPeer{
			AllowedIPs:   ips,
			PublicKey:    peer.PublicKey,
			PreSharedKey: peer.PreSharedKey,
			ServerOptions: option.ServerOptions{
				Server:     address,
				ServerPort: uint16(portUInt64),
			},
		})
	}

	return &option.WireGuardOutboundOptions{
		MTU:          uint32(device.MTU),
		Peers:        peers,
		PrivateKey:   device.SecretKey,
		LocalAddress: device.Endpoint,
	}, nil
}

func WarpGetOutbounds(tag string, endpoint string, nested bool) (string, error) {
	options := []option.Outbound{}
	primaryTag := tag
	if nested {
		primaryTag = "primary"
		conf, err := ws.ParseConfig("./warp-secondary/wgcf-profile.ini", endpoint)
		if err != nil {
			return "", err
		}
		wgOptions, err := convertConfig(conf.Device)
		if err != nil {
			return "", err
		}
		wgOptions.Detour = primaryTag
		options = append(options, option.Outbound{
			Type:             "wireguard",
			Tag:              tag,
			WireGuardOptions: *wgOptions,
		})
	}
	conf, err := ws.ParseConfig("./warp-primary/wgcf-profile.ini", endpoint)
	if err != nil {
		return "", err
	}
	wgOptions, err := convertConfig(conf.Device)
	if err != nil {
		return "", err
	}
	options = append(options, option.Outbound{
		Type:             "wireguard",
		Tag:              primaryTag,
		WireGuardOptions: *wgOptions,
	})
	jsonData, err := json.Marshal(options)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
