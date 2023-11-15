# goGetMyIP
just a small library for Golang projects to be able to go, and get your external IP. 

To be entirely honest, this is written because I needed to get the external IP at runtime a while ago, and this was the best solution I could find. If this is the best way to get the external IP, then it's useful to have a library with updated API endpoints. Otherwise, let Cunningham's Law work its magic!

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/pyrohedgehog/goGetMyIP/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/pyrohedgehog/goGetMyIP)](https://goreportcard.com/report/github.com/pyrohedgehog/goGetMyIP)

To install, run:
`go get github.com/pyrohedgehog/goGetMyIP`

for 99% of use cases, use 
```golang 
ipString := goGetMyIP.GetExternalIP()
```
