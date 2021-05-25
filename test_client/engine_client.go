package test_client

import (
	"log"
	"suvvm.work/ToadOCRRpcClient/rpc"
	"suvvm.work/ToadOCRRpcClient/utils"
)

func RunTestEngineClient(appID, appSecret, discoverUrl string) error {
	client := rpc.NewEngineClient(appID, appSecret, discoverUrl)
	if err := client.InitEngineClient(); err != nil {
		log.Printf("init engine client fail!")
		return err
	}
	filename := "resources/images/test_image_2.png"
	imageBytes, err := utils.GetImageGrayBytes(filename)
	if err != nil {
		log.Printf("GetImageGrayBytes fail")
		return err
	}
	preVal, err := client.Predict("snn", imageBytes)
	if err != nil {
		log.Printf("call rpc fail")
		return err
	}
	log.Printf("success label:%v", preVal)
	return nil
}
