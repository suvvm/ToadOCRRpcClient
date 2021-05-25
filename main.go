package main

import (
	"log"
	"os"
	"suvvm.work/ToadOCRRpcClient/common"
	"suvvm.work/ToadOCRRpcClient/test_client"
)

func main() {
	if len(os.Args) < 2 {
		log.Printf("Please provide command parameters\n Running with " +
			"`help` to show currently supported commands")
		return
	}
	cmd := os.Args[1]
	if _, ok := common.CMDMap[cmd]; !ok {
		log.Printf("Unknow command!\n")
	} else if cmd == common.CmdHelp {
		log.Printf("\nToad OCR Rpc Client Help:\n" +
			"help: use command `%s` to show supported command\n" +
			"run: use command `%s {{app_id}} {{app_secret}} {{(optional)discover_url}}`" +
			" to run test rpc client to sent one snn predict ",common.CmdHelp, common.CmdRun)
	} else if cmd == common.CmdRun {
		if len(os.Args) < 4 {
			log.Printf("missing required parameters, please provide app_id and app_secret!")
			return
		}
		appID := os.Args[2]
		appSecet := os.Args[3]
		var discoverUrl string
		if len(os.Args) >= 5 {
			discoverUrl = os.Args[4]
		}
		if err := test_client.RunTestEngineClient(appID, appSecet, discoverUrl); err != nil {
			log.Printf("rnn engine error:%v", err)
		}
		if err := test_client.RunTestProcessorClient(appID, appSecet, discoverUrl); err != nil {
			log.Printf("rnn processor error:%v", err)
		}
	} else {
		log.Printf("Unknow command!\n")
	}
}
