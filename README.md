# sendtelegram
![](https://img.shields.io/github/v/release/gstracquadanio/sendtelegram)
![](https://github.com/gstracquadanio/sendtelegram/workflows/main/badge.svg)

A Telegram replacement for sendmail.

`sendtelegram` reads from `stidin` and sends messages on a specific chat using a Telegram bot.

`sendtelegram` adheres to the 12 factor principles to manage settings.

It requires two environment variables:
1. SENDTELEGRAM_API_TOKEN: the Telegram API token.
2. SENDTELEGRAM_CHAT_ID: id of the chat between the bot and the target user.


## Usage

```
Usage of sendtelegram:
  -api-token string
    	Telegram API token.
  -chat-id string
    	Telegram chat ID.
```

## Examples

Sending a text file as a telegram message:
```
cat README.md | sendtelegram
```
## Changelong
* 0.1.2: Build `sendtelegram` with options to reduce filesize. 
* 0.1.1: Rewriting `sendtelegram` in GO. 

## Author

* Giovanni Stracquadanio, giovanni.stracquadanio@ed.ac.uk
