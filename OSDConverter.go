func ConvertPMNSubscriberDataToProto(data *PMNSubscriberData) (*lte_protos_models.OperatorSpecificData, error) {
	opcData := &lte_protos_models.OperatorSpecificData{}

	// Convert SubscriberId
	if data.Osd != nil && data.Osd.SubscriberId != nil {
		subscriberId := &lte_protos_models.SubscriberId{
			DataType: data.Osd.SubscriberId.DataType,
			SIdValue: data.Osd.SubscriberId.SIdValue,
		}
		opcData.SubscriberId = subscriberId
	}

	// Convert SubscriberInfo
	if data.Osd != nil && data.Osd.SubscriberInfo != nil {
		subscriberInfo := &lte_protos_models.SubscriberInfo{
			DataType: data.Osd.SubscriberInfo.DataType,
		}
		if data.Osd.SubscriberInfo.SiValue != nil {
			siValue := &lte_protos_models.SIValue{
				PricingPlanType:         data.Osd.SubscriberInfo.SiValue.PricingPlanType,
				PlanName:                data.Osd.SubscriberInfo.SiValue.PlanName,
				HomeLocation:            data.Osd.SubscriberInfo.SiValue.HomeLocation,
				UserNotification:        data.Osd.SubscriberInfo.SiValue.UserNotification,
				SubscriberEmailId:       data.Osd.SubscriberInfo.SiValue.SubscriberEmailId,
				SubscriberValidity:      data.Osd.SubscriberInfo.SiValue.SubscriberValidity,
				SubscriberCategory:      data.Osd.SubscriberInfo.SiValue.SubscriberCategory,
				TopupServicesSubscribed: data.Osd.SubscriberInfo.SiValue.TopupServicesSubscribed,
			}
			subscriberInfo.SiValue = siValue
		}
		opcData.SubscriberInfo = subscriberInfo
	}

	// Convert ServiceSubscription
	if data.Osd != nil && data.Osd.ServiceSubscription != nil {
		serviceSubscription := &lte_protos_models.ServiceSubscription{
			DataType: data.Osd.ServiceSubscription.DataType,
		}
		if data.Osd.ServiceSubscription.SsValue != nil {
			for _, ssValue := range data.Osd.ServiceSubscription.SsValue {
				service := &lte_protos_models.SSValue{
					ServiceName:          ssValue.ServiceName,
					ServiceId:            ssValue.ServiceId,
					BillingStartDate:     ssValue.BillingStartDate,
					BillingEndDate:       ssValue.BillingEndDate,
					QuotaStartDate:       ssValue.QuotaStartDate,
					QuotaEndDate:         ssValue.QuotaEndDate,
					MonitoringKey:        ssValue.MonitoringKey,
					RatingGroup:          ssValue.RatingGroup,
					TotalThreshold:       ssValue.TotalThreshold,
					RecurringQuotaReset:  ssValue.RecurringQuotaReset,
					CurrentRolloverCount: ssValue.CurrentRolloverCount,
				}
				serviceSubscription.SsValue = append(serviceSubscription.SsValue, service)
			}
		}
		opcData.ServiceSubscription = serviceSubscription
	}

	// Convert VolumeAccounting
	if data.Osd != nil && data.Osd.VolumeAccounting != nil {
		volumeAccounting := &lte_protos_models.VolumeAccounting{
			DataType: data.Osd.VolumeAccounting.DataType,
		}
		if data.Osd.VolumeAccounting.VaValue != nil {
			for _, vaValue := range data.Osd.VolumeAccounting.VaValue {
				volume := &lte_protos_models.VAValue{
					ServiceName:    vaValue.ServiceName,
					ServiceId:      vaValue.ServiceId,
					TotalUsedQuota: vaValue.TotalUsedQuota,
					UlUsedQuota:    vaValue.UlUsedQuota,
					DlUsedQuota:    vaValue.DlUsedQuota,
					MonitoringKey:  vaValue.MonitoringKey,
					GracePeriod:    vaValue.GracePeriod,
				}
				volumeAccounting.VaValue = append(volumeAccounting.VaValue, volume)
			}
		}
		opcData.VolumeAccounting = volumeAccounting
	}

	return opcData, nil
}
