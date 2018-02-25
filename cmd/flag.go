package cmd

import "github.com/spf13/cobra"

var (
	configFile string
	srcHost    string
	srcPort    string
	dstHost    string
	dstPort    string
)

func init() {
	initGlobalFlags(rootCmd)
}

func initGlobalFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", `Set the path of config file`)

	cmd.PersistentFlags().StringVar(&srcHost, "src-host", "", `Set source address`)
	cmd.PersistentFlags().StringVar(&srcPort, "src-port", "", `Set source port`)
	cmd.PersistentFlags().StringVar(&dstHost, "dst-host", "", `Set destination address`)
	cmd.PersistentFlags().StringVar(&dstPort, "dst-port", "", `Set destination port`)

}
