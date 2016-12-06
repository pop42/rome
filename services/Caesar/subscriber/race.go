package subscriber

import (
	"net/http"

	"github.com/bluefoxcode/rome/services/Caesar/lib/util"
	"github.com/bluefoxcode/rome/services/Caesar/model/race"
	proto "github.com/bluefoxcode/rome/services/Notekeeper/proto"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

// Race is a struct that contains Race handlers
type Race struct {
	Client client.Client
}

// Handle will respond to relevant messages on the topic it is registered
func (e *Race) Handle(ctx context.Context, msg *proto.Race) error {

	item := race.Item{
		RaceID:      msg.RaceID,
		RaceName:    msg.RaceName,
		First:       msg.First,
		Last:        msg.Last,
		Party:       msg.Party,
		CandidateID: msg.CandidateID,
		VoteCount:   msg.VoteCount,
		Source:      msg.Source,
	}
	var w http.ResponseWriter
	var req *http.Request
	c := util.Context(w, req)
	race.Create(c.DB, item)
	return nil
}
