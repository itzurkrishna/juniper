
func ConvertPMNSubscriberDataToProto() *models.OperatorSpecificData {

	subscriberId := &models.SubscriberId{
		DataType: "string",
		SIdValue: "A",
	}

	qosProfileNames := []*models.QosProfileName{
		{
			QosProfileName:     "",
			PcrfARPPrioLevel:   "",
			PcrfPreEmptionCap:  "",
			PcrfPreEmptionVuln: "",
			PcrfQoSClassName:   "",
			PcrfMaxReqBrUL:     "",
			PcrfMaxReqBrDL:     "",
			PcrfGuarBrUL:       "",
			PcrfGuarBrDL:       "",
			PcrfAPNAggMaxBrUL:  "",
			PcrfAPNAggMaxBrDL:  "",
		},
	}

	topupServicesSubscribed := []string{"", "", "", ""}

	siValue := &models.SIValue{
		PricingPlanType:         "",
		PlanName:                "199",
		QosProfileName:          qosProfileNames,
		HomeLocation:            "",
		UserNotification:        "",
		SubscriberEmailId:       "",
		SubscriberValidity:      "",
		SubscriberCategory:      "",
		TopupServicesSubscribed: topupServicesSubscribed,
	}

	subscriberInfo := &models.SubscriberInfo{
		DataType: "object",
		SiValue:  siValue,
	}

	ssValue := []*models.SsValue{
		{
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
		},
	}

	serviceSubscription := &models.ServiceSubscription{
		DataType: "object",
    		SsValue: ssValue,
	}
	vaValue := []*models.VaValue{
		{
			ServiceName:    "",
			ServiceId:      "",
			TotalUsedQuota: "",
			UlUsedQuota:    "",
			DlUsedQuota:    "",
			MonitoringKey:  "",
			GracePeriod:    "",
		},
	}

	volumeAccounting := &models.VolumeAccounting{
		DataType: "object",
    		VaValue: vaValue,
	}

	return &models.OperatorSpecificData{
		SubscriberId:         subscriberId,
		SubscriberInfo:       subscriberInfo,
		ServiceSubscription:  serviceSubscription,
		VolumeAccounting:     volumeAccounting,
	}
}
