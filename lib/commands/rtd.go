package commands

import (
	"fmt"
	"github.com/babofitos/babogo00/config"
	irc "github.com/fluffle/goirc/client"
	"math/rand"
	"time"
)

func (c *Command) Rtd(conn *irc.Conn, args []string) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(6) + 1
	conn.Privmsg(config.Get().Chan, fmt.Sprintf("%v", r))
}
