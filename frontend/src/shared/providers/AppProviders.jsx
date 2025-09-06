// Centralized provider wrapper for the application.
// - Ensures main.jsx remains clean
// - Add here global providers like Theme, Router, Query, State, i18n, etc.
import { ThemeProvider } from "@/shared/providers/ThemeProvider"
import { MemoryRouter } from "react-router-dom"

export function AppProviders({ children }) {
  return (
    <ThemeProvider>
      <MemoryRouter>
        {children}
      </MemoryRouter>
    </ThemeProvider>
  )
}
