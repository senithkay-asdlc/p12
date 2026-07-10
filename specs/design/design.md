---
skillsApplied:
  - high-level-architecture
  - go
  - openapi-conventions
---

# Design: Hello World API

## 1. Overview

A minimal backend system consisting of a single HTTP service that returns a
friendly greeting and reports its own liveness. There is no UI, no
authentication, and no persistent storage.

## 2. Components

- **hello-api** — `service`. Exposes the greeting and health endpoints; the
  entire system in one deployable unit.

## 3. Capabilities

### hello-api

- **Greeting** — respond to `GET /hello` with a JSON greeting message,
  optionally personalized via a `name` query parameter (requirements 4.1, 4.2).
- **Health check** — respond to `GET /health` with `200 OK` for platform
  liveness/readiness probes (requirement 4.3).

## 4. Data model

No persistent entities. The only exchanged shape is the greeting response:

- **Greeting** — `message: string`.

## 5. Roles & access

No authentication or roles — the API is open and unauthenticated, per
non-goals.

## 6. Interactions

None — `hello-api` has no dependencies on other components or external
systems; it serves requests standalone.

## 7. Data flow

1. A client issues `GET /hello` (optionally with `?name=Alice`).
2. `hello-api` builds the greeting message in-process (no external calls) and
   returns `200 OK` with the JSON body.
3. A platform liveness probe issues `GET /health`; `hello-api` returns
   `200 OK` immediately.
