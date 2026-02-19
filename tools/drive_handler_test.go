package tools

import (
	"strings"
	"testing"
)

// --- search_drive_files ---

func TestDriveHandlerSearchMissingQuery(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "search_drive_files", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "query") {
		t.Errorf("expected error mentioning 'query', got %q", text)
	}
}

func TestDriveHandlerSearchAuthFailure(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "search_drive_files", map[string]any{
		"query": "test document",
	})
	if !isError {
		t.Fatal("expected isError=true for auth failure")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "credentials") && !strings.Contains(lower, "authenticating") && !strings.Contains(lower, "token") {
		t.Errorf("expected error about credentials/auth, got %q", text)
	}
}

// --- get_drive_file_content ---

func TestDriveHandlerGetFileContentMissingFileID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_drive_file_content", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "file_id") {
		t.Errorf("expected error mentioning 'file_id', got %q", text)
	}
}

func TestDriveHandlerGetFileContentAuthFailure(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_drive_file_content", map[string]any{
		"file_id": "abc123",
	})
	if !isError {
		t.Fatal("expected isError=true for auth failure")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "credentials") && !strings.Contains(lower, "authenticating") && !strings.Contains(lower, "token") {
		t.Errorf("expected error about credentials/auth, got %q", text)
	}
}

// --- get_drive_file_download_url ---

func TestDriveHandlerGetFileDownloadURLMissingFileID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_drive_file_download_url", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "file_id") {
		t.Errorf("expected error mentioning 'file_id', got %q", text)
	}
}

// --- list_drive_items ---
// list_drive_items has no strictly required params (folder_id defaults to "root"),
// so the first error path is auth failure.

func TestDriveHandlerListDriveItemsAuthFailure(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "list_drive_items", nil)
	if !isError {
		t.Fatal("expected isError=true for auth failure")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "credentials") && !strings.Contains(lower, "authenticating") && !strings.Contains(lower, "token") {
		t.Errorf("expected error about credentials/auth, got %q", text)
	}
}

// --- get_drive_file_permissions ---

func TestDriveHandlerGetFilePermissionsMissingFileID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_drive_file_permissions", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "file_id") {
		t.Errorf("expected error mentioning 'file_id', got %q", text)
	}
}

// --- check_drive_file_public_access ---

func TestDriveHandlerCheckPublicAccessMissingFileName(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "check_drive_file_public_access", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "file_name") {
		t.Errorf("expected error mentioning 'file_name', got %q", text)
	}
}

// --- get_drive_shareable_link ---

func TestDriveHandlerGetShareableLinkMissingFileID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_drive_shareable_link", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "file_id") {
		t.Errorf("expected error mentioning 'file_id', got %q", text)
	}
}

// --- create_drive_file ---

func TestDriveHandlerCreateFileMissingFileName(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "create_drive_file", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "file_name") {
		t.Errorf("expected error mentioning 'file_name', got %q", text)
	}
}

func TestDriveHandlerCreateFileMissingContentAndURL(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "create_drive_file", map[string]any{
		"file_name": "test.txt",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "content") && !strings.Contains(lower, "fileurl") {
		t.Errorf("expected error mentioning 'content' or 'fileUrl', got %q", text)
	}
}

// --- import_to_google_doc ---

func TestDriveHandlerImportToGoogleDocMissingFileName(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "import_to_google_doc", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "file_name") {
		t.Errorf("expected error mentioning 'file_name', got %q", text)
	}
}

func TestDriveHandlerImportToGoogleDocMissingSource(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "import_to_google_doc", map[string]any{
		"file_name": "test.md",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "content") && !strings.Contains(lower, "file_path") && !strings.Contains(lower, "file_url") {
		t.Errorf("expected error mentioning source params, got %q", text)
	}
}

// --- update_drive_file ---

func TestDriveHandlerUpdateFileMissingFileID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "update_drive_file", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "file_id") {
		t.Errorf("expected error mentioning 'file_id', got %q", text)
	}
}

// --- copy_drive_file ---

func TestDriveHandlerCopyFileMissingFileID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "copy_drive_file", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "file_id") {
		t.Errorf("expected error mentioning 'file_id', got %q", text)
	}
}

// --- share_drive_file ---

func TestDriveHandlerShareFileMissingFileID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "share_drive_file", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "file_id") {
		t.Errorf("expected error mentioning 'file_id', got %q", text)
	}
}

