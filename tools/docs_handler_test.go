package tools

import (
	"strings"
	"testing"
)

// --- search_docs ---

func TestDocsHandlerSearchDocsMissingQuery(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "search_docs", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "query") {
		t.Errorf("expected error mentioning 'query', got %q", text)
	}
}

func TestDocsHandlerSearchDocsAuthFailure(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "search_docs", map[string]any{
		"query": "test doc",
	})
	if !isError {
		t.Fatal("expected isError=true for auth failure")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "credentials") && !strings.Contains(lower, "authenticating") && !strings.Contains(lower, "token") {
		t.Errorf("expected error about credentials/auth, got %q", text)
	}
}

// --- get_doc_content ---

func TestDocsHandlerGetDocContentMissingDocumentID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_doc_content", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "document_id") {
		t.Errorf("expected error mentioning 'document_id', got %q", text)
	}
}

func TestDocsHandlerGetDocContentAuthFailure(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_doc_content", map[string]any{
		"document_id": "doc123",
	})
	if !isError {
		t.Fatal("expected isError=true for auth failure")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "credentials") && !strings.Contains(lower, "authenticating") && !strings.Contains(lower, "token") {
		t.Errorf("expected error about credentials/auth, got %q", text)
	}
}

// --- list_docs_in_folder ---
// list_docs_in_folder has no strictly required params (folder_id defaults to "root"),
// so the first error path is auth failure.

func TestDocsHandlerListDocsInFolderAuthFailure(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "list_docs_in_folder", nil)
	if !isError {
		t.Fatal("expected isError=true for auth failure")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "credentials") && !strings.Contains(lower, "authenticating") && !strings.Contains(lower, "token") {
		t.Errorf("expected error about credentials/auth, got %q", text)
	}
}

// --- create_doc ---

func TestDocsHandlerCreateDocMissingTitle(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "create_doc", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "title") {
		t.Errorf("expected error mentioning 'title', got %q", text)
	}
}

func TestDocsHandlerCreateDocAuthFailure(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "create_doc", map[string]any{
		"title": "Test Document",
	})
	if !isError {
		t.Fatal("expected isError=true for auth failure")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "credentials") && !strings.Contains(lower, "authenticating") && !strings.Contains(lower, "token") {
		t.Errorf("expected error about credentials/auth, got %q", text)
	}
}

// --- inspect_doc_structure ---

func TestDocsHandlerInspectDocStructureMissingDocumentID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "inspect_doc_structure", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "document_id") {
		t.Errorf("expected error mentioning 'document_id', got %q", text)
	}
}

// --- debug_table_structure ---

func TestDocsHandlerDebugTableStructureMissingDocumentID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "debug_table_structure", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "document_id") {
		t.Errorf("expected error mentioning 'document_id', got %q", text)
	}
}

// --- export_doc_to_pdf ---

func TestDocsHandlerExportDocToPDFMissingDocumentID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "export_doc_to_pdf", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "document_id") {
		t.Errorf("expected error mentioning 'document_id', got %q", text)
	}
}

// --- modify_doc_text ---

func TestDocsHandlerModifyDocTextMissingDocumentID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "modify_doc_text", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "document_id") {
		t.Errorf("expected error mentioning 'document_id', got %q", text)
	}
}

func TestDocsHandlerModifyDocTextMissingTextAndFormatting(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "modify_doc_text", map[string]any{
		"document_id": "doc123",
		"start_index": 1,
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "text") && !strings.Contains(lower, "formatting") {
		t.Errorf("expected error mentioning 'text' or 'formatting', got %q", text)
	}
}

// --- find_and_replace_doc ---

func TestDocsHandlerFindAndReplaceDocMissingDocumentID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "find_and_replace_doc", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "document_id") {
		t.Errorf("expected error mentioning 'document_id', got %q", text)
	}
}

func TestDocsHandlerFindAndReplaceDocMissingFindText(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "find_and_replace_doc", map[string]any{
		"document_id": "doc123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "find_text") {
		t.Errorf("expected error mentioning 'find_text', got %q", text)
	}
}

func TestDocsHandlerFindAndReplaceDocMissingReplaceText(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "find_and_replace_doc", map[string]any{
		"document_id": "doc123",
		"find_text":   "hello",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "replace_text") {
		t.Errorf("expected error mentioning 'replace_text', got %q", text)
	}
}

// --- insert_doc_elements ---

func TestDocsHandlerInsertDocElementsMissingDocumentID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "insert_doc_elements", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "document_id") {
		t.Errorf("expected error mentioning 'document_id', got %q", text)
	}
}

func TestDocsHandlerInsertDocElementsMissingElementType(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "insert_doc_elements", map[string]any{
		"document_id": "doc123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "element_type") {
		t.Errorf("expected error mentioning 'element_type', got %q", text)
	}
}

