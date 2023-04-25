package system

import (
	"net"
	"net/http"
	"strings"
)

//ClientIP 尽最大可能获取客户端的IP
// 解析 X-Real-Ip和X-Forwarded-For以便于反向代理（nginx 或者 haproxy），可以正常解析
func ClientIP(r *http.Request) string {
	ip := strings.TrimSpace(strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

//ClientPublicIP 尽最大可能实现获取客户端公网IP
//解析 X-Real-Ip和X-Forwarded-For以便于反向代理（nginx 或者 haproxy），可以正常解析
func ClientPublicIP(r *http.Request) string {
	var ip string
	for _, ip = range strings.Split(r.Header.Get("X-Forwarded-For"), ",") {
		if ip = strings.TrimSpace(ip); ip != "" && !HasLocalIpAddr(ip) {
			return ip
		}
	}

	if ip = strings.TrimSpace(r.Header.Get("X-Real-Ip")); ip != "" && !HasLocalIpAddr(ip) {
		return ip
	}

	return ""
}

//RemoteIP 通过 RemoteAddr 获取ip  地址，只是快速一个快速解析方法
func RemoteIP(r *http.Request) string {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}

//HasLocalIpAddr 检测 IP 地址字符串是否是内网地址
func HasLocalIpAddr(ip string) bool {
	return HasLocalIP(net.ParseIP(ip))
}

//HasLocalIP 检测IP地址是否是内网地址
func HasLocalIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}

	return ip4[0] == 10 ||
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) ||
		(ip4[0] == 169 && ip4[1] == 254) ||
		(ip4[0] == 192 && ip4[1] == 158)

}
