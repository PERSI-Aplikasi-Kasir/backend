package service

import (
	"backend/common/utils"
	"backend/microservices/whatsappclient/chore/entity"
	"backend/microservices/whatsappclient/chore/interfaces"
	"backend/microservices/whatsappclient/config"
	"context"
	"errors"
	"fmt"

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

	isIDStored := s.client.Store.ID != nil
	isAutenticated := s.client.IsLoggedIn()

	return isIDStored && isAutenticated, nil
}

func (s *whatsappService) GetLoginQR() (<-chan *[]byte, error) {
	isLogin, err := s.CheckDevice()
	if err != nil {
		log.Error().Err(err).Msg("Failed to check device")
		return nil, err
	}

	if isLogin {
		log.Info().Msg("Client is already logged in")
		return nil, nil
	}

	if s.client != nil {
		config.ResyncClient(&s.client)
	}

	qrChan, err := s.client.GetQRChannel(context.Background())
	if err != nil {
		if errors.Is(err, whatsmeow.ErrQRStoreContainsID) {
			_ = s.client.Connect()
			if s.client.IsLoggedIn() {
				log.Info().Msg("Client is already logged in")
				return nil, nil
			}
			log.Error().Err(err).Msg("Failed to get QR channel")
			return nil, err
		} else {
			log.Error().Err(err).Msg("Failed to get QR channel")
			return nil, err
		}
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

	err = s.client.Connect()
	if err != nil {
		log.Error().Err(err).Msg("Failed to sync client")
		return nil, err
	}

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
		byteData, err := utils.LoadImage(req.ImageID)
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

func (s *whatsappService) ResetLoggedDevice() error {
	if s.client == nil {
		log.Error().Msg("Client is not initialized")
		return fmt.Errorf("client is not initialized")
	}

	s.client.Store.Container.DeleteDevice(s.client.Store)
	s.client.Store.Delete()
	s.client.Disconnect()

	log.Info().Msg("Logged device has been reset")
	return nil
}
