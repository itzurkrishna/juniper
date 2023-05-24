func OperatorSpecificDataConverter(subscriberID string, subscriberInfo *models.SubscriberInfo, serviceSubscription []*models.ServiceSubscription, volumeAccounting []*models.VolumeAccounting) *protos.OperatorSpecificData {
	subscriberIDProto := &protos.Subscriberid{
		DataType: "string",
		Value:    subscriberID,
	}

	subscriberInfoProto := &protos.Subscriberinfo{
		DataType: "string",
		Value1: &protos.Value1{
			PricingPlanType:        subscriberInfo.PricingPlanType,
			PlanName:               subscriberInfo.PlanName,
			QoSProfileName:         subscriberInfo.QoSProfileName,
			HomeLocation:           subscriberInfo.HomeLocation,
			UserNotification:       subscriberInfo.UserNotification,
			SubscriberEmailId:      subscriberInfo.SubscriberEmailId,
			SubscriberValidity:     subscriberInfo.SubscriberValidity,
			SubscriberCategory:     subscriberInfo.SubscriberCategory,
			TopupServicesSubscribed: subscriberInfo.TopupServicesSubscribed,
		},
	}

	serviceSubscriptionProto := make([]*protos.Servicesubscription, len(serviceSubscription))
	for i, subscription := range serviceSubscription {
		serviceSubscriptionProto[i] = &protos.Servicesubscription{
			DataType: "string",
			Value2: &protos.Value2{
				ServiceName:           subscription.ServiceName,
				ServiceId:             subscription.ServiceId,
				BillingStartDate:      subscription.BillingStartDate,
				BillingEndDate:        subscription.BillingEndDate,
				QuotaStartDate:        subscription.QuotaStartDate,
				QuotaEndDate:          subscription.QuotaEndDate,
				MonitoringKey:         subscription.MonitoringKey,
				RatingGroup:           subscription.RatingGroup,
					TotalThreshold:        subscription.TotalThreshold,
				RecurringQuotaReset:   subscription.RecurringQuotaReset,
				CurrentRolloverCount:  subscription.CurrentRolloverCount,
			},
		}
	}

	volumeAccountingProto := make([]*protos.Volumeaccounting, len(volumeAccounting))
	for i, accounting := range volumeAccounting {
		volumeAccountingProto[i] = &protos.Volumeaccounting{
			DataType: "string",
			Value3: &protos.Value3{
				ServiceName:     accounting.ServiceName,
				ServiceId:       accounting.ServiceId,
				TotalUsedQuota:  accounting.TotalUsedQuota,
				ULUsedQuota:     accounting.ULUsedQuota,
				DLUsedQuota:     accounting.DLUsedQuota,
				MonitoringKey:   accounting.MonitoringKey,
				GracePeriod:     accounting.GracePeriod,
			},
		}
	}

	return &protos.OperatorSpecificData{
		SubscriberId:        subscriberIDProto,
		SubscriberInfo:      subscriberInfoProto,
		ServiceSubscription: serviceSubscriptionProto,
		VolumeAccounting:    volumeAccountingProto,
	}
}
