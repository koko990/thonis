package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/thonis/pkg/log"
	"github.com/thonis/server"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start a http reverse proxy server",
	Long:  `start a http reverse proxy server`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Infoln("Start proxy server")

		doneCh := make(chan struct{})
		go func() {
			select {
			case <-doneCh:
				fmt.Println("stop")
			}
		}()

		sIp, err := checkSrcHost()
		if err != nil {
			log.Errorln(err)
			return
		}

		dIp, err := checkDstHost()
		if err != nil {
			log.Errorln(err)
			return
		}

		sPort, err := checkSrcPort()
		if err != nil {
			log.Errorln(err)
			return
		}

		dPort, err := checkDstPort()
		if err != nil {
			log.Errorln(err)
			return
		}

		newProxy := server.Proxy{
			SourceHost:      sIp,
			SourcePort:      sPort,
			DestinationHost: dIp,
			DestinationPort: dPort,
		}

		newProxy.Run(doneCh)

		// Do Stuff Here

	},
}
