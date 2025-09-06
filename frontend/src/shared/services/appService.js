// src/shared/services/appService.js
//
// Service layer to interact with Go (Wails) backend.
// - Wraps the auto-generated bindings in a safe API.
// - Adds centralized error handling via loggerService.
// - Organized by domain (app, projects, config).
//
// Future-proof: expand projects/config when new Go bindings exist.

import { appBindings } from "@/shared/bindings/appBindings"
import { loggerService } from "./loggerService"

/**
 * safe
 *
 * Wrapper for async calls.
 * - Logs error centrally
 * - Returns fallback value if provided
 */
async function safe(promise, fallback = null) {
  try {
    return await promise
  } catch (err) {
    loggerService.error("[appService] error", { err })
    return fallback
  }
}

export const appService = {
  /** Get app version from Go binding */
  async getVersion() {
    const version = await safe(appBindings.getVersion(), null)
    if (!version) throw new Error("App version unavailable")
    return version
  },

  /** Project-related domain calls (placeholders for now) */
  projects: {
    async getAll() {
      loggerService.info("[appService.projects] getAll called")
      // TODO: replace with Go binding
      return []
    },
    async getById(id) {
      loggerService.info("[appService.projects] getById", { id })
      // TODO: replace with Go binding
      return null
    },
  },

  /** Config-related domain calls (placeholders for now) */
  config: {
    async get() {
      loggerService.info("[appService.config] get called")
      // TODO: replace with Go binding
      return {}
    },
  },
}
