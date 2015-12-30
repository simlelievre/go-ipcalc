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

	netadv "github.com/simlelievre/go-netadv"
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

		pNet, err := netadv.ConvertIPv4InIPv6Network(args[0], args[1])
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
