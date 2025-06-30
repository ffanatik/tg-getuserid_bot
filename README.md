# tg-getuserid_bot

Super simple telegram bot allowing the user to get their UserID written in Go

## Description

A Telegram bot that helps users retrieve their Telegram User ID. Simply start the bot and it will display your unique User ID.

## Prerequisites

- Go 1.24 or higher
- Telegram Bot Token (get it from @BotFather)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/ffanatik/tg-getuserid_bot.git
cd tg-getuserid_bot
```
2. Set your bot token:
```bash
export TG_BOT_TOKEN="your_bot_token_here"
```
3. Build the application:
```bash
go run main.go
```
or build and run with Docker:
```bash
docker build -t tg-getuserid-bot .
docker run -e TG_BOT_TOKEN="your_bot_token_here" tg-getuserid-bot
```
## Usage

1. Start a chat with your bot
2. Send /start command
3. Get your Telegram User ID

## You can find the bot here:
[https://t.me/letsgetuserid_bot](https://t.me/letsgetuserid_bot)