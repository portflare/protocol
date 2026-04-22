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
