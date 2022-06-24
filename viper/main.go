package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("This is a Viper Example...")

	pflag.Int("flagname", 9999, "help message for flagname")
	if viper.IsSet("flagname") {
		i := viper.GetInt("flagname")
		fmt.Printf("flagname=%d\n", i)
	} else {
		fmt.Println("Didnt set!!")
	}

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	i := viper.GetInt("flagname") // retrieve values from viper instead of pflag
	fmt.Printf("flagname=%d\n", i)

	fmt.Println("Finishing.")
}
