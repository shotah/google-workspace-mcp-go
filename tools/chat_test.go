package tools

import (
	"testing"

	chat "google.golang.org/api/chat/v1"
)

func TestMessagesWithSpace(t *testing.T) {
	t.Parallel()
	msgs := []*chat.Message{
		{Name: "spaces/A/messages/1", Text: "hi"},
		{Name: "spaces/A/messages/2", Text: "yo"},
	}
	got := messagesWithSpace(msgs, "General")
	if len(got) != 2 {
		t.Fatalf("len=%d", len(got))
	}
	if got[0].spaceName != "General" || got[0].msg.Name != "spaces/A/messages/1" {
		t.Fatalf("unexpected first entry: %+v", got[0])
	}
	if got := messagesWithSpace(nil, "x"); len(got) != 0 {
		t.Fatalf("expected empty slice, got %#v", got)
	}
}
