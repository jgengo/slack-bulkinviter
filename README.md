# slack-bulkinviter

Super quick and dirty Golang script to bulk invite ALL users in a slack team to a specific channel. We need to do this regulary at [HiveHelsinki](http://hive.fi) when we create a new channel that everyone should be in.

To use:
* export your slack_token into the env var : SLACK_TOKEN
* compile it : go build -o inviter main.go
* ./inviter -c <channel_id>
* sit back and let it do its work

Inspired a lot on: https://github.com/robby-dermody/slack-bulkinviter but this version take care to check if it's a public channel (channel), a private channel (group) or a public channel which has been turned to a private one (conversation) in the time.
