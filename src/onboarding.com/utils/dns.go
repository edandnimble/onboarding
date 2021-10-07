package utils

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
)

func GetServiceDNS(serviceName string) (ip string, port string, err error) {
	//_, srvs, err := net.LookupSRV(serviceName, "tcp", "consul")
	_, srvs, err := net.LookupSRV("", "", serviceName+".service.consul")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		return "", "", err
	}
	if len(srvs) == 0 {
		return "", "", fmt.Errorf("No SRV record found")
	}

	r := rand.Intn(len(srvs))
	port = strconv.FormatUint(uint64(srvs[r].Port), 10)
	target := srvs[r].Target
	ips, err := net.LookupIP(target)
	if len(ips) != 1 {
		return "", "", fmt.Errorf("Too many IPs detected for record target: %v IPs: %v", target, ips)
	}
	ip = ips[0].String()

	fmt.Printf("get DNS service: %v ip %v port %v\n", serviceName, ip, port)
	return ip, port, nil
}
