package goGetMyIP

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPGetter(t *testing.T) {
	//setup a getter for each default endpoint, then cache our IP
	ips := []string{}
	for _, api := range defaultAPIEndpoints {
		ipg := NewIPGetterWith([]string{api})
		ips = append(ips, ipg.GetIPString())
	}
	ips = append(ips, NewIPGetter().GetIPString())
	ips = append(ips, GetExternalIP())
	for i := 1; i < len(ips); i++ {
		assert.Equal(t,
			ips[i], ips[i-1],
			"different IPs from different endpoints",
		)
	}

}