// --- insert_doc_image ---

func TestDocsHandlerInsertDocImageMissingDocumentID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "insert_doc_image", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "document_id") {
		t.Errorf("expected error mentioning 'document_id', got %q", text)
	}
}

func TestDocsHandlerInsertDocImageMissingImageSource(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "insert_doc_image", map[string]any{
		"document_id": "doc123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "image_source") {
		t.Errorf("expected error mentioning 'image_source', got %q", text)
	}
}

// --- update_doc_headers_footers ---

func TestDocsHandlerUpdateDocHeadersFootersMissingDocumentID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "update_doc_headers_footers", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "document_id") {
		t.Errorf("expected error mentioning 'document_id', got %q", text)
	}
}

func TestDocsHandlerUpdateDocHeadersFootersMissingSectionType(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "update_doc_headers_footers", map[string]any{
		"document_id": "doc123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "section_type") {
		t.Errorf("expected error mentioning 'section_type', got %q", text)
	}
}

func TestDocsHandlerUpdateDocHeadersFootersMissingContent(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "update_doc_headers_footers", map[string]any{
		"document_id":  "doc123",
		"section_type": "header",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "content") {
		t.Errorf("expected error mentioning 'content', got %q", text)
	}
}

// --- batch_update_doc ---

func TestDocsHandlerBatchUpdateDocMissingDocumentID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "batch_update_doc", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "document_id") {
		t.Errorf("expected error mentioning 'document_id', got %q", text)
	}
}

func TestDocsHandlerBatchUpdateDocMissingOperations(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "batch_update_doc", map[string]any{
		"document_id": "doc123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "operations") {
		t.Errorf("expected error mentioning 'operations', got %q", text)
	}
}

// --- create_table_with_data ---

func TestDocsHandlerCreateTableWithDataMissingDocumentID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "create_table_with_data", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "document_id") {
		t.Errorf("expected error mentioning 'document_id', got %q", text)
	}
}

func TestDocsHandlerCreateTableWithDataMissingTableData(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "create_table_with_data", map[string]any{
		"document_id": "doc123",
		"index":       1,
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "table_data") {
		t.Errorf("expected error mentioning 'table_data', got %q", text)
	}
}

// --- update_paragraph_style ---

func TestDocsHandlerUpdateParagraphStyleMissingDocumentID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "update_paragraph_style", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "document_id") {
		t.Errorf("expected error mentioning 'document_id', got %q", text)
	}
}

func TestDocsHandlerUpdateParagraphStyleInvalidStartIndex(t *testing.T) {
	s := newToolTestServer(t)
	// start_index defaults to 0 via GetInt, which is < 1 → validation error
	text, isError := callTool(t, s, "update_paragraph_style", map[string]any{
		"document_id": "doc123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "start_index") {
		t.Errorf("expected error mentioning 'start_index', got %q", text)
	}
}

func TestDocsHandlerUpdateParagraphStyleInvalidEndIndex(t *testing.T) {
	s := newToolTestServer(t)
	// end_index defaults to 0 via GetInt, which is <= start_index (1)
	text, isError := callTool(t, s, "update_paragraph_style", map[string]any{
		"document_id": "doc123",
		"start_index": 1,
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "end_index") {
		t.Errorf("expected error mentioning 'end_index', got %q", text)
	}
}

// --- Document Comment Tools (via RegisterCommentTools) ---

func TestDocsHandlerReadDocumentCommentsMissingDocumentID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "read_document_comments", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "document_id") {
		t.Errorf("expected error mentioning 'document_id', got %q", text)
	}
}

func TestDocsHandlerCreateDocumentCommentMissingDocumentID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "create_document_comment", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "document_id") {
		t.Errorf("expected error mentioning 'document_id', got %q", text)
	}
}

func TestDocsHandlerCreateDocumentCommentMissingContent(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "create_document_comment", map[string]any{
		"document_id": "doc123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "comment_content") {
		t.Errorf("expected error mentioning 'comment_content', got %q", text)
	}
}

func TestDocsHandlerReplyToDocumentCommentMissingDocumentID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "reply_to_document_comment", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "document_id") {
		t.Errorf("expected error mentioning 'document_id', got %q", text)
	}
}

func TestDocsHandlerReplyToDocumentCommentMissingCommentID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "reply_to_document_comment", map[string]any{
		"document_id": "doc123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "comment_id") {
		t.Errorf("expected error mentioning 'comment_id', got %q", text)
	}
}

func TestDocsHandlerResolveDocumentCommentMissingDocumentID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "resolve_document_comment", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "document_id") {
		t.Errorf("expected error mentioning 'document_id', got %q", text)
	}
}
