package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"

	proto "github.com/bluefoxcode/rome/services/Notekeeper/proto"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/registry/nats"
	_ "github.com/micro/go-plugins/transport/nats"
	"google.golang.org/grpc/metadata"
)

// Race is model for exported race
type Race struct {
	RaceID      string `json:"raceID"`
	RaceName    string `json:"raceName"`
	First       string `json:"first"`
	Last        string `json:"last"`
	Party       string `json:"party"`
	CandidateID string `json:"candidateID"`
	VoteCount   int32  `json:"voteCount"`
}

// Races is list of races
type Races []Race

func main() {
	cmd.Init()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ticker := time.NewTicker(time.Second * 2)
	races := initRaces()

	go updateVotes(ticker, &races)

	wg.Wait()

}

func initRaces() Races {
	return Races{
		{"0",
			"President of the United States",
			"Donald",
			"Trump",
			"GOP",
			"9677",
			0,
		},
		{"0",
			"President of the United States",
			"Hillary",
			"Clinton",
			"Dem",
			"9678",
			0,
		},
		{"0",
			"President of the United States",
			"Gary",
			"Johnson",
			"Lib",
			"9675",
			0,
		},
		{"0",
			"President of the United States",
			"Jill",
			"Stein",
			"Grn",
			"9676",
			0,
		}}
}

func updateVotes(ticker *time.Ticker, races *Races) {

	for range ticker.C {
		if rand.Intn(10) > 6 {
			n := rand.Intn(len(*races))
			(*races)[n].VoteCount += rand.Int31() % 1000
			publish((*races)[n])
		}
	}
}

func publish(r Race) {

	ctx := metadata.NewContext(context.Background(), metadata.MD{"X-User-Id": []string{"Elections_AP"}})

	msg := client.NewPublication("go.micro.srv.Notekeeper.Race", &proto.Race{
		RaceID:      r.RaceID,
		RaceName:    r.RaceName,
		First:       r.First,
		Last:        r.Last,
		Party:       r.Party,
		CandidateID: r.CandidateID,
		VoteCount:   r.VoteCount,
		Source:      "AP",
	})

	if err := client.Publish(ctx, msg); err != nil {
		log.Println("publish err: ", err)
	}

}
