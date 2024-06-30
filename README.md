# Webapper - Telegram Mini Apps URL Embedder

When developing Telegram Mini Apps, you need a domain and an HTTPS connection between your server and Telegram servers. This binary/docker container helps you send a message with a Telegram Mini App embedded with any URL you choose.

## Usage

Run the binary/docker container with the `TOKEN` environment variable set to your Telegram Bot Token. Then send `/start` to your bot.

You can also set the `LOG` environment variable to any value if you need a detailed log of Telegram requests to your bot.