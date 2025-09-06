// src/shared/bindings/appBindings.js
// Thin wrapper around Wails auto-generated Go bindings.
// This is the ONLY place that imports from "@wailsjs/go/..."
// Everywhere else in the app should import from this file.

import { Version } from "@wailsjs/go/wailsapp/App"

export const appBindings = {
  getVersion: Version,
}
