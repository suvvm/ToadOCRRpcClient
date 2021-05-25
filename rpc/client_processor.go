package rpc

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"log"
	"strconv"
	pb_processor "suvvm.work/ToadOCRRpcClient/rpc/toad_ocr_preprocessor/idl"
	"suvvm.work/ToadOCRRpcClient/utils"
	"time"
)

type ProcessorClient struct {
	AppID string
	AppSecret string
	DiscoverUrl string
	toadOCRPreprocessorClient pb_processor.ToadOcrPreprocessorClient
}

func (processor *ProcessorClient) InitProcessorClient(appID, appSecret, discoverUrl string) error {
	flag.Parse()
	var url string
	url = discoverUrl
	if discoverUrl == ""{
		log.Printf("discoverUrl is empty using default at http://localhost:2379")
		url = *defaultDiscoverUrl
	}
	if appID == "" || appSecret == "" {
		return fmt.Errorf("appID or appSecret is empty")
	}
	r := NewResolver(url, *engineServ)
	resolver.Register(r)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(ctx, r.Scheme()+"://authority/"+url,
		grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name), grpc.WithBlock())
	if err != nil {
		return err
	}
	processor.toadOCRPreprocessorClient = pb_processor.NewToadOcrPreprocessorClient(*conn)
	return nil
}

func (processor *ProcessorClient) Process(netFlag, appID string, image []byte) ([]string, error) {
	req := &pb_processor.ProcessRequest{
		AppId: appID,
		NetFlag: netFlag,
		Image: image,
	}
	token, err :=  utils.GetBasicToken(processor.AppSecret, req.NetFlag + strconv.Itoa(len(req.Image)))
	if err != nil {
		return nil, err
	}
	req.BasicToken = token
	resp, err := processor.toadOCRPreprocessorClient.Process(context.Background(), req)
	if err != nil {
		return nil, err
	}
	if resp.Code != int32(*successCode) {
		err = fmt.Errorf("resp code not success code:%v message:%v", resp.Code, resp.Message)
		return nil, err
	}
	return resp.Labels, nil
}

