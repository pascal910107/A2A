# AGENTS.md

## Cursor Cloud specific instructions

This is a **documentation/specification project** (A2A Protocol), not a runnable application. The primary "service" is the MkDocs documentation site.

### Key commands

| Task | Command |
|---|---|
| Install deps | `pip install -r requirements-docs.txt` |
| Dev server | `DISABLE_MKDOCS_2_WARNING=true mkdocs serve -a 0.0.0.0:8000` |
| Build site | `mkdocs build` |
| Full build (with proto/SDK docs) | `./scripts/build_docs.sh` |
| Lint (Python) | `ruff check .` |
| Format check | `ruff format --check .` |

### Gotchas

- The `mkdocs serve` command currently emits a large warning about MkDocs 2.0 compatibility. Set `DISABLE_MKDOCS_2_WARNING=true` to suppress it.
- `ruff` is not listed in `requirements-docs.txt` but is used for linting Python files (configured in `.ruff.toml`). Install separately with `pip install ruff`.
- The `sdk/python.md` nav entry and some anchor links in docs produce warnings during build — these are pre-existing and not blocking.
- Proto-to-JSON-Schema regeneration (`scripts/proto_to_json_schema.sh`) requires `protoc`, `protoc-gen-jsonschema`, and `jq`. These are optional for docs work; the devcontainer `setup.sh` installs them.
- The dev server runs on port **8000** by default.
