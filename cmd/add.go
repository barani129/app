/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/barani129/app/todo"
)

// addCmd represents the add command
var priority int

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add will create a new todo item to the list`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.GetString("datafile"))
		items, _ := todo.ReadItems(viper.GetString("datafile"))
		for _, x := range args {
			item := todo.Item{Text: x}
			item.Label()
			item.SetPriority(priority)
			items = append(items, item)
		}
		if err := todo.SaveItems(viper.GetString("datafile"), items); err != nil {
			fmt.Errorf("%v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")
}
