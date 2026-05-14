#!/usr/bin/env bash
set -euo pipefail

push_after_commit=false
commit_message="chore: update scraped inbox"

while [ "$#" -gt 0 ]; do
  case "$1" in
    --push)
      push_after_commit=true
      shift
      ;;
    --message)
      commit_message="${2:?--message requires a value}"
      shift 2
      ;;
    *)
      echo "usage: scripts/scrape-and-commit.sh [--push] [--message MESSAGE]" >&2
      exit 2
      ;;
  esac
done

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$repo_root"

scripts/scrape.sh

if [ -z "$(git status --porcelain -- research-inbox book-inbox .scrapem/seen-research-urls.json .scrapem/seen-book-urls.json)" ]; then
  echo "No scraped inbox changes to commit."
  exit 0
fi

git add research-inbox book-inbox .scrapem/seen-research-urls.json .scrapem/seen-book-urls.json
git commit -m "$commit_message"

if [ "$push_after_commit" = true ]; then
  git push origin HEAD
fi
