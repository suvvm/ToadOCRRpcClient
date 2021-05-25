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
	pb_processor "github.com/suvvm/ToadOCRRpcClient/rpc/toad_ocr_preprocessor/idl"
	"github.com/suvvm/ToadOCRRpcClient/utils"
	"time"
)

type ProcessorClient struct {
	AppID string
	AppSecret string
	DiscoverUrl string
	toadOCRPreprocessorClient pb_processor.ToadOcrPreprocessorClient
}

func NewProcessorClient(appID, appSecret, discoverUrl string) *ProcessorClient {
	return &ProcessorClient{
		AppID: appID,
		AppSecret: appSecret,
		DiscoverUrl: discoverUrl,
	}
}

func (processor *ProcessorClient) InitProcessorClient() error {
	flag.Parse()
	if processor.DiscoverUrl == ""{
		log.Printf("discoverUrl is empty using default at http://localhost:2379")
		processor.DiscoverUrl = *defaultDiscoverUrl
	}
	if processor.AppID == "" || processor.AppSecret == "" {
		return fmt.Errorf("appID or appSecret is empty")
	}
	r := NewResolver(processor.DiscoverUrl, *processorServ)
	resolver.Register(r)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(ctx, r.Scheme()+"://authority/"+processor.DiscoverUrl,
		grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name), grpc.WithBlock())
	if err != nil {
		return err
	}
	processor.toadOCRPreprocessorClient = pb_processor.NewToadOcrPreprocessorClient(*conn)
	return nil
}

func (processor *ProcessorClient) Process(netFlag string, image []byte) ([]string, error) {
	req := &pb_processor.ProcessRequest{
		AppId: processor.AppID,
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

