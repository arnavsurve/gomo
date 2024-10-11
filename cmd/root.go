package cmd

import (
	"log"
	"os"

	"github.com/arnavsurve/gomo/pkg/models"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gomo",
	Short: "A minimal pomodoro utility for the CLI",
	Long: `gomo is a minimal pomodoro utility for the CLI.
	Written in Go with the Bubbletea TUI framework and the only CLI commander that matters, Cobra.
	Created by Arnav Surve (arnav@surve.dev, linkedin.com/in/arnavsurve, github.com/arnavsurve)`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			switch args[0] {
			case "start":
				// Unmarshal config yaml to be passed into models
				config := models.Config{}
				c := *config.GetConf()

				// Start timer with focus duration
				p := tea.NewProgram(models.NewStartModel(c.Focus * 60))
				if _, err := p.Run(); err != nil {
					log.Fatal(err)
				}
			case "config":
				p := tea.NewProgram(models.NewConfigModel())
				if _, err := p.Run(); err != nil {
					log.Fatal(err)
				}
			}
		} else if len(args) == 0 {
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/gomo/config.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
