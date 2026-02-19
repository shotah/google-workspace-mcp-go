package tools

import (
	"strings"
	"testing"
)

// --- list_contacts ---
// list_contacts has no strictly required params, so first error is auth failure.

func TestContactsHandlerListContactsAuthFailure(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "list_contacts", nil)
	if !isError {
		t.Fatal("expected isError=true for auth failure")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "credentials") && !strings.Contains(lower, "authenticating") {
		t.Errorf("expected error about credentials/auth, got %q", text)
	}
}

// --- get_contact ---

func TestContactsHandlerGetContactMissingContactID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_contact", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "contact_id") {
		t.Errorf("expected error mentioning 'contact_id', got %q", text)
	}
}

func TestContactsHandlerGetContactAuthFailure(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_contact", map[string]any{
		"contact_id": "c1234567890",
	})
	if !isError {
		t.Fatal("expected isError=true for auth failure")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "credentials") && !strings.Contains(lower, "authenticating") {
		t.Errorf("expected error about credentials/auth, got %q", text)
	}
}

// --- search_contacts ---

func TestContactsHandlerSearchContactsMissingQuery(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "search_contacts", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "query") {
		t.Errorf("expected error mentioning 'query', got %q", text)
	}
}

// --- list_contact_groups ---
// list_contact_groups has no strictly required params, so first error is auth failure.

func TestContactsHandlerListContactGroupsAuthFailure(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "list_contact_groups", nil)
	if !isError {
		t.Fatal("expected isError=true for auth failure")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "credentials") && !strings.Contains(lower, "authenticating") {
		t.Errorf("expected error about credentials/auth, got %q", text)
	}
}

// --- get_contact_group ---

func TestContactsHandlerGetContactGroupMissingGroupID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_contact_group", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "group_id") {
		t.Errorf("expected error mentioning 'group_id', got %q", text)
	}
}

// --- create_contact ---
// create_contact has no RequireString params. It calls resolveEmail (succeeds via env),
// then newPeopleService (auth failure), then buildPersonBody check.
// Auth failure comes before the buildPersonBody check.

func TestContactsHandlerCreateContactAuthFailure(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "create_contact", map[string]any{
		"given_name": "Test",
	})
	if !isError {
		t.Fatal("expected isError=true for auth failure")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "credentials") && !strings.Contains(lower, "authenticating") {
		t.Errorf("expected error about credentials/auth, got %q", text)
	}
}

func TestContactsHandlerCreateContactNoFields(t *testing.T) {
	// create_contact with no fields calls newPeopleService first (auth failure),
	// then buildPersonBody. Since auth fails first, we can only test auth here.
	// But let's test that an empty call hits auth before field validation.
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "create_contact", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	// Auth failure comes before field validation
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "credentials") && !strings.Contains(lower, "authenticating") {
		t.Errorf("expected error about credentials/auth, got %q", text)
	}
}

// --- update_contact ---

func TestContactsHandlerUpdateContactMissingContactID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "update_contact", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "contact_id") {
		t.Errorf("expected error mentioning 'contact_id', got %q", text)
	}
}

// --- delete_contact ---

func TestContactsHandlerDeleteContactMissingContactID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "delete_contact", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "contact_id") {
		t.Errorf("expected error mentioning 'contact_id', got %q", text)
	}
}

// --- batch_create_contacts ---
// contacts is checked via args type assertion, not RequireString.

func TestContactsHandlerBatchCreateContactsNoContacts(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "batch_create_contacts", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "contact") {
		t.Errorf("expected error mentioning 'contact', got %q", text)
	}
}

func TestContactsHandlerBatchCreateContactsEmptyArray(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "batch_create_contacts", map[string]any{
		"contacts": []any{},
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "contact") {
		t.Errorf("expected error mentioning 'contact', got %q", text)
	}
}

// --- batch_update_contacts ---

func TestContactsHandlerBatchUpdateContactsNoUpdates(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "batch_update_contacts", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "update") {
		t.Errorf("expected error mentioning 'update', got %q", text)
	}
}

// --- batch_delete_contacts ---

func TestContactsHandlerBatchDeleteContactsMissingContactIDs(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "batch_delete_contacts", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "contact_id") {
		t.Errorf("expected error mentioning 'contact_id', got %q", text)
	}
}

func TestContactsHandlerBatchDeleteContactsEmptyArray(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "batch_delete_contacts", map[string]any{
		"contact_ids": []any{},
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "contact") {
		t.Errorf("expected error mentioning 'contact', got %q", text)
	}
}

// --- create_contact_group ---

func TestContactsHandlerCreateContactGroupMissingName(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "create_contact_group", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "name") {
		t.Errorf("expected error mentioning 'name', got %q", text)
	}
}

// --- update_contact_group ---

func TestContactsHandlerUpdateContactGroupMissingGroupID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "update_contact_group", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "group_id") {
		t.Errorf("expected error mentioning 'group_id', got %q", text)
	}
}

func TestContactsHandlerUpdateContactGroupMissingName(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "update_contact_group", map[string]any{
		"group_id": "abc123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "name") {
		t.Errorf("expected error mentioning 'name', got %q", text)
	}
}

// --- delete_contact_group ---

func TestContactsHandlerDeleteContactGroupMissingGroupID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "delete_contact_group", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "group_id") {
		t.Errorf("expected error mentioning 'group_id', got %q", text)
	}
}

// --- modify_contact_group_members ---

func TestContactsHandlerModifyContactGroupMembersMissingGroupID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "modify_contact_group_members", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "group_id") {
		t.Errorf("expected error mentioning 'group_id', got %q", text)
	}
}

func TestContactsHandlerModifyContactGroupMembersNoAddOrRemove(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "modify_contact_group_members", map[string]any{
		"group_id": "abc123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "add_contact_ids") && !strings.Contains(lower, "remove_contact_ids") {
		t.Errorf("expected error mentioning add/remove contact IDs, got %q", text)
	}
}
