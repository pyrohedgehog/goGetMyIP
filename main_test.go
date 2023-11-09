package goGetMyIP

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPGetter(t *testing.T) {
	apis := []string{
		"https://api.ipify.org",
		"http://myexternalip.com/raw",
		// "http://api.ident.me",
	}
	//setup a getter for each default endpoint, then cache our IP
	ips := []string{}
	for _, api := range apis {
		ipg := NewIPGetterWith([]string{api})
		ips = append(ips, ipg.GetIPString())
	}
	ips = append(ips, NewIPGetter().GetIPString())
	for i := 1; i < len(ips); i++ {
		assert.Equal(t,
			ips[i], ips[i-1],
			"different IPs from different endpoints",
		)
	}

}
