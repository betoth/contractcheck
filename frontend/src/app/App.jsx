import { Suspense } from "react"
import { Routes, Route } from "react-router-dom"
import { SidebarProvider } from "@/components/ui/sidebar"
import { AppSidebar } from "@/shared/components/organisms/AppSidebar"
import { AppHeader } from "@/shared/components/organisms/AppHeader"
import { useAppVersion, useCurrentBreadcrumbs, useLogNavigation } from "@/shared/hooks"
import { routes } from "@/shared/routes/routes"
import { Skeleton } from "@/components/ui/skeleton"
import NotFoundPage from "@/features/common/NotFoundPage" // Fallback page for unknown routes
import "./styles/tailwind.css"

/**
 * Root application component.
 *
 * Responsibilities:
 * - Provides global layout (Sidebar + Header + Main content).
 * - Registers routes dynamically from central config.
 * - Handles breadcrumbs, app version, and 404 fallback.
 * - Logs route navigation via `useLogNavigation`.
 * - Wraps everything with sidebar provider for layout consistency.
 */
export default function App() {
  // Fetch current app version from backend
  const version = useAppVersion()

  // Resolve breadcrumb chain dynamically based on current route
  const breadcrumbs = useCurrentBreadcrumbs()

  // Log navigation events (path, time) to backend + console
  useLogNavigation()

  return (
    <SidebarProvider>
      {/* Application wrapper with ARIA role for accessibility */}
      <div className="flex h-screen" role="application">
        {/* Sidebar navigation (routes + version) */}
        <AppSidebar version={version} routes={routes} />

        {/* Main content area */}
        <div className="flex flex-1 flex-col">
          {/* Persistent header with breadcrumbs */}
          <AppHeader breadcrumbs={breadcrumbs} role="banner" />

          {/* Scrollable main container */}
          <main
            className="flex-1 overflow-auto p-4"
            aria-label="Main content"
            role="main"
          >
            {/* Suspense shows skeleton while lazy-loaded routes resolve */}
            <Suspense fallback={<Skeleton className="h-8 w-32" />}>
              <Routes>
                {/* Dynamically map routes */}
                {routes.map(({ path, element }) => (
                  <Route key={path} path={path} element={element} />
                ))}

                {/* Fallback: unknown routes → 404 page */}
                <Route path="*" element={<NotFoundPage />} />
              </Routes>
            </Suspense>
          </main>
        </div>
      </div>
    </SidebarProvider>
  )
}
