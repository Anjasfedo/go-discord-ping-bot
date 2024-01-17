package bot // Declare a bot package, which handles Discord bot functionality.

import (
	"fmt"                                             // Import the "fmt" package for formatted I/O.
	"github.com/Anjasfedo/go-discord-ping-bot/config" // Import the custom config package.
	"github.com/bwmarrin/discordgo"                   // Import the "discordgo" package for Discord bot functionality.
)

var BotID string             // Declare a global variable to store the ID of the bot user.
var GoBot *discordgo.Session // Declare a global variable to store the Discord bot session.

// Start function initializes and starts the Discord bot.
func Start() {
	// Create a new Discord bot session using the bot token from the configuration.
	GoBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error()) // Print an error message if there is an issue creating the bot session.
		return
	}

	// Retrieve information about the bot user (e.g., BotID).
	user, err := GoBot.User("@me")
	if err != nil {
		fmt.Println(err.Error()) // Print an error message if there is an issue retrieving bot user information.
		return
	}

	BotID = user.ID // Set the global variable BotID with the ID of the bot user.

	GoBot.AddHandler(messageHandler) // Register the messageHandler function to handle incoming messages.

	err = GoBot.Open() // Open a connection to the Discord gateway.
	if err != nil {
		fmt.Println(err.Error()) // Print an error message if there is an issue opening the connection.
		return
	}

	fmt.Println("Bot is Running") // Print a message indicating that the bot is running.
}

// messageHandler function is the handler for incoming Discord messages.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return // Ignore messages sent by the bot itself.
	}

	if m.Content == "ping" {
		// Respond with "pong" to messages containing "ping".
		_, err := s.ChannelMessageSend(m.ChannelID, "pong")
		if err != nil {
			fmt.Println("Error sending message:", err) // Print an error message if there is an issue sending the response.
		}
	}

	if m.Content == "anjas" {
		// Respond with "Anjas Gantenk" to messages containing "anjas".
		_, err := s.ChannelMessageSend(m.ChannelID, "Anjas Gantenk")
		if err != nil {
			fmt.Println("Error sending message:", err) // Print an error message if there is an issue sending the response.
		}
	}
}
