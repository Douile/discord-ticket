package main;

import (
  "strings"

  "github.com/bwmarrin/discordgo"
)

func OnMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
  if message.Author.ID == session.State.User.ID { return; }

  commandName := strings.TrimPrefix(message.Content, CONFIG.prefix);
  if (commandName == message.Content) { return; } // Doesn't start with prefix TODO: Maybe use HasPrefix
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
