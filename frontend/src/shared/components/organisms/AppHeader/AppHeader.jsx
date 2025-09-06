// src/shared/components/organisms/AppHeader/AppHeader.jsx
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb"

/**
 * AppHeader
 *
 * - Displays breadcrumb navigation when provided.
 * - Falls back to a simple page title if no breadcrumbs exist.
 * - Supports optional right-side actions (e.g., buttons, filters).
 * - Follows accessibility best practices using proper landmark roles and aria labels.
 */
export default function AppHeader({ title, breadcrumbs = [], actions }) {
  return (
    <header
      role="banner" // Accessibility: marks this header as the banner landmark
      className="sticky top-0 z-10 flex h-12 items-center justify-between border-b bg-background/80 px-4 backdrop-blur"
    >
      {/* Breadcrumb navigation if available */}
      {breadcrumbs.length > 0 ? (
        <Breadcrumb aria-label="Breadcrumb">
          <BreadcrumbList>
            {breadcrumbs.map((crumb, idx) => (
              <BreadcrumbItem key={crumb.href || crumb.label}>
                {crumb.href ? (
                  <BreadcrumbLink href={crumb.href}>{crumb.label}</BreadcrumbLink>
                ) : (
                  <BreadcrumbPage>{crumb.label}</BreadcrumbPage>
                )}
                {/* Render separator only between items */}
                {idx < breadcrumbs.length - 1 && <BreadcrumbSeparator />}
              </BreadcrumbItem>
            ))}
          </BreadcrumbList>
        </Breadcrumb>
      ) : (
        // Fallback: show title when no breadcrumbs are passed
        <h1 className="text-sm font-medium">{title}</h1>
      )}

      {/* Right-side actions if provided */}
      {actions && <div className="flex items-center gap-2">{actions}</div>}
    </header>
  )
}
