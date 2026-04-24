package scanner


import (
	"net"
)


func inc(ip net.IP) {
	ip4 := ip.To4()

	for i := len(ip4) -1; i >= 0; i-- {
		ip4[i]++
		if ip4[i] != 0{
			return
		}
	}
}

func GenerateIP(cidr string) ([]net.IP,error) {
	var ip net.IP
	var ipnet *net.IPNet
	var err error
	var result []net.IP

	ip, ipnet, err = net.ParseCIDR(cidr)

	if err != nil {
		ip = net.ParseIP(cidr)
		if ip == nil || ip.To4() == nil {
			return nil, err
		}
		ipnet = &net.IPNet{
			IP: ip,
			Mask: net.CIDRMask(32,32),
		}
		
	}
	
    ip4 := ip.To4()

	network := ip4.Mask(ipnet.Mask)
	broadcast := make(net.IP,len(network))
	for i := range network{
		broadcast[i] = network[i] | ^ipnet.Mask[i]
	}
	if network.Equal(broadcast) {
		result = append(result, network)
		return result, nil
	}

	current := make(net.IP, len(network))
	copy(current, network)

	for  {
		inc(current)
		if !ipnet.Contains(current)	|| current.Equal(broadcast){
			break
		}
		result = append(result, append(net.IP(nil), current...))
	}
	return result,nil
}