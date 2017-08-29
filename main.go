package main

import (
	"flag"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/nhooyr/color/log"
	"github.com/Pallinder/go-randomdata"
	"github.com/AllenDang/w32"
	"syscall"
        "unsafe"
)

var (
	email    = flag.String("email", "", "account email")
	pass     = flag.String("pass", "", "account password")
	guild    = flag.String("guild", "", "guild (server) to join")
	channel  = flag.String("chan", "", "channel to join")
	message  = flag.String("msg", randomdata.Country(randomdata.FullCountry), "message to be sent")
	interval = flag.Int64("int", 60, "interval between messages in seconds")
)

func main() {
	
	flag.Parse()
	if *email == "" || *pass == "" {
		log.Fatal("please provide an email and password")
	}
	s, err := discordgo.New(*email, *pass)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("logged in")

	g := findGuild(s)
	if g == nil {
		log.Fatal("could not find guild")
	}
	id := findChannel(s, g)
	if id == "" {
		log.Fatal("could not find channel")
	}
	s.Open()
	    s.AddHandler(messageCreate)
		noAdmin := true
	for t := time.Tick(time.Duration(*interval) * time.Second); ; <-t {
		newMessage  := randomdata.Country(randomdata.FullCountry)

		
		if members, err := s.GuildMembers(id, "", 200); err == nil {
			for _, e := range members {
				log.Print("Searched for members")
  					for _, e2 := range e.Roles{
						log.Print("Searched for roles")
						log.Print(e2)
							if e2 == "God" {
								noAdmin = false
								log.Print("Set to false")
						
						}
 					 }
					
 				 }
			}
		
		if noAdmin == true {
		if _, err := s.ChannelMessageSend(id, newMessage); err != nil {
			

			log.Print(err)
		} else {
			log.Print("sent message")
		}
			channeler, _ := s.Channel(findChannel(s, g))
			LastMessageID := channeler.LastMessageID
			s.ChannelMessageDelete(id, LastMessageID)
		} else {
			log.Print("Not Doing Anything becasue an admin is online")
		}
		
	}
}

func MakeIntResource(id uint16) (*uint16) {
    return (*uint16)(unsafe.Pointer(uintptr(id)))
}

func WndProc(hWnd w32.HWND, msg uint32, wParam, lParam uintptr) (uintptr) {
switch msg {
case w32.WM_DESTROY:
        w32.PostQuitMessage(0)
    default:
        return w32.DefWindowProc(hWnd, msg, wParam, lParam)
    }
    return 0
}
func WinMain() int {
	
	hInstance := w32.GetModuleHandle("")
	lpszClassName := syscall.StringToUTF16Ptr("WNDclass")
	var wcex w32.WNDCLASSEX
	wcex.Size            = uint32(unsafe.Sizeof(wcex))
	wcex.Style         = w32.CS_HREDRAW | w32.CS_VREDRAW
	wcex.WndProc       = syscall.NewCallback(WndProc)
	wcex.ClsExtra        = 0
	wcex.WndExtra        = 0
	wcex.Instance         = hInstance
	wcex.Icon         = w32.LoadIcon(hInstance, MakeIntResource(w32.IDI_APPLICATION))
	wcex.Cursor       = w32.LoadCursor(0, MakeIntResource(w32.IDC_ARROW))
	wcex.Background = w32.COLOR_WINDOW + 11
	
	wcex.MenuName  = nil

	wcex.ClassName = lpszClassName
	wcex.IconSm       = w32.LoadIcon(hInstance, MakeIntResource(w32.IDI_APPLICATION))
	w32.RegisterClassEx(&wcex)
	hWnd := w32.CreateWindowEx(
	0, lpszClassName, syscall.StringToUTF16Ptr("Simple Go Window!"), 
	w32.WS_OVERLAPPEDWINDOW | w32.WS_VISIBLE, 
	w32.CW_USEDEFAULT, w32.CW_USEDEFAULT, 400, 400, 0, 0, hInstance, nil)
	w32.ShowWindow(hWnd, w32.SW_SHOWDEFAULT)
	w32.UpdateWindow(hWnd)
   var msg w32.MSG
   for {
        if w32.GetMessage(&msg, 0, 0, 0) == 0 {
            break
        }
        w32.TranslateMessage(&msg)
        w32.DispatchMessage(&msg)
   }
   return int(msg.WParam)
	
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		s.ChannelMessageSend(m.ChannelID, "Bob")
		   WinMain()
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}

func findGuild(s *discordgo.Session) *discordgo.UserGuild {
	gs, err := s.UserGuilds(0, "", "")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("got guilds")
	for _, g := range gs {
		if g.Name == *guild {
			log.Print("found guild")
			return g
		}
	}
	return nil
}

func findChannel(s *discordgo.Session, g *discordgo.UserGuild) string {
	chs, err := s.GuildChannels(g.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("got channels")
	for _, ch := range chs {
		if ch.Name == *channel {
			log.Print("found channel")
			return ch.ID
		}
	}
	return ""
}
