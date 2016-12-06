package subscriber

import (
	"log"

	proto "github.com/bluefoxcode/rome/services/Notekeeper/proto"
	"github.com/bluefoxcode/rome/services/Notekeeper/publisher"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

// Race is a struct that contains Race handlers
type Race struct {
	Client client.Client
}

// Handle will respond to relevant messages on the topic it is registered
func (e *Race) Handle(ctx context.Context, msg *proto.Race) error {
	log.Print("Handler received race data. Publishing...", msg)
	publisher.PublishRace(msg)
	return nil
}
