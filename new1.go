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
  dataType := "exampleType"
	sIdValue := "exampleIdValue"
  dt1 := "xyz"
  qos := []*QosProfileName{
    {
      QosProfileName    : "ExQosProfileName"
      PcrfARPPrioLevel  : "ExPcrfARPPrioLevel"
      PcrfPreEmptionCap : "ExPcrfPreEmptionCap"
      PcrfPreEmptionVuln: "ExPcrfPreEmptionVul"
      PcrfQoSClassName  : "ExPcrfQoSClassName"
      PcrfMaxReqBrUL    : "ExPcrfMaxReqBrUL"
      PcrfMaxReqBrDL    : "ExPcrfMaxReqBrDL"
      PcrfGuarBrUL      : "ExPcrfGuarBrUL"
      PcrfGuarBrDL      : "ExPcrfGuarBrDL"
      PcrfAPNAggMaxBrUL : "ExPcrfAPNAggMaxBrUL"
      PcrfAPNAggMaxBrDL : "ExPcrfAPNAggMaxBrDL"
    }
  }
  
  siValue := &SIValue{
		PricingPlanType:         "examplePlanType",
		PlanName:                "examplePlanName",
		QosProfileName:          qos,
		HomeLocation:            "exampleHomeLocation",
		UserNotification:        "exampleNotification",
		SubscriberEmailId:       "exampleEmail",
		SubscriberValidity:      "exampleValidity",
		SubscriberCategory:      "exampleCategory",
		TopupServicesSubscribed: []string{"exampleService1", "exampleService2"},
	}
  
  ssValue := []*SSValue{
		{
			ServiceName:          "exampleService1",
			ServiceId:            "exampleServiceId1",
			BillingStartDate:     "exampleBillingStartDate1",
			BillingEndDate:       "exampleBillingEndDate1",
			QuotaStartDate:       "exampleQuotaStartDate1",
			QuotaEndDate:         "exampleQuotaEndDate1",
			MonitoringKey:        "exampleMonitoringKey1",
			RatingGroup:          "exampleRatingGroup1",
			TotalThreshold:       "exampleTotalThreshold1",
			RecurringQuotaReset:  "exampleRecurringQuotaReset1",
			CurrentRolloverCount: "exampleRolloverCount1",
		},
		{
			ServiceName:          "exampleService2",
			ServiceId:            "exampleServiceId2",
			BillingStartDate:     "exampleBillingStartDate2",
			BillingEndDate:       "exampleBillingEndDate2",
			QuotaStartDate:       "exampleQuotaStartDate2",
			QuotaEndDate:         "exampleQuotaEndDate2",
			MonitoringKey:        "exampleMonitoringKey2",
			RatingGroup:          "exampleRatingGroup2",
			TotalThreshold:       "exampleTotalThreshold2",
			RecurringQuotaReset:  "exampleRecurringQuotaReset2",
			CurrentRolloverCount: "exampleRolloverCount2",
		},
	}
  
  	vaValue := []*VAValue{
		{
			ServiceName:    "exampleService1",
			ServiceId:      "exampleServiceId1",
			TotalUsedQuota: "exampleTotalUsedQuota1",
			UlUsedQuota:    "exampleUlUsedQuota1",
			DlUsedQuota:    "exampleDlUsedQuota1",
			MonthlyQuota:   "exampleMonthlyQuota1",
		},
		{
			ServiceName:    "exampleService2",
			ServiceId:      "exampleServiceId2",
			TotalUsedQuota: "exampleTotalUsedQuota2",
			UlUsedQuota:    "exampleUlUsedQuota2",
			DlUsedQuota:    "exampleDlUsedQuota2",
			MonthlyQuota:   "exampleMonthlyQuota2",
		},
	}

 client := protos.NewPMNSubscriberConfigServicerClient(cc)

  
  
 
 request := OperatorSpecificDataConverter()
 client.PMNSubscriberConfig(context.Background(), request)
}

func CreateOperatorSpecificData(){
  
  
  subscriberId := &models.SubscriberId{
		DataType: dataType,
		SIdValue: sIdValue,
	}
  
  subscriberInfo := &models.SubscriberInfo{
		DataType: dt1,
		SiValue:  siValue,
	}
  
  serviceSubscription := &ServiceSubscription{
		DataType: "SSValue",
		SsValue:  ssValue,
	}

	volumeAccounting := &VolumeAccounting{
		DataType: "VAValue",
		VaValue:  vaValue,
	}
  
 osd := &models.OperatorSpecificData{
		SubscriberId:        subscriberId,
		SubscriberInfo:      subscriberInfo,
		ServiceSubscription: serviceSubscription,
		VolumeAccounting:    volumeAccounting,
}
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
