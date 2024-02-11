package utils

import (
	"crypto/sha256"
	"fmt"
	"net"

	"github.com/axgle/mahonia"
	uuid "github.com/satori/go.uuid"
)

func GetInnetIPAddress() []string {
	var ips []string
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return ips
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						ips = append(ips, ipnet.IP.String())
					}
				}
			}
		}
	}
	return ips
}

func SHA256Check(password, sha string) bool {
	sum := sha256.Sum256([]byte(password))
	sha2 := fmt.Sprintf("%x", sum)
	//fmt.Println("sha2:", sha2, "sha:", sha)
	if sha2 != sha {
		return false
	}
	return true
}

func GetUUID() string {
	u1 := uuid.Must(uuid.NewV4(), nil)
	return u1.String()
}

func GBK2UTF8(str string) string {
	dec := mahonia.NewDecoder("GBK")
	return dec.ConvertString(str)
}

func UTF82GBK(str string) string {
	enc := mahonia.NewEncoder("GBK")
	return enc.ConvertString(str)

}
