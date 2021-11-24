package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	name1 string
)

func main() {
	rootCmd := &cobra.Command{
		Use: "app",
		Run: func(c *cobra.Command, args []string) {
			fmt.Println("name1:", name1)
			if name2, err := c.PersistentFlags().GetString("name2"); err == nil {
				fmt.Println("name2:", name2)
			}
		},
	}

	rootCmd.PersistentFlags().StringVar(&name1, "name1", "name1", "your name1")
	rootCmd.PersistentFlags().StringP("name2", "n", "name2", "your name2")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
