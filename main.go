package main

import (
	"flag"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/nhooyr/color/log"
	"math/rand"
	//"github.com/Pallinder/go-randomdata"
	//"os/exec"
	
)

var (
	email    = flag.String("email", "", "account email")
	pass     = flag.String("pass", "", "account password")
	guild    = flag.String("guild", "", "guild (server) to join")
	channel  = flag.String("chan", "", "channel to join")
	message  = flag.String("msg","yeet", "message to be sent")
	interval = flag.Int64("int", 30, "interval between messages in seconds")
)
// randomdata.Country(randomdata.FullCountry)
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
	//s.Open()
	   // s.AddHandler(messageCreate)
		noAdmin := true
	for t := time.Tick(time.Duration(*interval) * time.Minute); ; <-t {
		newMessage  := "yeet"
		
		for t2 := 0; t2 > rand.Intn(4); t2++ {
		
			minutesPass := rand.Intn(4)	
			
		for t3 := time.Tick(time.Duration(minutesPass) * time.Minute); ; <-t3 {
			
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
			//channeler, _ := s.Channel(findChannel(s, g))
			//LastMessageID := channeler.LastMessageID
		//	s.ChannelMessageDelete(id, LastMessageID)
		} else {
			log.Print("Not Doing Anything becasue an admin is online")
		}
		
	}
   }
   }
}

/**
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		s.ChannelMessageSend(m.ChannelID, "Bob")
	cmd := exec.Command("sleep", "5")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}


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
**/

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
