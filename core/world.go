package core

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/fsnotify/fsnotify"
	"github.com/gliderlabs/ssh"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

type World struct {
	srv    *ssh.Server
	PubSub *gochannel.GoChannel
}

func NewWorld() *World {
	return &World{}
}

func (w *World) Init() {
	w.setupConfig()
	w.setupLogger()
	w.outputConfig()
	w.setupPubSub()
	w.setupSSHServer()
	w.startSSHServer()
}

func (w *World) setupConfig() {
	// Read configuration
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("fatal error config file")
	}

	// Update configuration on change
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Info().Str("file", e.Name).Msg("Config file changed")
		w.outputConfig()
	})
	viper.WatchConfig()
}

func (w *World) setupLogger() {
	// Setup logger
	zerolog.SetGlobalLevel(zerolog.InfoLevel) // Default to Info

	// Attempt to parse the log_level from the config file
	logLevel, err := zerolog.ParseLevel(viper.GetString("log_level"))
	if err != nil {
		log.Fatal().Err(err).Msg("unable to parse log_level")
	} else {
		zerolog.SetGlobalLevel(logLevel)
	}

	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func (w *World) outputConfig() {
	// Output configuration
	for _, key := range viper.AllKeys() {
		log.Debug().
			Str("key", key).
			Str("value", viper.GetString(key)).
			Msg("Config")
	}
}

func (w *World) setupPubSub() {
	// Setup PubSub
	w.PubSub = gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(true, false),
	)

	playerMessages, err := w.PubSub.Subscribe(context.Background(), "player")
	if err != nil {
		log.Fatal().Err(err).Msg("Error subscribing to player messages")
	}

	roomMessages, err := w.PubSub.Subscribe(context.Background(), "room")
	if err != nil {
		log.Fatal().Err(err).Msg("Error subscribing to room messages")
	}

	go process(playerMessages)
	go process(roomMessages)

}

func (w *World) setupSSHServer() {
	serverHost := viper.GetString("server.host")
	serverPort := viper.GetString("server.port")
	serverIdleTimeout := viper.GetDuration("server.idle_timeout")

	srv := &ssh.Server{
		Addr:        net.JoinHostPort(serverHost, serverPort),
		IdleTimeout: serverIdleTimeout,
		ConnectionFailedCallback: func(conn net.Conn, err error) {
			log.Error().Err(err).Msg("Connection failed")
			conn.Close()
		},
		Handler: w.Handler,
	}

	w.srv = srv
}

func (w *World) startSSHServer() {
	// Start SSH server
	serverHost := viper.GetString("server.host")
	serverPort := viper.GetString("server.port")
	log.Info().Str("host", serverHost).Str("port", serverPort).Msg("Starting server")

	if err := w.srv.ListenAndServe(); err != nil {
		log.Error().Err(err).Msg("Error starting server")
	}
}

func process(messages <-chan *message.Message) {
	for msg := range messages {
		fmt.Printf("received message: %s, payload: %s\n", msg.UUID, string(msg.Payload))
		msg.Ack()
	}
}

func (w *World) Handler(s ssh.Session) {
	p := &Player{
		PubSub: w.PubSub,
		Account: &Account{
			// Username: "Jasrags",
			Characters: []string{"Jasrags"},
		},
		Character: Character{
			Name: "Jasrags",
		},
		Session: s,
	}

	p.MoveTo(TheVoid.Rooms[0], func() {})

	log.Info().
		Str("user", s.User()).
		Str("remote_addr", s.RemoteAddr().String()).
		Str("session_id", s.Context().SessionID()).
		Msg("New SSH connection")

	// w.Users[s.Context().SessionID()] = u
	// u.Room.Enter(s)
	// u.Room.Render(s)

	t := term.NewTerminal(s, "> ")
	for {

		input, err := t.ReadLine()
		if err != nil {
			log.Error().Err(err).Str("user", s.User()).Msg("Error reading input")
			break
		}

		if input != "" {
			log.Debug().Str("user", s.User()).Str("input", input).Msg("Received input")

			// cmd := w.parseCommand(s.Context().SessionID(), input)
			// w.handleCommand(cmd)
		}
	}
}

// func (w *World) handleCommand(cmd CommandMsg) {
// 	switch cmd.Command {
// 	case "say":
// 	case "echo":
// 		io.WriteString(w.Users[cmd.FromID].Session, fmt.Sprintf("ECHO: %s\n", strings.Join(cmd.Args, " ")))
// 	case "help":
// 		io.WriteString(w.Users[cmd.FromID].Session, "Help message\n")
// 	case "exit":
// 		io.WriteString(w.Users[cmd.FromID].Session, "Goodbye\n")
// 		w.Users[cmd.FromID].Session.Close()
// 		defer delete(w.Users, cmd.FromID)
// 	default:
// 		io.WriteString(w.Users[cmd.FromID].Session, "Unknown command\n")
// 	}
// }

// func (w *World) parseCommand(id, input string) CommandMsg {
// 	// Split the input into a command and arguments
// 	args := strings.Fields(input)
// 	if len(args) == 0 {
// 		return CommandMsg{FromID: id}
// 	}

// 	return CommandMsg{
// 		FromID:  id,
// 		Command: strings.ToLower(args[0]),
// 		Args:    args[1:],
// 	}
// }

// type CommandMsg struct {
// 	FromID  string
// 	Command string
// 	Args    []string
// }

// type User struct {
// 	Session   ssh.Session
// 	Character struct {
// 		Name string
// 	}
// 	// Room *core.Room
// }

// func (r *Room) Render(s ssh.Session) {
// 	// userCount := len(r.Users)
// 	io.WriteString(s, cfmt.Sprintf("{{%s\n}}::green", r.Title))
// 	io.WriteString(s, cfmt.Sprintf("{{%s\n}}::bold", wordwrap.WrapString(r.Description, uint(viper.GetInt("word_wrap_min")))))
// 	// io.WriteString(s, cfmt.Sprintf("There are {{%d}}::yellow|bold other %s here\n", userCount, pluralise.WithCount("user", userCount)))
// }

// func (r *Room) Enter(s ssh.Session) {
// 	// r.Users[s.Context().SessionID()] = s
// }

// func (r *Room) Leave(s ssh.Session) {
// 	// delete(r.Users, s.Context().SessionID())
// }

// func (r *Room) Say(s ssh.Session, msg string) {
// 	io.WriteString(s, cfmt.Sprintf("You said '%s'\n", msg))

// 	for _, user := range r.Users {
// 		if user != s {
// 			io.WriteString(user, msg)
// 		}
// 	}
// }

// var theVoid = &Room{
// 	// Users:       make(map[string]ssh.Session),
// 	Title:       "The Void",
// 	Description: "You are floating in a formless void, detached from all sensation of physical matter, surrounded by swirling glowing light, which fades into the relative darkness around you without any trace of edges or shadow.",
// }
