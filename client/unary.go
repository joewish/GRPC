package main
import (
  "context"
  "time"
  pb "github.com/GRPC/proto"
)
func callSayHello(client pb.GreetServiceClient){
  ctx,cancel:= context.withTimeout(context.Background(),time.Second)
  defer cancel()

  res,err:= client.SayHello(ctx,&pb.NoParam{})
  if err!=nil{
    log.Fatal(err)
  }
  log.Pritnf("%s",res.Message)
}