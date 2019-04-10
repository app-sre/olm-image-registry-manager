package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile    string
	gitRepoURL string
	gitBranch  string
	gitDir     string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "olm-image-registry-manager",
	Short: "Tooling to help with OLM operator CD",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	rootCmd.PersistentFlags().StringVar(&gitRepoURL, "git-repo-url", "", "Git repo url")
	viper.BindPFlag("git-repo-url", rootCmd.PersistentFlags().Lookup("git-repo-url"))

	rootCmd.PersistentFlags().StringVar(&gitRepoURL, "git-branch", "", "Git branch")
	viper.BindPFlag("git-branch", rootCmd.PersistentFlags().Lookup("git-branch"))

	rootCmd.PersistentFlags().StringVar(&gitDir, "git-dir", "", "Git dir")
	viper.BindPFlag("git-dir", rootCmd.PersistentFlags().Lookup("git-dir"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".olm-image-registry-manager" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".olm-image-registry-manager")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
