package loader

import (
	"fmt"
	"github.com/babofitos/babogo00/lib/commands"
	irc "github.com/fluffle/goirc/client"
	"reflect"
	"strings"
)

var com = new(commands.Command)

func LoadCommand(conn *irc.Conn, line *irc.Line) {
	msg := line.Text()
	msgSplit := strings.Split(msg, " ")
	command := msgSplit[0][1:]
	command = strings.Title(command)
	args := msgSplit[1:]
	argsValue := reflect.ValueOf(args)
	connValue := reflect.ValueOf(conn)
	reflect.ValueOf(com).MethodByName(command).Call([]reflect.Value{connValue, argsValue})
}
