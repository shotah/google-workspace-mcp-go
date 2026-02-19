package tools

import (
	"strings"
	"testing"
)

// --- list_task_lists ---
// list_task_lists has no strictly required params (email resolved via env),
// so the first error path is auth failure.

func TestTasksHandlerListTaskListsAuthFailure(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "list_task_lists", nil)
	if !isError {
		t.Fatal("expected isError=true for auth failure")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "credentials") && !strings.Contains(lower, "authenticating") {
		t.Errorf("expected error about credentials/auth, got %q", text)
	}
}

// --- get_task_list ---

func TestTasksHandlerGetTaskListMissingTaskListID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_task_list", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "task_list_id") {
		t.Errorf("expected error mentioning 'task_list_id', got %q", text)
	}
}

func TestTasksHandlerGetTaskListAuthFailure(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_task_list", map[string]any{
		"task_list_id": "abc123",
	})
	if !isError {
		t.Fatal("expected isError=true for auth failure")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "credentials") && !strings.Contains(lower, "authenticating") {
		t.Errorf("expected error about credentials/auth, got %q", text)
	}
}

// --- create_task_list ---

func TestTasksHandlerCreateTaskListMissingTitle(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "create_task_list", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "title") {
		t.Errorf("expected error mentioning 'title', got %q", text)
	}
}

// --- update_task_list ---

func TestTasksHandlerUpdateTaskListMissingTaskListID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "update_task_list", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "task_list_id") {
		t.Errorf("expected error mentioning 'task_list_id', got %q", text)
	}
}

func TestTasksHandlerUpdateTaskListMissingTitle(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "update_task_list", map[string]any{
		"task_list_id": "abc123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "title") {
		t.Errorf("expected error mentioning 'title', got %q", text)
	}
}

// --- delete_task_list ---

func TestTasksHandlerDeleteTaskListMissingTaskListID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "delete_task_list", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "task_list_id") {
		t.Errorf("expected error mentioning 'task_list_id', got %q", text)
	}
}

// --- list_tasks ---

func TestTasksHandlerListTasksMissingTaskListID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "list_tasks", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "task_list_id") {
		t.Errorf("expected error mentioning 'task_list_id', got %q", text)
	}
}

// --- get_task ---

func TestTasksHandlerGetTaskMissingTaskListID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_task", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "task_list_id") {
		t.Errorf("expected error mentioning 'task_list_id', got %q", text)
	}
}

func TestTasksHandlerGetTaskMissingTaskID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "get_task", map[string]any{
		"task_list_id": "abc123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "task_id") {
		t.Errorf("expected error mentioning 'task_id', got %q", text)
	}
}

// --- create_task ---

func TestTasksHandlerCreateTaskMissingTaskListID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "create_task", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "task_list_id") {
		t.Errorf("expected error mentioning 'task_list_id', got %q", text)
	}
}

func TestTasksHandlerCreateTaskMissingTitle(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "create_task", map[string]any{
		"task_list_id": "abc123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "title") {
		t.Errorf("expected error mentioning 'title', got %q", text)
	}
}

// --- update_task ---

func TestTasksHandlerUpdateTaskMissingTaskListID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "update_task", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "task_list_id") {
		t.Errorf("expected error mentioning 'task_list_id', got %q", text)
	}
}

func TestTasksHandlerUpdateTaskMissingTaskID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "update_task", map[string]any{
		"task_list_id": "abc123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "task_id") {
		t.Errorf("expected error mentioning 'task_id', got %q", text)
	}
}

// --- delete_task ---

func TestTasksHandlerDeleteTaskMissingTaskListID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "delete_task", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "task_list_id") {
		t.Errorf("expected error mentioning 'task_list_id', got %q", text)
	}
}

func TestTasksHandlerDeleteTaskMissingTaskID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "delete_task", map[string]any{
		"task_list_id": "abc123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "task_id") {
		t.Errorf("expected error mentioning 'task_id', got %q", text)
	}
}

// --- move_task ---

func TestTasksHandlerMoveTaskMissingTaskListID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "move_task", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "task_list_id") {
		t.Errorf("expected error mentioning 'task_list_id', got %q", text)
	}
}

func TestTasksHandlerMoveTaskMissingTaskID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "move_task", map[string]any{
		"task_list_id": "abc123",
	})
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "task_id") {
		t.Errorf("expected error mentioning 'task_id', got %q", text)
	}
}

// --- clear_completed_tasks ---

func TestTasksHandlerClearCompletedTasksMissingTaskListID(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "clear_completed_tasks", nil)
	if !isError {
		t.Fatal("expected isError=true")
	}
	if !strings.Contains(strings.ToLower(text), "task_list_id") {
		t.Errorf("expected error mentioning 'task_list_id', got %q", text)
	}
}

func TestTasksHandlerClearCompletedTasksAuthFailure(t *testing.T) {
	s := newToolTestServer(t)
	text, isError := callTool(t, s, "clear_completed_tasks", map[string]any{
		"task_list_id": "abc123",
	})
	if !isError {
		t.Fatal("expected isError=true for auth failure")
	}
	lower := strings.ToLower(text)
	if !strings.Contains(lower, "credentials") && !strings.Contains(lower, "authenticating") {
		t.Errorf("expected error about credentials/auth, got %q", text)
	}
}
