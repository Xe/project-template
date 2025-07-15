#!/usr/bin/env bash
set -euo pipefail

# 1. Determine repo root and name
repo_root="$(git rev-parse --show-toplevel)"
repo_name="$(basename "$repo_root")"

# 2. Patch devcontainer.json
json_file=".devcontainer/devcontainer.json"
tmp_json="${json_file}.tmp"

if [[ ! -f "$json_file" ]]; then
  echo "❌ Error: $json_file not found." >&2
  exit 1
fi

jq --arg wf "/workspace/$repo_name" \
   '.workspaceFolder = $wf' \
   "$json_file" > "$tmp_json" \
&& mv "$tmp_json" "$json_file"

echo "✅ Set workspaceFolder to /workspace/$repo_name in $json_file"

# 3. Patch docker-compose.yaml
yaml_file=".devcontainer/docker-compose.yaml"
backup_suffix=".bak"

if [[ ! -f "$yaml_file" ]]; then
  echo "❌ Error: $yaml_file not found." >&2
  exit 1
fi

# Replace lines like "- ../:/workspace/project-template:cached"
# with "- ../:/workspace/<repo_name>:cached"
sed -i"$backup_suffix" -E \
  "s@^(\s*-\s*\.\./:)[^:]+(:cached)@\1/workspace/$repo_name\2@" \
  "$yaml_file"

# drop backup if you don't need it
rm -f "${yaml_file}${backup_suffix}"

echo "✅ Updated volume mapping to /workspace/$repo_name:cached in $yaml_file"