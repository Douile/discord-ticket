package main;

import (
  "flag"
);

var (
  TOKEN string;
  PREFIX string;
)

func init() {
  flag.StringVar(&TOKEN, "t", "", "Discord API bot token");
  flag.StringVar(&TOKEN, "p", "!", "Bot command prefix");
  flag.Parse();
}
