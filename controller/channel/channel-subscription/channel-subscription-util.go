package channel_subscription

import (
	"fmt"
	"timtubeWebAdmin/domain"
	user3 "timtubeWebAdmin/io/channel/channel"
	dcs "timtubeWebAdmin/io/channel/channel-subscription"
	user2 "timtubeWebAdmin/io/user/user"
)

func GetRealChannelSubscription(id string) ([]domain.ChannelSubscription, error) {
	var channelSubscriptionList []domain.ChannelSubscription
	channels, err := dcs.ReadChannelSubscriptionByChannel(id)
	if err != nil {
		fmt.Println(err, "error reading channels")
		return []domain.ChannelSubscription{}, err
	}
	for _, channelSubscription := range channels {
		cso, err := GetCompleteChannelSubscription(channelSubscription)
		if err != nil {
			fmt.Println(err, " error reading real channel subscription")
		}
		channelSubscriptionList = append(channelSubscriptionList, cso)
	}
	return channelSubscriptionList, nil
}
func GetCompleteChannelSubscription(channelSubscription domain.ChannelSubscription) (domain.ChannelSubscription, error) {
	user, err := user2.ReadUser(channelSubscription.UserId)
	if err != nil {
		fmt.Println(err, " error reading complete user channel")
		return domain.ChannelSubscription{}, err
	}
	channel, err := user3.ReadChannel(channelSubscription.ChannelId)
	if err != nil {
		fmt.Println(err, " error reading complete channel")
		return domain.ChannelSubscription{}, err
	}
	channelSubscriptionObject := domain.ChannelSubscription{channelSubscription.Id, channel.Name, user.Name, channelSubscription.Date}
	return channelSubscriptionObject, nil
}
