package tools

// search.go has no unexported pure helper functions.
// executeSearch requires a context and makes API calls (not a pure function).
// newCustomSearchService reads environment variables and creates external service connections.
// No pure functions to unit test.
