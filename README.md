
do some operation on ip


### Convert an ipv6 network and an ipv4 into 6rd cisco compatible
```
$ ./go-ipcalc convert 2001:0101:FFC0::/47 151.151.100.4
2001:101:ffc0:c808::/63
```


### Split an ipv6 network in subnets
```
$ go-ipcalc split 2001:41D0:FFC0::/63 64
2001:41d0:ffc0::/64
2001:41d0:ffc0:1::/64

$ go-ipcalc split 2001:41D0:FFC0::/62 64
2001:41d0:ffc0::/64
2001:41d0:ffc0:1::/64
2001:41d0:ffc0:2::/64
2001:41d0:ffc0:3::/64

$ go-ipcalc split 2001:41D0:FFC0::/61 64
2001:41d0:ffc0::/64
2001:41d0:ffc0:1::/64
2001:41d0:ffc0:2::/64
2001:41d0:ffc0:3::/64
2001:41d0:ffc0:4::/64
2001:41d0:ffc0:5::/64
2001:41d0:ffc0:6::/64
2001:41d0:ffc0:7::/64
```


