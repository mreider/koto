package services

import (
	"context"

	"github.com/mreider/koto/backend/common"
	"github.com/mreider/koto/backend/userhub/rpc"
)

type notificationService struct {
	*BaseService
}

func NewNotification(base *BaseService) rpc.NotificationService {
	return &notificationService{
		BaseService: base,
	}
}

func (s *notificationService) Count(ctx context.Context, _ *rpc.Empty) (*rpc.NotificationCountResponse, error) {
	me := s.getMe(ctx)
	total, unread, err := s.repos.Notification.Counts(me.ID)
	if err != nil {
		return nil, err
	}
	return &rpc.NotificationCountResponse{
		Total:  int32(total),
		Unread: int32(unread),
	}, nil
}

func (s *notificationService) Notifications(ctx context.Context, _ *rpc.Empty) (*rpc.NotificationNotificationsResponse, error) {
	me := s.getMe(ctx)
	notifications, err := s.repos.Notification.Notifications(me.ID)
	if err != nil {
		return nil, err
	}
	rpcNotifications := make([]*rpc.Notification, len(notifications))
	for i, notification := range notifications {
		rpcNotifications[i] = &rpc.Notification{
			Id:        notification.ID,
			Text:      notification.Text,
			Type:      notification.Type,
			Data:      notification.Data.String(),
			CreatedAt: common.TimeToRPCString(notification.CreatedAt),
			ReadAt:    common.NullTimeToRPCString(notification.ReadAt),
		}
	}

	return &rpc.NotificationNotificationsResponse{
		Notifications: rpcNotifications,
	}, nil
}

func (s *notificationService) Clean(ctx context.Context, r *rpc.NotificationCleanRequest) (*rpc.Empty, error) {
	me := s.getMe(ctx)
	err := s.repos.Notification.Clean(me.ID, r.LastKnownId)
	if err != nil {
		return nil, err
	}
	return &rpc.Empty{}, nil
}

func (s *notificationService) MarkRead(ctx context.Context, r *rpc.NotificationMarkReadRequest) (*rpc.Empty, error) {
	me := s.getMe(ctx)
	err := s.repos.Notification.MarkRead(me.ID, r.LastKnownId)
	if err != nil {
		return nil, err
	}
	return &rpc.Empty{}, nil
}
