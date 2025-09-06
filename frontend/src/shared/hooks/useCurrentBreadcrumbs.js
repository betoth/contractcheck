import { useLocation, useParams } from "react-router-dom"
import { routes } from "@/shared/routes/routes"

/**
 * Hook to resolve breadcrumbs dynamically based on the current route.
 *
 * - Detects the active route based on `location.pathname`
 * - Supports dynamic segments (e.g. /projects/:id)
 * - Substitutes placeholders like `:projectName` with values from route params
 *
 * Usage:
 *   const breadcrumbs = useCurrentBreadcrumbs()
 *
 * Returns:
 *   Array<{ label: string, href?: string }>
 *   Example:
 *     /projects           -> [{ label: "Projects" }]
 *     /projects/123       -> [{ label: "Projects", href: "/projects" }, { label: "123" }]
 *
 * Future-proof:
 *   - Can fetch project name by ID (instead of raw "123") if integrated with a data service.
 */
export function useCurrentBreadcrumbs() {
  const location = useLocation()
  const params = useParams()

  // Flatten routes with children so nested routes are also searchable
  const flattenRoutes = (routes) =>
    routes.flatMap((r) => [r, ...(r.children || [])])

  const allRoutes = flattenRoutes(routes)

  // Find current route by pathname (supports both static and dynamic)
  const currentRoute =
    allRoutes.find((r) => {
      // Convert route.path ("/projects/:id") into a regex
      const pattern = new RegExp(
        "^" +
          r.path
            .replace(/:[^/]+/g, "[^/]+") // replace params with wildcard
            .replace(/\//g, "\\/") + // escape slashes
          "$"
      )
      return pattern.test(location.pathname)
    }) || null

  if (!currentRoute) {
    return [{ label: "Not Found" }]
  }

  // Replace placeholders (e.g. ":projectName") with param values
  return (currentRoute.breadcrumbs || []).map((crumb) => {
    if (crumb.label?.startsWith(":")) {
      const key = crumb.label.slice(1) // remove ":"
      return { ...crumb, label: params[key] || crumb.label }
    }
    return crumb
  })
}
