package main

import (
    "context"
    "fmt"
    protos "magma/lte/cloud/go/protos"
    models "magma/lte/cloud/go/protos/models"
    "google.golang.org/grpc"
    //"github.com/go-openapi/swag"
    "log"
)

func main() {
 fmt.Println("Hello client ...")

 opts := grpc.WithInsecure()
 cc, err := grpc.Dial("localhost:50051", opts)
 if err != nil {
  log.Fatal(err)
 }
 defer cc.Close()

 client := protos.NewPMNSubscriberConfigServicerClient(cc)
 
 smsSubscribed := true
 sharedSmsSubsDataId := "12345678"
 request := SmsSubscriptionDataConverter(smsSubscribed, sharedSmsSubsDataId)
 client.PMNSubscriberConfig(context.Background(), request)
}


func SmsSubscriptionDataConverter(smsSubscribed bool, sharedSmsSubsDataId string) *models.SmsSubscriptionData {
	sd := &models.SmsSubscriptionData{
		SmsSubscribed:        smsSubscribed,
		SharedSmsSubsDataId: sharedSmsSubsDataId,
	}
  return &protos.PMNSubscriberData{
      Sms_data: sd,
  }
}
