package tools

import (
	"strings"
	"testing"
)

// --- create_form ---

func TestFormsHandlerCreateFormMissingTitle(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "create_form", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "title") {
		t.Errorf("expected error mentioning 'title', got %q", text)
	}
}

func TestFormsHandlerCreateFormAuthFailure(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "create_form", map[string]any{
		"title": "Test Form",
	})
	if !isError {
		t.Fatal("expected isError=true for auth failure")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "credentials") && !strings.Contains(lower, "authenticating") && !strings.Contains(lower, "authentication") && !strings.Contains(lower, "token") {
		t.Errorf("expected error about credentials/auth, got %q", text)
	}
}

// --- get_form ---

func TestFormsHandlerGetFormMissingFormID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_form", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "form_id") {
		t.Errorf("expected error mentioning 'form_id', got %q", text)
	}
}

func TestFormsHandlerGetFormAuthFailure(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_form", map[string]any{
		"form_id": "form123",
	})
	if !isError {
		t.Fatal("expected isError=true for auth failure")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "credentials") && !strings.Contains(lower, "authenticating") && !strings.Contains(lower, "authentication") && !strings.Contains(lower, "token") {
		t.Errorf("expected error about credentials/auth, got %q", text)
	}
}

// --- set_publish_settings ---

func TestFormsHandlerSetPublishSettingsMissingFormID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "set_publish_settings", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "form_id") {
		t.Errorf("expected error mentioning 'form_id', got %q", text)
	}
}

// --- get_form_response ---

func TestFormsHandlerGetFormResponseMissingFormID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_form_response", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "form_id") {
		t.Errorf("expected error mentioning 'form_id', got %q", text)
	}
}

func TestFormsHandlerGetFormResponseMissingResponseID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_form_response", map[string]any{
		"form_id": "form123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "response_id") {
		t.Errorf("expected error mentioning 'response_id', got %q", text)
	}
}

// --- list_form_responses ---

func TestFormsHandlerListFormResponsesMissingFormID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "list_form_responses", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "form_id") {
		t.Errorf("expected error mentioning 'form_id', got %q", text)
	}
}

// --- batch_update_form ---

func TestFormsHandlerBatchUpdateFormMissingFormID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "batch_update_form", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "form_id") {
		t.Errorf("expected error mentioning 'form_id', got %q", text)
	}
}

func TestFormsHandlerBatchUpdateFormMissingRequests(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "batch_update_form", map[string]any{
		"form_id": "form123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "requests") {
		t.Errorf("expected error mentioning 'requests', got %q", text)
	}
}

func TestFormsHandlerBatchUpdateFormRequestsNotArray(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "batch_update_form", map[string]any{
		"form_id":  "form123",
		"requests": "not-an-array",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "requests") && !strings.Contains(lower, "array") {
		t.Errorf("expected error mentioning 'requests' or 'array', got %q", text)
	}
}
