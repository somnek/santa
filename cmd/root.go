package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	session string

	rootCmd = &cobra.Command{
		Use:   "santa [command]",
		Short: "\n🎄 Santa is a CLI to download inputs from AOC.\n",
	}

	sessionCmd = &cobra.Command{
		Use:   "session [token]",
		Short: "Session value from inspect element -> application -> cookies -> session",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			viper.Set("aoc_session", args[0])
			if err := viper.WriteConfig(); err != nil {
				fmt.Printf("Can't write config: %s \n", err)
				os.Exit(1)
			}
			fmt.Println("\nSession set complete.\nRun `santa day [day]` to download input.")
		},
	}

	downloadCmd = &cobra.Command{
		Use:   "day [day]",
		Short: "Download input for a given day",
		Args:  cobra.ExactArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			if viper.GetString("aoc_session") == "" {
				fmt.Println("\nNo session token set. Run `santa -s [session]` to set it.")
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			day := args[0]
			if _, err := strconv.Atoi(day); err != nil {
				fmt.Println("\nDay must be a number.")
				os.Exit(1)
			}
			download(cmd, day)
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Santa",
		Run: func(cmd *cobra.Command, args []string) {
			println("Santa v0.1.0")
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.santa.yaml)")
	rootCmd.AddCommand(sessionCmd)
	rootCmd.AddCommand(versionCmd, downloadCmd)

}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)

	} else {
		// find home
		configHome, err := homedir.Dir()
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		configName := ".santa" // we're using .santa.yaml, but viper doesn't require ext
		configType := "yaml"

		cobra.CheckErr(err)
		viper.AddConfigPath(configHome)
		viper.SetConfigType(configType)
		viper.SetConfigName(configName)

		// create config if it doesn't exist
		configPath := filepath.Join(configHome, configName+"."+configType)
		if _, err := os.Stat(configPath); err == nil {
			viper.ReadInConfig()

		} else if errors.Is(err, os.ErrNotExist) {
			viper.SetDefault("aoc_session", "")
			if err := viper.SafeWriteConfig(); err != nil {
				fmt.Printf("Can't write config: %s \n", err)
				os.Exit(1)
			}
		}
	}
}
