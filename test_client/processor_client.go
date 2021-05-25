package test_client

import (
	"io/ioutil"
	"log"
	"suvvm.work/ToadOCRRpcClient/rpc"
)

func RunTestProcessorClient (appID, appSecret, discoverUrl string) error {
	client := rpc.NewProcessorClient(appID, appSecret, discoverUrl)
	if err := client.InitProcessorClient(); err != nil {
		log.Printf("init processor client fail!")
		return err
	}
	filename := "resources/images/test_image_1.jpg"
	imageBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	log.Printf("image size:%v", len(imageBytes))
	preVal, err := client.Process("snn", imageBytes)
	if err != nil {
		log.Printf("call rpc fail")
		return err
	}
	log.Printf("success labels:%v", preVal)
	return nil
}
