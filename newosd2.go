func ConvertPMNSubscriberDataToProto() *lte_protos_models.OperatorSpecificData {
	subscriberId := &lte_protos_models.SubscriberId{
		DataType: "string",
		SIdValue: "A",
	}

	var qosProfileName []*lte_protos_models.QosProfileName
	for _, qos := range SIValue.QosProfileName {
		qosProfile := &lte_protos_models.QosProfileName{
			QosProfileName:    "",
			PcrfARPPrioLevel:  "",
			PcrfPreEmptionCap: "",
			PcrfPreEmptionVuln: "",
			PcrfQoSClassName:  "",
			PcrfMaxReqBrUL:    "",
			PcrfMaxReqBrDL:    "",
			PcrfGuarBrUL:      "",
			PcrfGuarBrDL:      "",
			PcrfAPNAggMaxBrUL: "",
			PcrfAPNAggMaxBrDL: "",
		}
		qosProfileName = append(qosProfileName, qosProfile)
	}

	topupServicesSubscribed := []string{"", "", "", ""}

	siValue := &lte_protos_models.SIValue{
		PricingPlanType:         "",
		PlanName:                "199",
		QosProfileName:          qosProfileName,
		HomeLocation:            "",
		UserNotification:        "",
		SubscriberEmailId:       "",
		SubscriberValidity:      "",
		SubscriberCategory:      "",
		TopupServicesSubscribed: topupServicesSubscribed,
	}

	subscriberInfo := &lte_protos_models.SubscriberInfo{
		DataType: "object",
		SiValue:  siValue,
	}

	var ssValue []*lte_protos_models.SSValue
	for _, ss := range ServiceSubscription.SsValue {
		service := &lte_protos_models.SSValue{
			ServiceName:          ss.ServiceName,
			ServiceId:            ss.ServiceId,
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
		ssValue = append(ssValue, service)
	}

	serviceSubscription := &lte_protos_models.ServiceSubscription{
		DataType: "object",
		SsValue:  ssValue,
	}

	var vaValue []*lte_protos_models.VAValue
	for _, va := range VolumeAccounting.VaValue {
		volume := &lte_protos_models.VAValue{
			ServiceName:    va.ServiceName,
			ServiceId:      va.ServiceId,
			TotalUsedQuota: "",
			UlUsedQuota:    "",
			DlUsedQuota:    "",
			MonitoringKey:  "",
			GracePeriod:    "",
		}
		vaValue = append(vaValue, volume)
	}

	volumeAccounting := &lte_protos_models.VolumeAccounting{
		DataType: "object",
		VaValue:  vaValue,
	}

	return &models.Osd{
		SubscriberId:        subscriberId,
		SubscriberInfo:      subscriberInfo,
		ServiceSubscription: serviceSubscription,
		VolumeAccounting:    volumeAccounting,
	}
}
