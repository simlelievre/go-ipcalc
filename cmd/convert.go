// Copyright © 2015 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert <ipv6 network> <ipv4>",
	Short: "convert ipv6 and ipv4 to 6rd cisco ip",
	Long:  `convert ipv6 and ipv4 to 6rd cisco ip`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Help()
			return
		}

		pNet, err := ConvertIPv4InIPv6Network(args[0], args[1])
		if err != nil {
			fmt.Println("Error on convertion :", err)
		}

		//  fmt.Printf("%#v\n", pNet)
		fmt.Println(pNet)
	},
}

func init() {
	RootCmd.AddCommand(convertCmd)
}

// ConvertIPv4InIPv6Network convert ipv6 network followed by ipv4
func ConvertIPv4InIPv6Network(network, ip string) (*net.IPNet, error) {
	// parse original network
	_, pNet1, err := net.ParseCIDR(network)
	if err != nil {
		return nil, err
	}

	// create an empty network
	_, pNet2, err := net.ParseCIDR("0::/63")
	if err != nil {
		return nil, err
	}

	// set 2 last bytes in new network
	ip2 := net.ParseIP(ip)
	pNet2.IP[0] = ip2[14]
	pNet2.IP[1] = ip2[15]

	// we shift network to follow the first ip
	shift, _ := pNet1.Mask.Size()
	pNet2.IP = SliceShiftRight(pNet2.IP, uint(shift))

	// return the OR of joined network
	return IPNetOr(pNet1, pNet2), nil
}

// IPNetOr make OR bitwise operation on net.IPNet
func IPNetOr(n1, n2 *net.IPNet) *net.IPNet {
	oip := SliceOr(n1.IP, n2.IP)
	omask := SliceOr(n1.Mask, n2.Mask)
	return &net.IPNet{IP: oip, Mask: omask}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// SliceOr return an OR on two slice of bytes
func SliceOr(s1, s2 []byte) []byte {
	mini := min(len(s1), len(s2))
	s := make([]byte, mini)
	for i := 0; i < mini; i++ {
		s[i] = s1[i] | s2[i]
	}
	return s
}

// SliceShiftRight return a slive of bytes shifted by "bits" bits
func SliceShiftRight(s1 []byte, bits uint) []byte {
	l := len(s1)
	s := make([]byte, l)

	r := bits / 8
	rs := bits - r*8

	for i := 0; i < l; i++ {
		// relative index
		ri := i - int(r)
		bl := byte(0)
		if ri >= 1 {
			bl = s1[ri-1] << (8 - rs)
		}

		bu := byte(0)
		if ri >= 0 {
			bu = s1[ri] >> rs
		}
		s[i] = bl | bu
	}

	return s
}
