import React from "react"
import ReactDOM from "react-dom/client"
import App from "./App.jsx"
import { AppProviders } from "@/shared/providers/AppProviders"
import "./styles/tailwind.css"

/**
 * Application entrypoint.
 *
 * - Safely locates the root DOM node.
 * - Mounts the App component wrapped with global providers.
 * - Uses React.StrictMode to surface potential issues during development.
 */

// Locate root element safely.
// Throwing an explicit error avoids silent failures
// if the expected DOM structure is missing.
const rootElement = document.getElementById("root")
if (!rootElement) {
  throw new Error(
    "Root element #root not found. Make sure index.html contains <div id='root'></div>"
  )
}

// Create the React root and render the app.
// AppProviders wraps App with global contexts (e.g. theme, query client, i18n).
ReactDOM.createRoot(rootElement).render(
  <React.StrictMode>
    <AppProviders>
      <App />
    </AppProviders>
  </React.StrictMode>
)
