package publisher

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"

	raceproto "github.com/bluefoxcode/rome/services/Notekeeper/proto"
	"github.com/micro/go-micro/client"
)

// PublishRace posts the race
func PublishRace(r *raceproto.Race) {

	ctx := metadata.NewContext(context.Background(), metadata.MD{"X-User-Id": []string{"Notekeeper"}})

	msg := client.NewPublication("go.db.Race.post", r)

	if err := client.Publish(ctx, msg); err != nil {
		log.Println("publish err: ", err)
	}

}
