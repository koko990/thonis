package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"net"
	"errors"
	"strconv"
)

var rootCmd = &cobra.Command{
	Use:   "thonis",
	Short: "thonis is a http reverse proxy",
	Long:  `thonis is a http reverse proxy`,
	Run: func(cmd *cobra.Command, args []string) {

		cmd.Help() // Do Stuff Here

	},
}

func Execute() {

	rootCmd.AddCommand(startCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}

//Check source host
func checkSrcHost() (net.IP, error) {
	ip := net.ParseIP(srcHost)

	if ip == nil {
		return nil, errors.New("Source host is wrong ")
	}

	return ip, nil
}

//Check source host
func checkDstHost() (net.IP, error) {
	ip := net.ParseIP(dstHost)

	if ip == nil {
		return nil, errors.New("Destination host is wrong ")
	}

	return ip, nil
}
func checkSrcPort() (int, error) {
	port, err := strconv.Atoi(srcPort)
	if err != nil {
		return 0, err
	}

	if port > 65535 {
		return 0, errors.New("Source port is wrong ")
	}

	return port, nil

}

func checkDstPort() (int, error) {
	port, err := strconv.Atoi(dstPort)
	if err != nil {
		return 0, err
	}

	if port > 65535 {
		return 0, errors.New("Destination port is wrong ")
	}

	return port, nil
}
