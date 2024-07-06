package dgsys

import (
	"encoding/binary"
	"errors"
	"net"
)

func GetLocalLanIps() []string {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if !ok || ipNet == nil {
			continue
		}

		// 只保存IPv4地址且不包含127开头的地址（如localhost）
		if ipNet.IP.To4() != nil && !ipNet.IP.IsLoopback() {
			ips = append(ips, ipNet.IP.String())
		}
	}

	return ips
}

func LocalLanIpToUint32() (uint32, error) {
	ips := GetLocalLanIps()
	if len(ips) == 0 {
		return 0, errors.New("no local LAN IP address found")
	}

	ip := net.ParseIP(ips[0])
	if ip == nil {
		return 0, errors.New("invalid IP address")
	}

	if ip4 := ip.To4(); ip4 != nil {
		return binary.BigEndian.Uint32(ip4), nil
	} else {
		return 0, errors.New("not an IPv4 address")
	}
}
