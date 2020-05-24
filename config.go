package main;

import (
  "fmt"
  "flag"
  "os"
);

type Config struct {
  token string;
  prefix string;
}

var CONFIG Config = Config{
  token: os.Getenv("DISCORD_API_TOKEN"),
  prefix: "!",
};

func init() {
  token := "";
  flag.StringVar(&token, "t", "", "Discord API bot token");
  flag.StringVar(&CONFIG.prefix, "p", CONFIG.prefix, "Bot command prefix");
  help := flag.Bool("h", false, "Print the help message");
  flag.Parse();
  if len(token) > 0 {
    CONFIG.token = token;
  }

  if *help || len(CONFIG.token) == 0 {
    fmt.Println("Help message");
    os.Exit(0);
  }
}
