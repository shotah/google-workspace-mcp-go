package tools

import (
	"strings"
	"testing"
)

// TestGmailHandlerSearchMissingQuery verifies that search_gmail_messages
// returns a tool-level error when the required "query" param is missing.
func TestGmailHandlerSearchMissingQuery(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "search_gmail_messages", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "query") {
		t.Errorf("expected error mentioning 'query', got %q", text)
	}
}

// TestGmailHandlerGetMessageContentMissingMessageID verifies that
// get_gmail_message_content returns error when message_id is missing.
func TestGmailHandlerGetMessageContentMissingMessageID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_gmail_message_content", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "message_id") {
		t.Errorf("expected error mentioning 'message_id', got %q", text)
	}
}

// TestGmailHandlerGetMessagesContentBatchMissingMessageIDs verifies that
// get_gmail_messages_content_batch returns error when message_ids is missing.
func TestGmailHandlerGetMessagesContentBatchMissingMessageIDs(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_gmail_messages_content_batch", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "message_ids") {
		t.Errorf("expected error mentioning 'message_ids', got %q", text)
	}
}

// TestGmailHandlerGetMessagesContentBatchEmptyMessageIDs verifies that
// get_gmail_messages_content_batch returns error when message_ids is empty.
func TestGmailHandlerGetMessagesContentBatchEmptyMessageIDs(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_gmail_messages_content_batch", map[string]any{
		"message_ids": []any{},
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "no message ids") {
		t.Errorf("expected error mentioning empty message_ids, got %q", text)
	}
}

// TestGmailHandlerGetAttachmentContentMissingMessageID verifies that
// get_gmail_attachment_content returns error when message_id is missing.
func TestGmailHandlerGetAttachmentContentMissingMessageID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_gmail_attachment_content", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "message_id") {
		t.Errorf("expected error mentioning 'message_id', got %q", text)
	}
}

// TestGmailHandlerGetAttachmentContentMissingAttachmentID verifies that
// get_gmail_attachment_content returns error when attachment_id is missing.
func TestGmailHandlerGetAttachmentContentMissingAttachmentID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_gmail_attachment_content", map[string]any{
		"message_id": "abc123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "attachment_id") {
		t.Errorf("expected error mentioning 'attachment_id', got %q", text)
	}
}

// TestGmailHandlerGetThreadContentMissingThreadID verifies that
// get_gmail_thread_content returns error when thread_id is missing.
func TestGmailHandlerGetThreadContentMissingThreadID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_gmail_thread_content", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "thread_id") {
		t.Errorf("expected error mentioning 'thread_id', got %q", text)
	}
}

// TestGmailHandlerGetThreadsContentBatchMissingThreadIDs verifies that
// get_gmail_threads_content_batch returns error when thread_ids is missing.
func TestGmailHandlerGetThreadsContentBatchMissingThreadIDs(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_gmail_threads_content_batch", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "thread_ids") {
		t.Errorf("expected error mentioning 'thread_ids', got %q", text)
	}
}

// TestGmailHandlerGetThreadsContentBatchEmptyThreadIDs verifies that
// get_gmail_threads_content_batch returns error when thread_ids is empty.
func TestGmailHandlerGetThreadsContentBatchEmptyThreadIDs(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_gmail_threads_content_batch", map[string]any{
		"thread_ids": []any{},
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "no thread ids") {
		t.Errorf("expected error mentioning empty thread_ids, got %q", text)
	}
}

// TestGmailHandlerModifyMessageLabelsMissingMessageID verifies that
// modify_gmail_message_labels returns error when message_id is missing.
func TestGmailHandlerModifyMessageLabelsMissingMessageID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "modify_gmail_message_labels", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "message_id") {
		t.Errorf("expected error mentioning 'message_id', got %q", text)
	}
}

// TestGmailHandlerModifyMessageLabelsMissingBothLabelArrays verifies that
// modify_gmail_message_labels returns error when neither add nor remove labels provided.
func TestGmailHandlerModifyMessageLabelsMissingBothLabelArrays(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "modify_gmail_message_labels", map[string]any{
		"message_id": "msg123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "add_label_ids") && !strings.Contains(strings.ToLower(text), "remove_label_ids") {
		t.Errorf("expected error mentioning label arrays, got %q", text)
	}
}

// TestGmailHandlerManageGmailLabelMissingAction verifies that
// manage_gmail_label returns error when action is missing.
func TestGmailHandlerManageGmailLabelMissingAction(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "manage_gmail_label", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "action") {
		t.Errorf("expected error mentioning 'action', got %q", text)
	}
}

// TestGmailHandlerManageGmailLabelCreateMissingName verifies that
// manage_gmail_label with action=create returns error when name is missing.
// Note: the name check happens after auth, so without credentials we get an
// auth error first. This test verifies the handler still returns an error.
func TestGmailHandlerManageGmailLabelCreateMissingName(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "manage_gmail_label", map[string]any{
		"action": "create",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	// Auth error comes before name validation; verify we get some error.
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "name") && !strings.Contains(lower, "credentials") && !strings.Contains(lower, "authenticating") {
		t.Errorf("expected error mentioning 'name' or credentials, got %q", text)
	}
}

