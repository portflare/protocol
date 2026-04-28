# Portflare Protocol

This repository contains shared wire-level types and validation helpers for Portflare.

Its purpose is to let the server and client share protocol definitions without making either repository depend on the other's implementation packages.

## What belongs here

- websocket message types shared by server and client
- common JSON payload structs
- protocol constants
- lightweight validation helpers used on both sides

## What does not belong here

- server business logic
- client runtime logic
- persistence code
- dashboard or UI code

## Current packages

- `types`: shared protocol structs and message type constants
- `validation`: shared lightweight validation helpers such as client key checks

## Goal

Keep `github.com/portflare/server` and `github.com/portflare/client` independent while giving both a stable place for shared wire contracts.

## Release and tagging

This repository is set up to use Release Please for automatic version PRs and tags.

Expected workflow:

- merge conventional commits into `main` or `master`
- Release Please opens or updates a release PR
- merging that release PR creates the version tag and GitHub release

Recommended commit prefixes include:

- `feat:` for new protocol surface
- `fix:` for protocol corrections
- `docs:` for documentation-only changes
- `test:` for test-only changes
- `refactor:` for internal reshaping without protocol behavior changes

When `server` and `client` start consuming published protocol versions, tag releases here first and then update those repositories to the new protocol version.
