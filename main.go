package goGetMyIP

import (
	"io"
	"net/http"
)

// IPGetter handles getting the IP from a list of API endpoints.
type IPGetter struct {
	ipString  string
	apiPoints []string
}

// create an IPGetter with some default API's
func NewIPGetter() *IPGetter {
	return NewIPGetterWith(
		//some default API endpoints that I've used before
		[]string{
			"https://api.ipify.org",
			"http://myexternalip.com/raw",
			// "http://api.ident.me",
		},
	)
}

// create an IPGetter with your own endpoints.
func NewIPGetterWith(apis []string) *IPGetter {
	ipg := IPGetter{
		apiPoints: apis,
	}
	return &ipg
}

// the start of the process of getting the IP's. Useful to call at the start of a program.
// will go through all the API's given, and attempt to get the external IP from whichever is fastest.
// only really needed if you really want to get all of your processing out of the way at the launch of the program.
func (ipg *IPGetter) CacheIP() {
	cacheChan := make(chan string)
	for _, api := range ipg.apiPoints {
		go func(api string) {
			resp, err := http.Get(api)
			if err != nil {
				return
			}
			ip, err := io.ReadAll(resp.Body)
			if err != nil {
				return
			}
			defer resp.Body.Close()
			cacheChan <- string(ip)

		}(api)
	}
	ipg.ipString = <-cacheChan
}

// get the IP as a string. checks to see if it already has been cached before calling APIs
func (ipg *IPGetter) GetIPString() string {
	if ipg.ipString == "" {
		ipg.CacheIP()
	}
	return ipg.ipString
}

// gets the external IP of this device. Most suitable for 99% of use cases.
func GetExternalIP() string {
	return NewIPGetter().GetIPString()
}
