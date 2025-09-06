import { routes } from "@/shared/routes/routes"

/**
 * Hook to resolve a route object by its key.
 *
 * - Useful when you only know the logical key of a route
 *   and want to retrieve its full definition (path, label, icon, etc.)
 * - Keeps route lookup logic centralized and reusable
 *
 * Usage:
 *   const route = useCurrentRoute("projects")
 *
 * Returns:
 *   Route object if found, otherwise null
 */
export function useCurrentRoute(routeKey) {
  return routes.find((r) => r.key === routeKey) || null
}
