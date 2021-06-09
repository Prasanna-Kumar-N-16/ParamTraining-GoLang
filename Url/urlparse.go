package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	p := fmt.Println
	s := "Mysql://admin:password@serverhost.com:8081/server/page1?key=value&key2=value2#X"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	p(u.Scheme)
	p(u.User.Username())
	p(u.User)
	pass, _ := u.User.Password()
	p(pass)
	p(u.Host)                                  //prints host and port
	host, port, _ := net.SplitHostPort(u.Host) //splits host name and port
	p(host)                                    //prints host
	p(port)
	p(u.Path)     //prints the path
	p(u.Fragment) //prints fragment path value
	p(u.RawQuery)

}
