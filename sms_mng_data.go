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
 
 supportedFeatures := "5G core"
 mtSmsSubscribed := true
 mtSmsBarringAll := true
 mtSmsBarringRoaming := true
 moSmsSubscribed := true
 moSmsBarringAll := true
 moSmsBarringRoaming := true
 request := SmsManagementSubscriptionDataConverter(supportedFeatures, mtSmsSubscribed, mtSmsBarringAll, mtSmsBarringRoaming, moSmsSubscribed, moSmsBarringAll, moSmsBarringRoaming)
 client.PMNSubscriberConfig(context.Background(), request)
}


func SmsManagementSubscriptionDataConverter(supportedFeatures string, mtSmsSubscribed, mtSmsBarringAll, mtSmsBarringRoaming, moSmsSubscribed, moSmsBarringAll, moSmsBarringRoaming bool) *protos.PMNSubscriberData {
	smd := &models.SmsManagementSubscriptionData{
		SupportedFeatures: supportedFeatures,
		MtSmsSubscribed:   mtSmsSubscribed,
		MtSmsBarringAll:   mtSmsBarringAll,
		MtSmsBarringRoaming: mtSmsBarringRoaming,
		MoSmsSubscribed:   moSmsSubscribed,
		MoSmsBarringAll:   moSmsBarringAll,
		MoSmsBarringRoaming: moSmsBarringRoaming,
	}
  	return &protos.PMNSubscriberData{
      		SmsMngData: smd,
  	}
}
