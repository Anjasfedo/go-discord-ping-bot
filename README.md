# Discord Ping Bot 🤖🚀

This Go program implements a simple Discord bot that responds to messages containing specific commands. The bot is designed to respond with "pong" when it receives the command "ping".

## Usage 🚀

1. Create a config.json file in the project root with the following format:

    ```config.json
    {
        "Token": "your_discord_bot_token",
        "BotPrefix": "!"
    }
    ```

2. run the program:

    ```
     go run main.go
    ```

The bot will now be running and responding to messages in Discord.

## Code Explanation 📜

### `main` Function 🚀

#### Load Configuration:

-   Calls the `ReadConfig` function from the `config` package to read the Discord bot token and command prefix from the `config.json` file.

#### Start Discord Bot:

-   Calls the `Start` function from the `bot` package to initialize and start the Discord bot.

### `config` Package 🌐

#### ReadConfig Function:

-   Reads the `config.json` file to retrieve the Discord bot token and command prefix.
-   Sets global variables `Token` and `BotPrefix` with the values from the configuration file.

### `bot` Package 🤖

#### Start Function:

-   Creates a new Discord bot session using the DiscordGo library.
-   Retrieves bot user information and sets the global variable `BotID` with the bot user ID.
-   Adds a message handler (`messageHandler`) to handle incoming messages.
-   Opens a connection to the Discord gateway and prints a message indicating that the bot is running.

#### messageHandler Function:

-   Ignores messages sent by the bot itself.
-   Responds with "pong" to messages containing the command "ping."
-   Responds with "Anjas Gantenk" to messages containing the command "anjas."

## Closing Notes 📝

Feel free to adjust the configuration, and if you encounter any issues or have suggestions for improvement, please open an issue or submit a pull request.

Happy coding! 🤖👨‍💻
