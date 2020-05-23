package main;

import (
  "fmt"
  "runtime"

  "github.com/bwmarrin/discordgo"
);

var CommandHello Command = Command{
  Call: func(session *discordgo.Session,message *discordgo.MessageCreate) {
    session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
      Title: "Hi there!",
      Description: "I'm but a simple discgobot.",
    });
  },
  Check: CheckNop,
  Help: "Hello!",
};

var CommandInfo Command = Command{
  Call: func(session *discordgo.Session,message *discordgo.MessageCreate) {
    session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
      Author: &discordgo.MessageEmbedAuthor{
        Name: "Info",
        IconURL: session.State.User.AvatarURL(""),
      },
      Description: fmt.Sprintf("%s\ndiscordgo: v%s", runtime.Version(), discordgo.VERSION),
    });
  },
  Check: CheckNop,
  Help: "Get runtime information",
}

var CommandHelp Command = Command{
  Call: func(session *discordgo.Session,message *discordgo.MessageCreate) {
    helpMessage := "";
    for k, v := range COMMANDS {
      helpMessage += fmt.Sprintf("`%s%s` - %s\n", PREFIX, k, v.Help);
    }

    session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
      Author: &discordgo.MessageEmbedAuthor{
        Name: "Help",
        IconURL: session.State.User.AvatarURL(""),
      },
      Description: helpMessage,
    });
  },
  Check: CheckNop,
  Help: "Get help with other commands",
};

func RegisterCommands() {
  COMMANDS["hello"] = CommandHello;
  COMMANDS["help"] = CommandHelp;
  COMMANDS["info"] = CommandInfo;
}
