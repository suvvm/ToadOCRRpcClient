package rpc

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"log"
	pb_engine "suvvm.work/ToadOCRRpcClient/rpc/toad_ocr_engine/idl"
	"suvvm.work/ToadOCRRpcClient/utils"
	"time"
)

type EngineClient struct {
	AppID string
	AppSecret string
	DiscoverUrl string
	toadOCREngineClient pb_engine.ToadOcrClient
}

func (engine *EngineClient) InitEngineClient(appID, appSecret, discoverUrl string) error {
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
	engine.toadOCREngineClient = pb_engine.NewToadOcrClient(*conn)
	return nil
}

func (engine *EngineClient) Predict(netFlag string, image []byte) (string, error) {
	req := &pb_engine.PredictRequest{
		AppId: engine.AppID,
		NetFlag: netFlag,
		Image: image,
	}
	token, err :=  utils.GetBasicToken(engine.AppSecret, req.NetFlag + utils.PixelHashStr(req.Image))
	if err != nil {
		return "", err
	}
	req.BasicToken = token
	resp, err := engine.toadOCREngineClient.Predict(context.Background(), req)
	if err != nil {
		return "", err
	}
	if resp.Code != int32(*successCode) {
		err = fmt.Errorf("resp code not success code:%v message:%v", resp.Code, resp.Message)
		return "", err
	}
	return resp.Label, nil
}
