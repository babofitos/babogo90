package main

import (
	"fmt"
	"github.com/babofitos/babogo00/config"
	"github.com/babofitos/babogo00/lib"
	irc "github.com/fluffle/goirc/client"
	"strings"
)

func main() {
	botcfg := config.Init()
	cfg := irc.NewConfig("baboGO00")
	cfg.SSL = false
	cfg.Server = botcfg.Server
	c := irc.Client(cfg)
	c.HandleFunc("connected",
		func(conn *irc.Conn, line *irc.Line) {
			ck := fmt.Sprintf("%s %s", botcfg.Chan, botcfg.Key)
			conn.Join(ck)
		})
	c.HandleFunc("privmsg",
		func(conn *irc.Conn, line *irc.Line) {
			//TODO: regex this or something
			if strings.HasPrefix(line.Text(), ".") {
				loader.LoadCommand(conn, line)
			}
		})
	quit := make(chan bool)
	c.HandleFunc("disconnected",
		func(conn *irc.Conn, line *irc.Line) { quit <- true })
	if err := c.Connect(); err != nil {
		fmt.Printf("Connection error: %s\n", err)
	}
	<-quit
}
