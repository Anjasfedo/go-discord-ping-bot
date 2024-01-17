package main // Declare a main package, which is the entry point of the program.

import (
	"fmt"                                             // Import the "fmt" package for formatted I/O.
	"github.com/Anjasfedo/go-discord-ping-bot/bot"    // Import the custom bot package.
	"github.com/Anjasfedo/go-discord-ping-bot/config" // Import the custom config package.
)

func main() {
	err := config.ReadConfig() // Call the ReadConfig function from the config package.
	if err != nil {
		fmt.Println(err.Error()) // Print an error message if there is an issue reading the configuration.
		return                   // Exit the program if there is an error.
	}

	bot.Start() // Start the Discord bot by calling the Start function from the bot package.

	<-make(chan struct{}) // Block the main goroutine by waiting on a channel. This is used to keep the program running indefinitely.
}
