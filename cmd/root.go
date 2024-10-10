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
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			switch args[0] {
			case "start":
				p := tea.NewProgram(models.NewStartModel(1200)) // TODO placeholder, pass user configured value (in seconds) to the model params. 1200 = 20 min * 60 sec
				if _, err := p.Run(); err != nil {
					log.Fatal(err)
				}

			case "config":
				p := tea.NewProgram(models.NewConfigModel())
				if _, err := p.Run(); err != nil {
					log.Fatal(err)
				}
			}
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
