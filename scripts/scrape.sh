#!/usr/bin/env bash
set -euo pipefail

config_path="scrape.yaml"
build=false

while [ "$#" -gt 0 ]; do
  case "$1" in
    --build)
      build=true
      shift
      ;;
    --config)
      config_path="${2:?--config requires a value}"
      shift 2
      ;;
    -h|--help)
      echo "usage: scripts/scrape.sh [--build] [--config scrape.yaml]" >&2
      exit 0
      ;;
    *)
      echo "usage: scripts/scrape.sh [--build] [--config scrape.yaml]" >&2
      exit 2
      ;;
  esac
done

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$repo_root"

if [ ! -f "$config_path" ]; then
  echo "config file not found: $config_path" >&2
  exit 1
fi

if [ "$build" = true ]; then
  docker compose build scraper
fi

docker compose run --rm scraper run --config "/vault/$config_path"
