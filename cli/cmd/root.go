// Copyright © 2018 eine <@github>
//
// Unless required by applicable law or agreed to in writing, this
// software is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "go-icemulti",
	Short: "Go(lang) tools to exploit warm/cold boot in iCE40 FPGAs",
	Long: `
go-icemulti is the main command.

DTD is a fast and flexible CLI library for Go to merge/pack, reorder,
and explore binary files contaning multiple bitstreams for iCE40 FPGAs.

Documentation is available at eine.github.io/icemulti`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[icemulti] Go(lang) tools to exploit warm/cold boot in iCE40 FPGAs")
		fmt.Println("This is the default cmd (no subcomand)")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var cfgFile string
var Verbose int
var Output string

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/icemulti.yaml)")
	rootCmd.PersistentFlags().IntVarP(&Verbose, "verbose", "v", 0, "0: disabled, 1: minimum, 2: full")
	rootCmd.PersistentFlags().StringVarP(&Output, "output", "o", "", "write output to file")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
