package config // Declare a config package, which handles configuration settings.

import (
	"encoding/json" // Import the "encoding/json" package for JSON encoding and decoding.
	"fmt"           // Import the "fmt" package for formatted I/O.
	"os"            // Import the "os" package for operating system functions.
)

var (
	Token     string        // Declare a global variable to store the Discord bot token.
	BotPrefix string        // Declare a global variable to store the bot command prefix.
	config    *configStruct // Declare a global variable to store the configuration settings.
)

// Define a struct to represent the structure of the configuration file.
type configStruct struct {
	Token      string `json:"Token"`     // Token field in the configuration file.
	BotkPrefix string `json:"BotPrefix"` // BotPrefix field in the configuration file.
}

// ReadConfig function reads the configuration from the "config.json" file.
func ReadConfig() error {
	fmt.Println("Read config.json") // Print a message indicating that the configuration is being read.

	// Read the contents of the "config.json" file.
	file, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error()) // Print an error message if there is an issue reading the file.
		return err
	}

	fmt.Println(string(file)) // Print the content of the configuration file.

	// Unmarshal the JSON data from the file into the configStruct variable.
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error()) // Print an error message if there is an issue unmarshalling the JSON data.
		return err
	}

	Token = config.Token          // Set the global variable Token with the value from the configuration file.
	BotPrefix = config.BotkPrefix // Set the global variable BotPrefix with the value from the configuration file.

	return nil // Return nil to indicate that the configuration reading was successful.
}
