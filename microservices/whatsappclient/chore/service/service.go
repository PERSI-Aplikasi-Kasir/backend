package service

import (
	"backend/common/utils"
	"backend/microservices/whatsappclient/chore/entity"
	"backend/microservices/whatsappclient/chore/interfaces"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog/log"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

var _ interfaces.WhatsappService = &whatsappService{}

type whatsappService struct {
	client *whatsmeow.Client
}

func NewWhatsappService(client *whatsmeow.Client) *whatsappService {
	return &whatsappService{
		client: client,
	}
}

func (s *whatsappService) CheckDevice() (bool, error) {
	if s.client == nil {
		log.Error().Msg("Client is not initialized")
		return false, fmt.Errorf("client is not initialized")
	}

	return s.client.Store.ID != nil, nil
}

func (s *whatsappService) GetLoginQR() (<-chan *[]byte, error) {
	if s.client != nil {
		s.client.Disconnect()
	}

	qrChan, err := s.client.GetQRChannel(context.Background())
	if err != nil {
		log.Error().Err(err).Msg("Failed to get QR channel")
		return nil, err
	}

	err = s.client.Connect()
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to WhatsApp")
		return nil, err
	}

	out := make(chan *[]byte)

	go func() {
		defer close(out)
		for evt := range qrChan {
			switch evt.Event {
			case "success":
				out <- nil
				return
			case "timeout":
				out <- nil
				return
			case "code":
				png, err := utils.GenerateQRCode(evt.Code)
				if err != nil {
					log.Error().Err(err).Msg("Failed to generate QR code image")
					continue
				}
				out <- png
			}
		}
	}()

	return out, nil
}

func (s *whatsappService) SendMessage(req *entity.MessageSend) error {
	logged, err := s.CheckDevice()
	if err != nil {
		log.Error().Err(err).Msg("Failed to check device")
		return err
	}
	if !logged {
		log.Error().Msg("Device is not logged in")
		return fmt.Errorf("device is not logged in")
	}

	targetJID := types.NewJID(utils.ParsePhoneNumber(req.To), types.DefaultUserServer)

	var message waE2E.Message
	if req.ImageID == "" {
		message.Conversation = proto.String(req.Message)
	} else {
		byteData, err := loadImage(req.ImageID)
		if err != nil {
			log.Error().Err(err).Msg("Failed to load image")
			return err
		}

		uploadRes, _ := s.client.Upload(context.Background(), *byteData, whatsmeow.MediaImage)
		message.ImageMessage = &waE2E.ImageMessage{
			Caption:       proto.String("Hello, world!"),
			Mimetype:      proto.String("image/png"),
			URL:           &uploadRes.URL,
			DirectPath:    &uploadRes.DirectPath,
			MediaKey:      uploadRes.MediaKey,
			FileEncSHA256: uploadRes.FileEncSHA256,
			FileSHA256:    uploadRes.FileSHA256,
			FileLength:    &uploadRes.FileLength,
		}
	}

	res, err := s.client.SendMessage(context.Background(), targetJID, &message)
	if err != nil {
		log.Error().Err(err).Msg("Failed to send message")
		return err
	}

	log.Info().Msg(fmt.Sprintf("Message sent to %s at %s", req.To, res.Timestamp))
	return nil
}

func loadImage(imageName string) (*[]byte, error) {
	file, err := os.Open(imageName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to open image file")
		return nil, err
	}
	defer file.Close()

	byteData, err := io.ReadAll(file)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read image file")
		return nil, err
	}

	return &byteData, nil
}
