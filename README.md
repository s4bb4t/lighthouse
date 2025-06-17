# 📡 LightHouse: Structured Errors, Logging & Alerts for Go

**LightHouse** is a modular Go library for error handling, structured logging, and alerting. You can use the entire
suite or select components independently, depending on your stack and requirements.

> 🚀 **Mission:** Help teams ship faster and debug smarter with rich, traceable, and readable errors — designed for
> production, logs, and humans.

---

## 🧩 Navigation

- [Overview](#overview)
- [SPError — Structured Errors](#sperror--structured-errors)
    - [Core Structure](#core-structure)
    - [Quick Start](#quick-start)
    - [Creating Errors](#creating-errors)
    - [Error Wrapping](#error-wrapping)
    - [Error Unwrapping](#error-unwrapping)
    - [Using Spin()](#using-spin)
- [Best Practices](#best-practices)
- [Predefined Errors](#predefined-errors)
- [Logger](#Logger)

---

## Overview

LightHouse helps you build and trace **rich error chains** that carry context, localization, levels, and full stack
traces.

```go
// Example: Creating and wrapping a structured error
err := sp.Any(errors.New("something broke"), "Contract failed", "Check contract address")
```

---

## SPError — Structured Errors

### Core Structure

```go
type Error struct {
Core struct {
Desc   string // What happened?
Hint   string // How to fix it
Source string // File:line
Cause  error // Underlying error
}

User struct {
Messages map[string]string // Localized messages
HttpCode int               // HTTP status
Level    levels.Level // Severity level
}

meta               map[string]any
underlying         *Error
remainsUnderlying  int
}
```

Every field is designed for a different context:

- `Desc` → for logs
- `Hint` → for internal usage or tools
- `Messages` → for users
- `Level` → for filtering in telemetry (e.g. show only user errors)

---

### Quick Start

```go
err := sp.Any(errors.New("sql: connection refused"), "DB failure", "Check connection string")
```

This produces a rich `Error` with:

- Default HTTP code (500)
- `LevelError`
- Source info auto-filled

---

### Creating Errors

Use `sp.New(...)` if you want full control:

```go
err := sp.New(sp.Sample{
Messages: map[string]string{
sp.En: "Data not found",
},
Desc:     "The user ID does not exist",
Hint:     "Check if the user is registered",
HttpCode: 404,
Level:    levels.LevelUser,
Meta: map[string]any{
"userID": id,
},
})
```

Or use the fluent builder API:

```go
err := sp.NewError().
SetDesc("User not found").
SetHint("Verify ID").
SetCode(404).
SetLevel(levels.LevelUser)
```

---

### Error Wrapping

Use `Wrap()` to build a chain of errors with growing context:

```go
base := sp.New(sp.Sample{Desc: "DB failed"})
wrapped := sp.WrapNew(base, sp.Sample{Desc: "User service failed"})
```

Or use the method form:

```go
outer := sp.New(sp.Sample{Desc: "App failure"}).Wrap(err)
```

---

### Error Unwrapping

You can unwrap structured errors like standard Go errors:

```go
errors.Unwrap(err) // returns .underlying or .Cause
```

---

### Using Spin()

Spin lets you extract the most relevant error **by level**:

```go
// Returns the last error with Level <= LevelUser
safe := err.Spin(levels.LevelUser)
fmt.Println(safe.Msg(sp.En))
```

Use this before displaying errors to users or telemetry.

---

## Best Practices

1. **Always wrap Exported functions errors.**
   ```go
   _, err := db.Ping()
   return sp.Any(err, "desc", "hint")
   ```

2. **Unexported functions should return structured errors directly.**
   ```go
   return someUnexportedFunc()
   ```

3. **Use Spin() before returning errors externally.**
   ```go
   safe := err.Spin(levels.LevelUser)
   return safe
   ```

4. **Use AddMeta to enrich errors with traceability.**

---

## Predefined Errors

For reuse and clarity, define common errors statically:

```go
var ErrUnavailable = sp.New(sp.Sample{
Messages: map[string]string{
sp.En: "Temporarily unavailable",
},
Desc:     "Initialization in progress",
Hint:     "Try again later",
HttpCode: 307,
Level:    levels.LevelUser,
})
```

Then reuse:

```go
return ErrUnavailable.Copy().AddMeta("contract", contract)
```

---

# Logger

The `Logger` module in LightHouse provides a robust, environment-aware logging interface built on top of Go's `slog`. It
adds seamless integration with the `sperror` package, ensuring consistent, readable logs with full error context.

---

## ✨ Key Features

- Supports `local`, `dev`, and `prod` modes
- Pretty-printed output in local, JSON logs in other modes
- Error-aware logging integrated with `sperror`
- Includes support for levels and localized messages
- No-op mode for testing

---

## ⚙️ Initialization

```go
logger := logger.New(logger.Dev, "en", os.Stdout)
```

You can configure it for different stages:

- `logger.Local` — human-readable colored output (for CLI/local dev)
- `logger.Dev` — JSON format with debug info
- `logger.Prod` — JSON format with error-level filtering

The logger respects the language code (`lg`) to render messages from `sperror`.

---

## ⚡ Usage Examples

### Logging a structured error:

```go
err := sp.New(sp.Sample{
Messages: map[string]string{
sp.En: "Failed to connect to DB",
},
Desc:  "PostgreSQL timeout",
Hint:  "Check DB availability",
})

logger.Error(err)
```

### Logging at specific level:

```go
logger.ErrorWithLevel(err, levels.LevelError)
```

### Logging a debug or info message:

```go
logger.Debug("Starting process", "module", "sync")
logger.Info("Listening on port", "port", 8080)
```

### Logging warnings with optional error:

```go
logger.Warn("Failed to load config", err)
```

---

## ⚖️ Integration with `sperror`

When logging `sperror.Error`, the logger will:

- Call `.Spin(level)` to extract the correct level of detail
- Inject all available metadata via structured slog attributes

### Example:

```go
spErr := sp.New(sp.Sample{
Messages: map[string]string{sp.En: "Connection timeout"},
Desc: "Timeout after 10s",
Hint: "Check your network settings",
Level: levels.LevelError,
Meta: map[string]any{"retry": true},
})

logger.Error(spErr)
```

---

## 🎨 Pretty Handler

For `logger.Local`, LightHouse provides a colored, aligned output handler using [
`fatih/color`](https://github.com/fatih/color):

```text
[Jan 02 - 15:04:05] ERROR: Connection timeout
	retry = true
```

This makes logs more readable and compact for local development.

---

## ✅ Summary

| Feature     | Description                             |
|-------------|-----------------------------------------|
| Levels      | Info, Warn, Error, Debug (customizable) |
| Environment | Local (pretty), Dev/Prod (JSON)         |
| Error-aware | Yes — via `sperror` integration         |
| Language    | Yes — localized messages supported      |
| Test mode   | Yes — use `logger.Noop()`               |

Want to extend it with tracing, alerting, or error hooks? You're set up for it.

---

## ✅ Summary

LightHouse empowers your Go services with production-grade structured errors, allowing you to:

- Decouple internal and user messages
- Trace deeply with rich metadata
- Control visibility with `.Spin()` and `.Level`

> Build once. Trace always. Deploy with confidence.

---

Want help integrating into your project? Ping @s4bb4t 🚀