func TestDriveHandlerShareFileMissingShareWith(t *testing.T) {
	s := newToolTestServer(t)
	// share_type defaults to "user", which requires share_with.
	// This param check happens before auth.
	text, isError := callTool(t, s, "share_drive_file", map[string]any{
		"file_id": "abc123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "share_with") {
		t.Errorf("expected error mentioning 'share_with', got %q", text)
	}
}

// --- batch_share_drive_file ---

func TestDriveHandlerBatchShareFileMissingFileID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "batch_share_drive_file", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "file_id") {
		t.Errorf("expected error mentioning 'file_id', got %q", text)
	}
}

func TestDriveHandlerBatchShareFileMissingRecipients(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "batch_share_drive_file", map[string]any{
		"file_id": "abc123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "recipients") {
		t.Errorf("expected error mentioning 'recipients', got %q", text)
	}
}

func TestDriveHandlerBatchShareFileEmptyRecipients(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "batch_share_drive_file", map[string]any{
		"file_id":    "abc123",
		"recipients": []any{},
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "recipients") && !strings.Contains(lower, "empty") {
		t.Errorf("expected error mentioning empty recipients, got %q", text)
	}
}

// --- update_drive_permission ---

func TestDriveHandlerUpdatePermissionMissingFileID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "update_drive_permission", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "file_id") {
		t.Errorf("expected error mentioning 'file_id', got %q", text)
	}
}

func TestDriveHandlerUpdatePermissionMissingPermissionID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "update_drive_permission", map[string]any{
		"file_id": "abc123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "permission_id") {
		t.Errorf("expected error mentioning 'permission_id', got %q", text)
	}
}

func TestDriveHandlerUpdatePermissionMissingRoleAndExpiration(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "update_drive_permission", map[string]any{
		"file_id":       "abc123",
		"permission_id": "perm456",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "role") && !strings.Contains(lower, "expiration") {
		t.Errorf("expected error mentioning 'role' or 'expiration_time', got %q", text)
	}
}

func TestDriveHandlerUpdatePermissionInvalidRole(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "update_drive_permission", map[string]any{
		"file_id":       "abc123",
		"permission_id": "perm456",
		"role":          "superadmin",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "invalid role") {
		t.Errorf("expected error mentioning 'invalid role', got %q", text)
	}
}

// --- remove_drive_permission ---

func TestDriveHandlerRemovePermissionMissingFileID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "remove_drive_permission", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "file_id") {
		t.Errorf("expected error mentioning 'file_id', got %q", text)
	}
}

func TestDriveHandlerRemovePermissionMissingPermissionID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "remove_drive_permission", map[string]any{
		"file_id": "abc123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "permission_id") {
		t.Errorf("expected error mentioning 'permission_id', got %q", text)
	}
}

// --- transfer_drive_ownership ---

func TestDriveHandlerTransferOwnershipMissingFileID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "transfer_drive_ownership", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "file_id") {
		t.Errorf("expected error mentioning 'file_id', got %q", text)
	}
}

func TestDriveHandlerTransferOwnershipMissingNewOwnerEmail(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "transfer_drive_ownership", map[string]any{
		"file_id": "abc123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "new_owner_email") {
		t.Errorf("expected error mentioning 'new_owner_email', got %q", text)
	}
}

// --- Additional validation tests ---

func TestDriveHandlerShareFileInvalidRole(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "share_drive_file", map[string]any{
		"file_id":    "abc123",
		"share_with": "user@example.com",
		"role":       "superadmin",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "invalid role") {
		t.Errorf("expected error mentioning 'invalid role', got %q", text)
	}
}

func TestDriveHandlerShareFileInvalidShareType(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "share_drive_file", map[string]any{
		"file_id":    "abc123",
		"share_with": "user@example.com",
		"share_type": "martian",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "invalid share_type") {
		t.Errorf("expected error mentioning 'invalid share_type', got %q", text)
	}
}

func TestDriveHandlerImportToGoogleDocMultipleSources(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "import_to_google_doc", map[string]any{
		"file_name": "test.md",
		"content":   "# Hello",
		"file_path": "/tmp/test.md",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "only one") && !strings.Contains(lower, "provide only") {
		t.Errorf("expected error about providing only one source, got %q", text)
	}
}

func TestDriveHandlerImportToGoogleDocUnsupportedFormat(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "import_to_google_doc", map[string]any{
		"file_name":     "test.xyz",
		"content":       "some content",
		"source_format": "xyz",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "unsupported") {
		t.Errorf("expected error mentioning 'unsupported', got %q", text)
	}
}