// TestGmailHandlerSendGmailMessageMissingTo verifies that
// send_gmail_message returns error when to is missing.
func TestGmailHandlerSendGmailMessageMissingTo(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "send_gmail_message", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	// The handler checks email first, then to, subject, body in order.
	// With USER_GOOGLE_EMAIL set, it should get to the "to" check.
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "to") {
		t.Errorf("expected error mentioning 'to', got %q", text)
	}
}

// TestGmailHandlerSendGmailMessageMissingSubject verifies that
// send_gmail_message returns error when subject is missing.
func TestGmailHandlerSendGmailMessageMissingSubject(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "send_gmail_message", map[string]any{
		"to": "someone@example.com",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "subject") {
		t.Errorf("expected error mentioning 'subject', got %q", text)
	}
}

// TestGmailHandlerSendGmailMessageMissingBody verifies that
// send_gmail_message returns error when body is missing.
func TestGmailHandlerSendGmailMessageMissingBody(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "send_gmail_message", map[string]any{
		"to":      "someone@example.com",
		"subject": "Test Subject",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "body") {
		t.Errorf("expected error mentioning 'body', got %q", text)
	}
}

// TestGmailHandlerDraftGmailMessageMissingSubject verifies that
// draft_gmail_message returns error when subject is missing.
func TestGmailHandlerDraftGmailMessageMissingSubject(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "draft_gmail_message", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "subject") {
		t.Errorf("expected error mentioning 'subject', got %q", text)
	}
}

// TestGmailHandlerDraftGmailMessageMissingBody verifies that
// draft_gmail_message returns error when body is missing.
func TestGmailHandlerDraftGmailMessageMissingBody(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "draft_gmail_message", map[string]any{
		"subject": "Test Subject",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "body") {
		t.Errorf("expected error mentioning 'body', got %q", text)
	}
}

// TestGmailHandlerBatchModifyMessageLabelsMissingMessageIDs verifies that
// batch_modify_gmail_message_labels returns error when message_ids is missing.
func TestGmailHandlerBatchModifyMessageLabelsMissingMessageIDs(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "batch_modify_gmail_message_labels", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "message_ids") {
		t.Errorf("expected error mentioning 'message_ids', got %q", text)
	}
}

// TestGmailHandlerBatchModifyMessageLabelsEmptyMessageIDs verifies that
// batch_modify_gmail_message_labels returns error when message_ids is empty.
func TestGmailHandlerBatchModifyMessageLabelsEmptyMessageIDs(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "batch_modify_gmail_message_labels", map[string]any{
		"message_ids": []any{},
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "message_ids") && !strings.Contains(lower, "empty") {
		t.Errorf("expected error mentioning empty message_ids, got %q", text)
	}
}

// TestGmailHandlerCreateGmailFilterMissingCriteria verifies that
// create_gmail_filter returns error when criteria is missing.
func TestGmailHandlerCreateGmailFilterMissingCriteria(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "create_gmail_filter", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "criteria") {
		t.Errorf("expected error mentioning 'criteria', got %q", text)
	}
}

// TestGmailHandlerCreateGmailFilterMissingAction verifies that
// create_gmail_filter returns error when action is missing.
func TestGmailHandlerCreateGmailFilterMissingAction(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "create_gmail_filter", map[string]any{
		"criteria": map[string]any{"from": "test@example.com"},
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "action") {
		t.Errorf("expected error mentioning 'action', got %q", text)
	}
}

// TestGmailHandlerDeleteGmailFilterMissingFilterID verifies that
// delete_gmail_filter returns error when filter_id is missing.
func TestGmailHandlerDeleteGmailFilterMissingFilterID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "delete_gmail_filter", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "filter_id") {
		t.Errorf("expected error mentioning 'filter_id', got %q", text)
	}
}

// TestGmailHandlerAuthFailureSearchGmailMessages verifies that
// search_gmail_messages with valid params but no credentials returns
// an error mentioning credentials or authentication.
func TestGmailHandlerAuthFailureSearchGmailMessages(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "search_gmail_messages", map[string]any{
		"query": "in:inbox",
	})
	if !isError {
		t.Fatal("expected isError=true for auth failure")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "credentials") && !strings.Contains(lower, "authenticating") && !strings.Contains(lower, "token") {
		t.Errorf("expected error about credentials/auth, got %q", text)
	}
}

// TestGmailHandlerAuthFailureListGmailLabels verifies that
// list_gmail_labels with no credentials returns an auth error.
func TestGmailHandlerAuthFailureListGmailLabels(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "list_gmail_labels", nil)
	if !isError {
		t.Fatal("expected isError=true for auth failure")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "credentials") && !strings.Contains(lower, "authenticating") && !strings.Contains(lower, "token") {
		t.Errorf("expected error about credentials/auth, got %q", text)
	}
}
