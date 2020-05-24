package main;

import (
  "fmt"

  "github.com/bwmarrin/discordgo"
)

var SHOULD_RESPOND = map[string]string{};

func OnReactAdd(session *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
  if reaction.UserID == session.State.User.ID { return; } // Ignore self

  _, exists := SHOULD_RESPOND[reaction.MessageID];
  if exists {
    err := session.MessageReactionRemove(reaction.ChannelID, reaction.MessageID, reaction.Emoji.Name, reaction.UserID);
    if err != nil {
      fmt.Println("Error removing reaction ", err);
    }

    userName := reaction.UserID;
    user, err := session.User(reaction.UserID);
    if err == nil {
      userName = user.Username;
    }
    session.ChannelMessageSend(reaction.ChannelID, fmt.Sprintf("%s reacted", userName));
  }
}
