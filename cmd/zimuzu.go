// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software // distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zhaoyao/showtime/zimuzu"
	"os"
)

func newCtx() (*zimuzu.Ctx, error) {
	account := viper.GetString("zimuzu.account")
	password := viper.GetString("zimuzu.password")
	ctx := zimuzu.New(account, password)
	return ctx, ctx.Login()
}

// zimuzuCmd represents the zimuzu command
var zimuzuCmd = &cobra.Command{
	Use:   "zimuzu",
	Short: "A brief description of your command",
	//Run: func(cmd *cobra.Command, args []string) {},
}

var zimuzuSearch = &cobra.Command{
	Use: "search",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, err := newCtx()
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		ret, err := ctx.Search(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		for _, res := range ret {
			fmt.Printf("%s: %s\n", res.Title, res.ItemID)

		}
	},
}

var (
	season int
	ep     int
)
var zimuzuList = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, err := newCtx()
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		listResource(ctx, args[0], season, ep)
	},
}

func init() {
	RootCmd.AddCommand(zimuzuCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// zimuzuCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// zimuzuCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	zimuzuCmd.PersistentFlags().String("account", "", "account to login zimuzu")
	zimuzuCmd.PersistentFlags().String("password", "", "password of account")

	viper.BindPFlag("zimuzu.account", RootCmd.PersistentFlags().Lookup("account"))
	viper.BindPFlag("zimuzu.password", RootCmd.PersistentFlags().Lookup("password"))

	zimuzuCmd.AddCommand(zimuzuSearch, zimuzuList)

	zimuzuList.Flags().IntVar(&season, "season", -1, "season")
	zimuzuList.Flags().IntVar(&ep, "ep", -1, "ep")
}

func listResource(ctx *zimuzu.Ctx, id string, season, ep int) {
	files, err := ctx.GetResource(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range files {
		if season > 0 && f.Season != season {
			continue
		}
		if ep > 0 && f.Episode != ep {
			continue
		}
		fmt.Printf("S%d EP%d: %s\n", f.Season, f.Episode, f.Format)
		for _, l := range f.Links {
			fmt.Printf("\t%s => %s\n", l.Type, l.URL)
		}
	}

}
