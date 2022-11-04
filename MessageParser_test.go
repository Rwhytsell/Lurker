package main

import "testing"

const messageToParse = "@badge-info=subscriber/27;badges=subscriber/24,premium/1;client-nonce=REDACTED;color=#9ACD32;display-name=REDACTED;emotes=;first-msg=0;flags=;id=REDACTED;mod=0;reply-parent-display-name=REDACTED;reply-parent-msg-body=REDACTED;reply-parent-msg-id=REDACTED;reply-parent-user-id=REDACTED;reply-parent-user-login=REDACTED;returning-chatter=0;room-id=REDACTED;subscriber=1;tmi-sent-ts=REDACTED;turbo=0;user-id=REDACTED;user-type= :REDACTED!REDACTED@REDACTED.tmi.twitch.tv PRIVMSG #chessbrah :@REDACTED he streamed yesterday; test"

func TestParseMessage(t *testing.T) {
	message := ParseMessage(messageToParse)
	if message.Channel != "chessbrah" {
		t.Fail()
	}
	if message.Username != "REDACTED" {
		t.Fail()
	}
	if message.UserId != "REDACTED" {
		t.Fail()
	}
	if message.IsSub != true {
		t.Fail()
	}
	if message.IsMod != false {
		t.Fail()
	}
	if message.IsTurbo != false {
		t.Fail()
	}
	if message.MessageText != "@REDACTED he streamed yesterday; test" {
		t.Fail()
	}
}
