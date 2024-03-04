package dgsys

import "net"

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
