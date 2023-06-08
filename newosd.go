func ConvertPMNSubscriberDataToProto() *lte_protos_models.OperatorSpecificData{
  
  subscriberId := &lte_protos_models.SubscriberId{
		DataType: "string",
		SIdValue: "A",
	}
  
  qosProfileName := &lte_protos_models.QosProfileName{
    QosProfileName:     ""
    PcrfARPPrioLevel:   ""
    PcrfPreEmptionCap:  ""
    PcrfPreEmptionVuln: ""
    PcrfQoSClassName:   ""
    PcrfMaxReqBrUL:     ""
    PcrfMaxReqBrDL:     ""
    PcrfGuarBrUL:       ""
    PcrfGuarBrDL:       ""
    PcrfAPNAggMaxBrUL:  ""
    PcrfAPNAggMaxBrDL:  ""
  }
  
  topupServicesSubscribed := []string{"", "", "", ""},
  
	siValue := &lte_protos_models.SIValue{
		PricingPlanType:         "",
		PlanName:                "199",
    QosProfileName:          qosProfileName
		HomeLocation:            "",
		UserNotification:        "",
		SubscriberEmailId:       "",
		SubscriberValidity:      "",
		SubscriberCategory:      "",
		TopupServicesSubscribed: topupServicesSubscribed,
	}
  
  subscriberInfo := &lte_protos_models.SubscriberInfo{
		DataType: "object",
    SiValue = siValue,
	}
  
  serviceSubscription := &lte_protos_models.ServiceSubscription{
		DataType: "object",
    SsValue: ssValue1,
    
	}
	ssValue1 := &lte_protos_models.SSValue{
		ServiceName:          "BasePlan",
		ServiceId:            "12345",
		BillingStartDate:     "",
		BillingEndDate:       "",
		QuotaStartDate:       "",
		QuotaEndDate:         "",
		MonitoringKey:        "",
		RatingGroup:          "",
		TotalThreshold:       "",
		RecurringQuotaReset:  "",
		CurrentRolloverCount: "",
	}
  
  volumeAccounting := &lte_protos_models.VolumeAccounting{
		DataType: "object",
    VaValue: vaValue1,
	}
	vaValue1 := &lte_protos_models.VAValue{
		ServiceName:    "",
		ServiceId:      "",
		TotalUsedQuota: "",
		UlUsedQuota:    "",
		DlUsedQuota:    "",
		MonitoringKey:  "",
		GracePeriod:    "",
	}
  
  
  return &models.Osd{
    SubscriberId: subscriberId,
    SubscriberInfo: subscriberInfo,
    ServiceSubscription: serviceSubscription,
    VolumeAccounting: volumeAccounting,
  }
  
}
