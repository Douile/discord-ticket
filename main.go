package main;

import (
  "fmt"
  "os"
  "os/signal"
  "syscall"

  "github.com/bwmarrin/discordgo"
)

var COMMANDS = map[string]Command{};

func main() {
  RegisterCommands();

  fmt.Println("Logging in...");

  dg, err := discordgo.New(fmt.Sprintf("Bot %s", TOKEN));
  if err != nil {
    fmt.Println("Error with discordgo", err);
    return;
  }

  dg.AddHandler(OnMessage);
  dg.AddHandler(OnReady);

  err = dg.Open();
  if err != nil {
    fmt.Println("Error connecting to discord", err);
    return;
  }

  // Wait here until CTRL-C or other term signal is received.
  fmt.Println("Bot is now running.  Press CTRL-C to exit.");
  sc := make(chan os.Signal, 1);
  signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill);
  <-sc;

  fmt.Println("Shutting down...");

  // Cleanly close down the Discord session.
  dg.Close();
}
