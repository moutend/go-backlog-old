package backlog

type ChangeLog struct {
	Field            string           `json:"field"`
	NewValue         string           `json:"newValue"`
	OriginalValue    string           `json:"originalValue"`
	AttachmentInfo   AttachmentInfo   `json:"attachmentInfo"`
	AttributeInfo    string           `json:"attributeInfo"`
	NotificationInfo NotificationInfo `json:"notificationInfo"`
}
