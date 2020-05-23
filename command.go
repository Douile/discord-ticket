package main;

import (
  "github.com/bwmarrin/discordgo"
);

type Command struct {
  Call func(*discordgo.Session,*discordgo.MessageCreate);
  Check func(*discordgo.Session,*discordgo.MessageCreate) bool;
  Help string;
};

func CheckNop(s *discordgo.Session,m *discordgo.MessageCreate) bool {
  return true;
};
