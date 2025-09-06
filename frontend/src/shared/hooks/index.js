/**
 * Centralized exports for all custom React hooks in the application.
 *
 * Purpose:
 * - Provides a single entry point for importing hooks, improving discoverability.
 * - Reduces repetitive import paths across the codebase.
 * - Encourages consistent usage of hooks by exposing only the intended public API.
 *
 * Hooks exported:
 * - useAppVersion         → Retrieves the current application version (from env/config).
 * - useCurrentBreadcrumbs → Manages and exposes breadcrumb state for the active route.
 * - useCurrentRoute       → Provides details about the active route (path, params, etc.).
 * - useLogNavigation      → Logs navigation events for analytics or debugging purposes.
 *
 * Usage:
 * ```tsx
 * import { useAppVersion, useCurrentRoute } from "@/hooks"
 *
 * const MyComponent = () => {
 *   const version = useAppVersion()
 *   const route = useCurrentRoute()
 *   ...
 * }
 * ```
 *
 * @module hooks
 */
export { useAppVersion } from "./useAppVersion"
export { useCurrentBreadcrumbs } from "./useCurrentBreadcrumbs"
export { useCurrentRoute } from "./useCurrentRoute"
export { useLogNavigation } from "./useLogNavigation"
