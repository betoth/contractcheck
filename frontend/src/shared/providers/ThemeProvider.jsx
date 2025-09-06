// src/shared/providers/ThemeProvider.jsx
import * as React from "react"
import { ThemeProvider as NextThemesProvider } from "next-themes"

/**
 * ThemeProvider
 *
 * - Wraps the application with `next-themes` provider.
 * - Enables theme switching between light, dark, and system preferences.
 * - Applies theme by toggling a `class` attribute on the <html> element.
 *
 * Props:
 * - children → React nodes to be wrapped
 * - ...props → forwarded to NextThemesProvider
 */
export function ThemeProvider({ children, ...props }) {
  return (
    <NextThemesProvider
      attribute="class" // applies `class="dark"` to <html>
      defaultTheme="system" // fallback to system preference if not set
      enableSystem
      disableTransitionOnChange // prevents flicker when switching themes
      {...props}
    >
      {children}
    </NextThemesProvider>
  )
}
