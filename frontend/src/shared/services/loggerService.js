// src/shared/services/loggerService.js
//
// Frontend logger service.
// - Always logs to browser console (for dev visibility).
// - Proxies structured logs to the Go Zap logger via Wails bindings.
// - Uses centralized formatting for consistency.

import { loggerBindings } from "@/shared/bindings/loggerBindings"

/**
 * format
 *
 * Build a structured log entry.
 * - Adds ISO timestamp
 * - Includes level, message, and optional metadata
 */
function format(level, msg, meta) {
  return {
    ts: new Date().toISOString(),
    level,
    msg,
    ...(meta && { meta }),
  }
}

export const loggerService = {
  info(msg, meta) {
    const entry = format("info", msg, meta)
    console.info("[frontend]", entry)
    loggerBindings.info(entry.msg, JSON.stringify(entry.meta || {}))
  },
  warn(msg, meta) {
    const entry = format("warn", msg, meta)
    console.warn("[frontend]", entry)
    loggerBindings.warn(entry.msg, JSON.stringify(entry.meta || {}))
  },
  error(msg, meta) {
    const entry = format("error", msg, meta)
    console.error("[frontend]", entry)
    loggerBindings.error(entry.msg, JSON.stringify(entry.meta || {}))
  },
  debug(msg, meta) {
    const entry = format("debug", msg, meta)
    console.debug("[frontend]", entry)
    loggerBindings.debug(entry.msg, JSON.stringify(entry.meta || {}))
  },
}
