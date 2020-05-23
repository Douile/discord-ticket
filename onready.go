package main;

import (
  "fmt"

  "github.com/bwmarrin/discordgo"
)

func OnReady(session *discordgo.Session, ready *discordgo.Ready) {
  fmt.Printf("Successfully logged in as %s [%s]\n", session.State.User.Username, session.State.User.ID);

  session.UpdateStatusComplex(discordgo.UpdateStatusData{
    Game: &discordgo.Game{
      Name: "discord-ticket",
      Type: discordgo.GameTypeGame,
    },
  });
}
