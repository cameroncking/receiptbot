# ReceiptBot

ReceiptBot is a Telegram bot designed to images and chat messages for tax
purposes. It organizes the received images and messages into user-specific
folders, making it easier to manage and retrieve them at the end of the year.

## Features

- Collects images and chat messages from both private chats and group chats.
- Supports images sent as photos and files.
- Organizes images and messages into user-specific folders.
- Appends chat messages to a `memo.txt` file in the respective folder.
- Runs in a Docker container for easy deployment.

## Prerequisites

- Go 1.20 or later
- Docker
- Docker Compose

## Project Structure

```
receiptbot/
├── main.go
├── bot/
│   └── bot.go
├── handlers/
│   └── handlers.go
├── utils/
│   └── utils.go
├── go.mod
├── go.sum
├── Dockerfile
├── docker-compose.yaml
├── .env
└── README.md
└── images/
    └── .dir
```

## Setup

1. **Clone the Repository:**

   ```sh
   git clone https://github.com/yourusername/receiptbot.git
   cd receiptbot
   ```

2. **Set Up Environment Variables:**
   Create a `.env` file in the root of your project directory and add your
   Telegram bot token:

   ```
   TELEGRAM_BOT_TOKEN=your_telegram_bot_token_here
   ```

3. **Build and Run the Docker Container:**

   ```sh
   docker-compose build
   docker-compose up
   ```

## Usage

1. **Add the Bot to a Group Chat:**
   - Invite the bot to a group chat by searching for its username and adding it
     to the group.
   - Ensure the bot has the necessary permissions to read messages in the
     group.
2. **Upload Images and Messages:**
   - Upload images as photos or files.
   - Send chat messages with or without captions.
   - The bot will organize the images and messages into user-specific folders.

## Code Overview

### `main.go`

The entry point of the application. Reads the API key from an environment
variable and starts the bot.

### `bot/bot.go`

Handles the initialization and update processing of the bot.

### `handlers/handlers.go`

Handles the logic for processing photos and text messages. Organizes images and
messages into user-specific folders.

### `utils/utils.go`

Contains utility functions for folder creation, file downloading, and appending
text to files.

### Dockerfile

Defines the Docker image for the bot. Builds the Go application and sets up the environment.

### docker-compose.yaml

Defines the Docker Compose configuration. Builds the Docker image and runs the
bot in a container.

### .env.sample

Contains a sample of .env, which contains environment variables for the bot,
such as the Telegram bot token.

Don't forget to edit .env to contain your own bot token!

```
cp .env.sample .env
```

### images/

This is the output folder where images will be stored.  docker-compose.yaml
specifies this as a mount path so that saved images can be accessed from
outside the container.

## Vibe Code Warning

This project was a fun tinker project that I did in two hours by 95% vibe
coding it.  The code is hot garbage.  It maintains all state in global
variables and probably has massive flaws that will cause it to delete all of
the files on your hard drive.

It was my expressed intent that I would not use my brain AT ALL if I could help
it.  I refused to even run `go get` to download packages until the LLM (Mistral
Nemo) identified the error and gave me the command to paste.

Here are my observations:

- AI is great at standing up low-quality proof-of-concept code and scaffolding.
- AI is really bad at business logic.  It's much concise (and easier IMHO) to
express the business logic myself as code than to try and explain the nuances
of my desired behavior to the LLM.
- AI really struggled with bugfixing multiple overalpping issues.  I ran into
an issue where I was trying to coax it to fix two bugs, and it could do so -
but it would always cause the other bug to regress.  The fix only spanned 3
lines in one file, but the LLM couldn't get both fixes to coexist.
- AI recommends insecure practices by default.  I had to coach it on something
as simple as moving api keys to .env.
- AI generates garbage code with redundant lines.  Golang cares about unused
imports, for example.  Every single file the LLM generated imported "log", but
it didn't put any log messages anywhere.  Therefore I had to discuss this with
the LLM for each file after showing it the error message, one file at a time.
- AI generated poor-quality code.  This would never pass code review for
anything more valuable than a proof of concept.

It was fun and informative though, and I got a tool out of it.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the ISC License. See the [LICENSE](LICENSE) file
for details.

