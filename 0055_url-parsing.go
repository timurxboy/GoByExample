package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)

	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)
	host, post, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(post)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

	s1 := "smtp://user:pass@hostname/domain"
	u1, err1 := url.Parse(s1)
	if err1 != nil {
		panic(err1)
	}

	fmt.Println(u1.Scheme)
	fmt.Println(u1.User)
	fmt.Println(u1.User.Username())
	p1, _ := u1.User.Password()
	fmt.Println(p1)
	fmt.Println(u1.Host)
}
