/**
 * Created with IntelliJ IDEA.
 * User: chris
 * Date: 08/09/2013
 * Time: 20:57
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"time"
	"net"
	"os"
	"github.com/coreos/go-etcd/etcd"
)

var ttl uint64

func main() {

	nodename := "test"
	ttl = 100

	c := etcd.NewClient()

	for true {
		name, err := os.Hostname()
		if err != nil {
			os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		} else {

			addrs, err := net.InterfaceAddrs()

			if err != nil {
				os.Stderr.WriteString("Oops: " + err.Error() + "\n")
			} else {

				for _, a := range addrs {
					if ipnet, ok := a.(*net.IPNet); ok && ipnet.IP.IsGlobalUnicast() {
						addr := ipnet.IP.String()
						c.Set(nodename + "/" + name, addr, ttl)
					}
				}
			}
		}
		time.Sleep(time.Second*300)
	}
}
