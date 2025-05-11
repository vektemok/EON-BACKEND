package model

const (
	ModerationPending  = "pending"  // — отправлена на модерацию
	ModerationApproved = "approved" // — одобрена, можно публиковать
	ModerationRejected = "rejected" // — отклонена модератором
	ModerationDisabled = "disabled" // — отключена вручную или по другим причинам (опционально)
)

var ValidModerationStatuses = map[string]bool{
	ModerationPending:  true,
	ModerationApproved: true,
	ModerationRejected: true,
	ModerationDisabled: true,
}

