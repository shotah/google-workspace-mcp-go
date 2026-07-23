# Google Workspace MCP (Go)

<p align="center">
  <a href="https://github.com/shotah/google-workspace-mcp-go/actions/workflows/ci.yml"><img src="https://github.com/shotah/google-workspace-mcp-go/actions/workflows/ci.yml/badge.svg" alt="CI"></a>
  <a href="https://github.com/shotah/google-workspace-mcp-go/actions/workflows/release.yml"><img src="https://github.com/shotah/google-workspace-mcp-go/actions/workflows/release.yml/badge.svg" alt="Release"></a>
  <a href="https://github.com/shotah/google-workspace-mcp-go/actions/workflows/ci.yml"><img src="https://github.com/shotah/google-workspace-mcp-go/raw/gh-pages/badges/coverage.svg" alt="Coverage"></a>
  <a href="https://pkg.go.dev/github.com/shotah/google-workspace-mcp-go"><img src="https://pkg.go.dev/badge/github.com/shotah/google-workspace-mcp-go.svg" alt="Go Reference"></a>
  <img src="https://img.shields.io/github/go-mod/go-version/shotah/google-workspace-mcp-go" alt="Go version">
  <a href="LICENSE"><img src="https://img.shields.io/github/license/shotah/google-workspace-mcp-go" alt="License"></a>
</p>

<p align="center">
  <strong>Give Claude, Cursor, and other MCP clients real access to your Google Workspace.</strong><br>
  Gmail, Drive, Calendar, Docs, Sheets, and more — one small binary, no Python runtime.
</p>

**137 tools · 12 services · single binary · OAuth that just works**

Drop it into your MCP config and ask your agent to search mail, clean up calendar duplicates, draft Docs, or update Sheets — with a permission surface you control (`read` / `edit` / `complete`).

| Service | What agents can do |
| ------- | ------------------ |
| **Gmail** | Search, read, send, labels, filters |
| **Drive** | Search, read, create, share |
| **Calendar** | List, create, modify, delete events |
| **Docs / Sheets / Slides** | Read and edit Workspace files |
| **Tasks · Contacts · Chat · Forms · Apps Script · Search** | Day-to-day Workspace automation |

