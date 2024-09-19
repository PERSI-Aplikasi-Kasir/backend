package interfaces

import (
	"backend/microservices/whatsappclient/chore/entity"
)

type WhatsappService interface {
	CheckDevice() (bool, error)
	GetLoginQR() (<-chan *[]byte, error)
	SendMessage(req *entity.MessageSend) error
}
