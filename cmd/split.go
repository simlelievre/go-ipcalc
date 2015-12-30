// Copyright © 2015 Simon Lelievre
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
	"strconv"

	"github.com/spf13/cobra"
)

// splitCmd represents the split command
var splitCmd = &cobra.Command{
	Use:   "split <network> <newsize>",
	Short: "split network in network smaller",
	Long:  `split network in network smaller`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Help()
			return
		}

		_, pNet, err := net.ParseCIDR(args[0])
		if err != nil {
			fmt.Println("Error on convertion :", err)
			return
		}

		// parse size parameter and verify if it integer
		newsize, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error on convertion :", err)
			return
		}
		if newsize < 1 {
			fmt.Println("newsize too small")
			return
		}
		// check is size is greater than pNet1.Size()
		oldsize, _ := pNet.Mask.Size()
		if newsize <= oldsize {
			fmt.Println("newsize had to be greater to", oldsize)
			return
		}

		aNet, err := SplitNetworks(pNet, uint(newsize))
		if err != nil {
			fmt.Println("Error on convertion :", err)
			return
		}
		for _, n := range aNet {
			fmt.Println(n)
		}
	},
}

func init() {
	RootCmd.AddCommand(splitCmd)
}

// SplitNetworks split network in a
func SplitNetworks(network *net.IPNet, newsize uint) ([]*net.IPNet, error) {
	var ret []*net.IPNet
	oldsize, _ := network.Mask.Size()

	// use a stack to avoid a recurvive function
	ret = append(ret, network)
	for i := oldsize; i < int(newsize); i++ {
		var tmp []*net.IPNet

		// parse network array to split all of them in two
		for _, n := range ret {
			one, two := SplitNetworkInTwo(n)
			tmp = append(tmp, one)
			tmp = append(tmp, two)
		}
		ret = tmp
	}

	return ret, nil
}

// SplitNetworkInTwo split a network in two. return two sub network
func SplitNetworkInTwo(network *net.IPNet) (*net.IPNet, *net.IPNet) {
	size, _ := network.Mask.Size()
	newMask := net.CIDRMask(size+1, 8*net.IPv6len)

	byte1 := net.IP{128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	ip2 := SliceShiftRight(byte1, uint(size))
	ip2 = SliceOr(network.IP, ip2)

	return &net.IPNet{IP: network.IP, Mask: newMask}, &net.IPNet{IP: ip2, Mask: newMask}
}