Built for **local, single-user** AI tool use. Need multi-user OAuth 2.1 or an HTTP server? Use the original [Python server](https://github.com/taylorwilsdon/google_workspace_mcp) — same tool surface, different deployment model.

## Why this one

- **Zero runtime** — download a binary (or `go install`) and run; no venv, no `uv`, no dependency churn
- **Same tool catalog** — 137 tools aligned with the Python reference server
- **Agent-friendly filters** — trim context with `--tools` / `--tool-tier`, and constrain writes with `--capability`
- **Local-first auth** — Desktop OAuth, tokens stored under `~/.google_workspace_mcp/credentials/`
- **Works where you already work** — Claude Code, Cursor, and any stdio MCP client

## Quick start

### 1. Google Cloud OAuth

Reuse credentials from the [Python server](https://github.com/taylorwilsdon/google_workspace_mcp) if you already have them. Otherwise:

1. Open [Google Cloud Console](https://console.cloud.google.com/)
2. Create or select a project → **APIs & Services → OAuth consent screen**
3. **Credentials → Create Credentials → OAuth Client ID → Desktop Application**
4. Copy the **Client ID** and **Client Secret**
5. Enable only the APIs you need:

<details>
<summary><strong>Enable APIs</strong> (click to expand)</summary>

- [Gmail API](https://console.cloud.google.com/flows/enableapi?apiid=gmail.googleapis.com)
- [Google Drive API](https://console.cloud.google.com/flows/enableapi?apiid=drive.googleapis.com)
- [Google Calendar API](https://console.cloud.google.com/flows/enableapi?apiid=calendar-json.googleapis.com)
- [Google Docs API](https://console.cloud.google.com/flows/enableapi?apiid=docs.googleapis.com)
- [Google Sheets API](https://console.cloud.google.com/flows/enableapi?apiid=sheets.googleapis.com)
- [Google Slides API](https://console.cloud.google.com/flows/enableapi?apiid=slides.googleapis.com)
- [Google Forms API](https://console.cloud.google.com/flows/enableapi?apiid=forms.googleapis.com)
- [Google Tasks API](https://console.cloud.google.com/flows/enableapi?apiid=tasks.googleapis.com)
- [Google Chat API](https://console.cloud.google.com/flows/enableapi?apiid=chat.googleapis.com)
- [People API (Contacts)](https://console.cloud.google.com/flows/enableapi?apiid=people.googleapis.com)
- [Apps Script API](https://console.cloud.google.com/flows/enableapi?apiid=script.googleapis.com)
- [Custom Search API](https://console.cloud.google.com/flows/enableapi?apiid=customsearch.googleapis.com) *(optional)*

</details>

### 2. Install

**Pre-built binary** (no Go required) — grab the archive for your platform from [Releases](https://github.com/shotah/google-workspace-mcp-go/releases):

| Platform | File |
| --- | --- |
| Linux x86_64 | `google-workspace-mcp-go_*_linux_amd64.tar.gz` |
| Linux ARM64 | `google-workspace-mcp-go_*_linux_arm64.tar.gz` |
| macOS Apple Silicon | `google-workspace-mcp-go_*_darwin_arm64.tar.gz` |
| macOS Intel | `google-workspace-mcp-go_*_darwin_amd64.tar.gz` |
| Windows x86_64 | `google-workspace-mcp-go_*_windows_amd64.zip` |

```bash
tar xzf google-workspace-mcp-go_*_linux_amd64.tar.gz
chmod +x google-workspace-mcp-go
mv google-workspace-mcp-go ~/.local/bin/
```

**Or with Go** (1.26+):

```bash
go install github.com/shotah/google-workspace-mcp-go@latest
```

### 3. Environment

```bash
export GOOGLE_OAUTH_CLIENT_ID="your-client-id.apps.googleusercontent.com"
export GOOGLE_OAUTH_CLIENT_SECRET="your-client-secret"
export USER_GOOGLE_EMAIL="you@gmail.com"  # optional but recommended
```

### 4. MCP client config

**Recommended** for local agents (small tool surface + everyday edit permissions):

```json
{
  "mcpServers": {
    "google-workspace": {
      "command": "google-workspace-mcp-go",
      "args": ["--tools", "gmail drive calendar", "--tool-tier", "core", "--capability", "edit"],
      "env": {
        "GOOGLE_OAUTH_CLIENT_ID": "your-client-id.apps.googleusercontent.com",
        "GOOGLE_OAUTH_CLIENT_SECRET": "your-client-secret",
        "USER_GOOGLE_EMAIL": "you@gmail.com"
      }
    }
  }
}
```

If those env vars are already exported in the shell that launches your MCP client, you can omit the `env` block:

```json
{
  "mcpServers": {
    "google-workspace": {
      "command": "google-workspace-mcp-go",
      "args": ["--tools", "gmail drive calendar", "--tool-tier", "core", "--capability", "edit"]
    }
  }
}
```

### 5. Authenticate once

Ask your assistant to call `start_google_auth` with your email. A browser opens, you approve, tokens land in `~/.google_workspace_mcp/credentials/`, and later runs refresh automatically.

> `start_google_auth` is registered with the `gmail` service. If your `--tools` list omits gmail, add it temporarily for first-time auth.

## Configuration

### CLI flags

| Flag | Description | Default |
| --- | --- | --- |
| `--tools` | Services to enable (e.g. `gmail drive calendar`) | all |
| `--tool-tier` | Depth: `core`, `extended`, or `complete` | `complete` |
| `--capability` | Permissions: `read`, `edit`, or `complete` | `complete` |
| `--read-only` | Shorthand for `--capability read` | `false` |

### Tool tiers (how many tools)

| Tier | Description | Count |
| --- | --- | --- |
| `core` | Everyday read/write per service | 45 |
| `extended` | Core + management tools | 91 |
| `complete` | Everything | 137 |

### Capabilities (what agents may do)

| Capability | Description | Count |
| --- | --- | --- |
| `read` | Read-only (same as `--read-only`) | 59 |
| `edit` | Everyday create/modify/delete; blocks high-impact ops | 131 |
| `complete` | Full surface including ownership transfer & bulk deletes | 137 |

Withheld under `edit`: `transfer_drive_ownership`, `batch_delete_contacts`, `delete_task_list`, `delete_contact_group`, `delete_script_project`, `clear_completed_tasks`.

### Environment variables

| Variable | Required | Description |
| --- | --- | --- |
| `GOOGLE_OAUTH_CLIENT_ID` | Yes | OAuth 2.0 Client ID |
| `GOOGLE_OAUTH_CLIENT_SECRET` | Yes | OAuth 2.0 Client Secret |
| `USER_GOOGLE_EMAIL` | No | Default account email |
| `WORKSPACE_MCP_CREDENTIALS_DIR` | No | Override credential directory |
| `GOOGLE_PSE_API_KEY` | No | Programmable Search Engine key |
| `GOOGLE_PSE_ENGINE_ID` | No | Programmable Search Engine ID |

### Available services

`gmail` `drive` `calendar` `docs` `sheets` `slides` `forms` `tasks` `chat` `contacts` `search` `appscript`

## When to use Go vs Python

| | This repo (Go) | [Python original](https://github.com/taylorwilsdon/google_workspace_mcp) |
| --- | --- | --- |
| Best for | Local Claude Code / Cursor / stdio MCP | Hosted or multi-user deployments |
| Install | Single binary | Python 3.10+ + deps |
| Tools | 137 | 137 |
| Transport | stdio | stdio + streamable HTTP |
| Auth | OAuth 2.0 (desktop) | OAuth 2.0 + OAuth 2.1 |
| Multi-user | No | Yes (sessions, Valkey, etc.) |

Same credentials work in both.

## Tools

<details>
<summary><strong>Full tool reference (137 tools across 12 services)</strong></summary>

### Gmail (15 tools)

| Tool                                | Tier     | Description                             |
| ----------------------------------- | -------- | --------------------------------------- |
| `search_gmail_messages`             | core     | Search messages with Gmail query syntax |
| `get_gmail_message_content`         | core     | Get full message content                |
| `get_gmail_messages_content_batch`  | core     | Batch retrieve up to 25 messages        |
| `send_gmail_message`                | core     | Send email with optional attachments    |
| `get_gmail_attachment_content`      | extended | Download attachment content             |
| `get_gmail_thread_content`          | extended | Get full conversation thread            |
| `modify_gmail_message_labels`       | extended | Add/remove labels from a message        |
| `list_gmail_labels`                 | extended | List all labels                         |
| `manage_gmail_label`                | extended | Create, update, or delete labels        |
| `draft_gmail_message`               | extended | Create draft email                      |
| `list_gmail_filters`                | extended | List mail filters                       |
| `create_gmail_filter`               | extended | Create new mail filter                  |
| `delete_gmail_filter`               | extended | Delete mail filter                      |
| `get_gmail_threads_content_batch`   | complete | Batch retrieve up to 25 threads         |
| `batch_modify_gmail_message_labels` | complete | Batch label operations                  |

**System tool** (registered with gmail, but not Gmail-specific):

| Tool                | Tier     | Description                       |
| ------------------- | -------- | --------------------------------- |
| `start_google_auth` | complete | Trigger OAuth authentication flow |

### Google Drive (16 tools)

| Tool                             | Tier     | Description                          |
| -------------------------------- | -------- | ------------------------------------ |
| `search_drive_files`             | core     | Search files with Drive query syntax |
| `get_drive_file_content`         | core     | Download file content                |
| `get_drive_file_download_url`    | core     | Get download URL                     |
| `create_drive_file`              | core     | Create new file                      |
| `import_to_google_doc`           | core     | Import file as Google Doc            |
| `share_drive_file`               | core     | Share file with users                |
| `get_drive_shareable_link`       | core     | Generate shareable link              |
| `list_drive_items`               | extended | List files in folder                 |
| `copy_drive_file`                | extended | Duplicate file                       |
| `update_drive_file`              | extended | Update file metadata/content         |
| `update_drive_permission`        | extended | Modify sharing permissions           |
| `remove_drive_permission`        | extended | Revoke access                        |
| `transfer_drive_ownership`       | extended | Transfer file ownership              |
| `batch_share_drive_file`         | extended | Batch sharing                        |
| `get_drive_file_permissions`     | complete | List all permissions                 |
| `check_drive_file_public_access` | complete | Check public sharing status          |

### Google Calendar (6 tools)

| Tool             | Tier     | Description                |
| ---------------- | -------- | -------------------------- |
| `list_calendars` | core     | List user's calendars      |
| `get_events`     | core     | Get events with time range |
| `create_event`   | core     | Create calendar event      |
| `modify_event`   | core     | Update event details       |
| `delete_event`   | core     | Delete event               |
| `query_freebusy` | extended | Check availability         |

### Google Docs (19 tools)

| Tool                         | Tier     | Description                 |
| ---------------------------- | -------- | --------------------------- |
| `get_doc_content`            | core     | Get document text content   |
| `create_doc`                 | core     | Create new document         |
| `modify_doc_text`            | core     | Edit document text          |
| `export_doc_to_pdf`          | extended | Export as PDF               |
| `search_docs`                | extended | Search documents            |
| `find_and_replace_doc`       | extended | Find and replace text       |
| `list_docs_in_folder`        | extended | List docs in Drive folder   |
| `insert_doc_elements`        | extended | Insert formatted elements   |
| `update_paragraph_style`     | extended | Change paragraph formatting |
| `insert_doc_image`           | complete | Insert image                |
| `update_doc_headers_footers` | complete | Edit headers/footers        |
| `batch_update_doc`           | complete | Batch document operations   |
| `inspect_doc_structure`      | complete | Analyze document structure  |
| `create_table_with_data`     | complete | Create and populate table   |
| `debug_table_structure`      | complete | Inspect table layout        |
| `read_document_comments`     | complete | Read all comments           |
| `create_document_comment`    | complete | Add comment                 |
| `reply_to_document_comment`  | complete | Reply to comment            |
| `resolve_document_comment`   | complete | Resolve comment             |

### Google Sheets (14 tools)

| Tool                            | Tier     | Description                  |
| ------------------------------- | -------- | ---------------------------- |
| `create_spreadsheet`            | core     | Create new spreadsheet       |
| `read_sheet_values`             | core     | Read cell values             |
| `modify_sheet_values`           | core     | Write/update cells           |
| `list_spreadsheets`             | extended | List user's spreadsheets     |
| `get_spreadsheet_info`          | extended | Get spreadsheet metadata     |
| `create_sheet`                  | complete | Add worksheet tab            |
| `format_sheet_range`            | complete | Format cell ranges           |
| `add_conditional_formatting`    | complete | Add conditional format rules |
| `update_conditional_formatting` | complete | Modify format rules          |
| `delete_conditional_formatting` | complete | Remove format rules          |
| `read_spreadsheet_comments`     | complete | Read all comments            |
| `create_spreadsheet_comment`    | complete | Add comment                  |
| `reply_to_spreadsheet_comment`  | complete | Reply to comment             |
| `resolve_spreadsheet_comment`   | complete | Resolve comment              |

### Google Slides (9 tools)

| Tool                            | Tier     | Description              |
| ------------------------------- | -------- | ------------------------ |
| `create_presentation`           | core     | Create presentation      |
| `get_presentation`              | core     | Get presentation content |
| `batch_update_presentation`     | extended | Batch slide operations   |
| `get_page`                      | extended | Get individual slide     |
| `get_page_thumbnail`            | extended | Get slide thumbnail      |
| `read_presentation_comments`    | complete | Read all comments        |
| `create_presentation_comment`   | complete | Add comment              |
| `reply_to_presentation_comment` | complete | Reply to comment         |
| `resolve_presentation_comment`  | complete | Resolve comment          |

### Google Forms (6 tools)

| Tool                   | Tier     | Description             |
| ---------------------- | -------- | ----------------------- |
| `create_form`          | core     | Create new form         |
| `get_form`             | core     | Get form details        |
| `list_form_responses`  | extended | List all responses      |
| `set_publish_settings` | complete | Configure publishing    |
| `get_form_response`    | complete | Get individual response |
| `batch_update_form`    | complete | Batch form updates      |

### Google Tasks (12 tools)

| Tool                    | Tier     | Description             |
| ----------------------- | -------- | ----------------------- |
| `get_task`              | core     | Get task details        |
| `list_tasks`            | core     | List tasks              |
| `create_task`           | core     | Create task             |
| `update_task`           | core     | Update task             |
| `delete_task`           | extended | Delete task             |
| `list_task_lists`       | complete | List task lists         |
| `get_task_list`         | complete | Get task list details   |
| `create_task_list`      | complete | Create task list        |
| `update_task_list`      | complete | Update task list        |
| `delete_task_list`      | complete | Delete task list        |
| `move_task`             | complete | Move task between lists |
| `clear_completed_tasks` | complete | Clear completed tasks   |

### Google Chat (4 tools)

| Tool              | Tier     | Description       |
| ----------------- | -------- | ----------------- |
| `send_message`    | core     | Send chat message |
| `get_messages`    | core     | Get messages      |
| `search_messages` | core     | Search messages   |
| `list_spaces`     | extended | List spaces/DMs   |

### Google Contacts (15 tools)

| Tool                           | Tier     | Description          |
| ------------------------------ | -------- | -------------------- |
| `search_contacts`              | core     | Search contacts      |
| `get_contact`                  | core     | Get contact details  |
| `list_contacts`                | core     | List contacts        |
| `create_contact`               | core     | Create contact       |
| `update_contact`               | extended | Update contact       |
| `delete_contact`               | extended | Delete contact       |
| `list_contact_groups`          | extended | List contact groups  |
| `get_contact_group`            | extended | Get group details    |
| `batch_create_contacts`        | complete | Batch create         |
| `batch_update_contacts`        | complete | Batch update         |
| `batch_delete_contacts`        | complete | Batch delete         |
| `create_contact_group`         | complete | Create group         |
| `update_contact_group`         | complete | Update group         |
| `delete_contact_group`         | complete | Delete group         |
| `modify_contact_group_members` | complete | Manage group members |

### Google Custom Search (3 tools)

| Tool                         | Tier     | Description                      |
| ---------------------------- | -------- | -------------------------------- |
| `search_custom`              | core     | Programmable Search Engine query |
| `search_custom_siterestrict` | extended | Site-restricted search           |
| `get_search_engine_info`     | complete | Get search engine config         |

### Google Apps Script (17 tools)

| Tool                    | Tier     | Description             |
| ----------------------- | -------- | ----------------------- |
| `list_script_projects`  | core     | List user's scripts     |
| `get_script_project`    | core     | Get script metadata     |
| `get_script_content`    | core     | Get script source code  |
| `create_script_project` | core     | Create new script       |
| `update_script_content` | core     | Update script code      |
| `run_script_function`   | core     | Execute script function |
| `generate_trigger_code` | core     | Generate trigger code   |
| `create_deployment`     | extended | Deploy script           |
| `list_deployments`      | extended | List deployments        |
| `update_deployment`     | extended | Update deployment       |
| `delete_deployment`     | extended | Remove deployment       |
| `delete_script_project` | extended | Delete script           |
| `list_versions`         | extended | List versions           |
| `create_version`        | extended | Create version          |
| `get_version`           | extended | Get version details     |
| `list_script_processes` | extended | List running processes  |
| `get_script_metrics`    | extended | Get script metrics      |

</details>

## Credential storage

Tokens live as JSON under `~/.google_workspace_mcp/credentials/` (`{email}.json`). Directory is `0700`, files are `0600`. Override with `WORKSPACE_MCP_CREDENTIALS_DIR`.

## Development

```bash
go test ./...
go test ./tools/ -run TestGmail
INTEGRATION_TEST_EMAIL="you@gmail.com" go test -tags integration ./tools/
```

- **Unit** — formatting, parsing, helpers
- **Protocol** — MCP `tools/call` validation and error paths
- **Mock API** — handlers against `httptest.Server` fixtures
- **Integration** — real Google APIs (`integration` build tag; needs `INTEGRATION_TEST_EMAIL`)

## Limitations

- **stdio only** — no HTTP server mode ([Python version](https://github.com/taylorwilsdon/google_workspace_mcp) has that)
- **Single-user** — no multi-user sessions or OAuth 2.1
- **Local MCP clients** — not aimed at hosted multi-tenant deployments

## Acknowledgments

Go rewrite of [google_workspace_mcp](https://github.com/taylorwilsdon/google_workspace_mcp) by [Taylor Wilsdon](https://github.com/taylorwilsdon). The Python project remains the full-featured reference for multi-user and HTTP deployments. MIT licensed.

## License

[MIT](LICENSE)
