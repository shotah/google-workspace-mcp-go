package tools

// appscript.go has no unexported pure helper functions worth unit testing.
//
// All unexported functions are either:
//   - newScriptService / newDriveServiceForScript: service constructors requiring auth
//   - make*Handler / register*: handler factories that make API calls
//
// The only self-contained logic is generateTriggerCodeHandler, which is an
// exported handler (not a pure function) that generates trigger code snippets.
// It will be covered by protocol-level handler tests (US-013) and does not
// warrant a separate pure-function test here.
