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
  // Register the commands to COMMANDS map
  RegisterCommands();

  fmt.Println("Logging in...");

  // Initialize discord class
  dg, err := discordgo.New(fmt.Sprintf("Bot %s", CONFIG.token));
  if err != nil {
    fmt.Println("Error with discordgo", err);
    return;
  }

  // Add event handlers
  dg.AddHandler(OnMessage);
  dg.AddHandler(OnReady);
  dg.AddHandler(OnReactAdd);

  // Connect to websocket
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
