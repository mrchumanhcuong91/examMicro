package handler

import (
	"fmt"
	"context"
	// "authservice/db/models"
	db "authservice/db/databases"
	log "github.com/micro/micro/v3/service/logger"
	utils "authservice/helper"
	authservice "authservice/proto"
)

type Authservice struct{
	PstDb * db.MyPostGres
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Authservice) Call(ctx context.Context, req *authservice.Request, rsp *authservice.Response) error {
	log.Info("Received Authservice.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
func (e *Authservice) GetNewToken(ctx context.Context, req *authservice.Request, rsp *authservice.Response) error {
	log.Info("Received Authservice.GetNewToken request")
	// rsp.Msg = "Hello " + req.Username
	// //create key
	tk,pubkey,encryptSecret, _ := utils.NewToken(map[string]interface{}{
		"user": req.Username,"expires":req.Expiretime,
	})
	// //save to db
	e.PstDb.InsertKey(fmt.Sprintf("%x",pubkey),fmt.Sprintf("%x",encryptSecret))
	// //create token
	rsp.Token = tk
	rsp.Pubkey = fmt.Sprintf("%x",pubkey)
	fmt.Printf("GetNewToken token %v , pubkey %x",rsp.Token, rsp.Pubkey)
	//return token + pubkey to client
	return nil
}
// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Authservice) Stream(ctx context.Context, req *authservice.StreamingRequest, stream authservice.Authservice_StreamStream) error {
	log.Infof("Received Authservice.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&authservice.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Authservice) PingPong(ctx context.Context, stream authservice.Authservice_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&authservice.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
