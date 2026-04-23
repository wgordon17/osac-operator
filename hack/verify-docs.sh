#!/usr/bin/env bash
set -euo pipefail

TMPDIR="$(mktemp -d)"
trap 'rm -rf "$TMPDIR"' EXIT

# Ensure crd-ref-docs is installed; resolve absolute path via repo root
make crd-ref-docs >/dev/null 2>&1
REPO_ROOT="$(git rev-parse --show-toplevel)"
CRD_REF_DOCS="$REPO_ROOT/bin/crd-ref-docs"

"$CRD_REF_DOCS" \
  --source-path=api/v1alpha1 \
  --renderer=markdown \
  --config=docs/reference/crd-ref-config.yaml \
  --output-path="$TMPDIR/crds.md"

if ! diff docs/reference/crds.md "$TMPDIR/crds.md"; then
  printf "CRD docs are stale. Run 'make docs-generate' to update.\n" >&2
  exit 1
fi
