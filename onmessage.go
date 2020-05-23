package main;

import (
  "strings"

  "github.com/bwmarrin/discordgo"
)

func OnMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
  if message.Author.ID == session.State.User.ID { return; }

  commandName := strings.TrimPrefix(message.Content, PREFIX);
  if (commandName == message.Content) { return; } // Doesn't start with prefix
  commandParts := strings.Split(commandName, " ");
  commandName = strings.ToLower(commandParts[0]);

  command, exists := COMMANDS[commandName];
  if exists {
    if command.Check(session, message) {
      command.Call(session, message);
    } else {
      session.ChannelMessageSend(message.ChannelID, "Sorry you don't have permission to use this command");
    }
  }
}
