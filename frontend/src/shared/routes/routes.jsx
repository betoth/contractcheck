import HomePage from "@/features/home/HomePage"
import ProjectsPage from "@/features/projects/ProjectsPage"
import { Home, FolderKanban } from "lucide-react"

/**
 * routes
 *
 * Centralized route definitions for the application.
 *
 * - Each route may define:
 *   • path: URL pattern
 *   • element: component to render
 *   • icon: optional icon for navigation
 *   • breadcrumbs: array for navigation hierarchy
 *   • children: nested routes (supports dynamic segments)
 *   • custom flags (e.g., isDefault)
 *
 * - Keeps sidebar, header, and router consistent.
 * - Allows dynamic breadcrumbs (e.g., ":projectName").
 */
export const routes = [
  {
    key: "home",
    path: "/",
    isDefault: true, // Custom flag to mark this as the default route
    label: "Home",
    icon: Home,
    element: <HomePage />,
    breadcrumbs: [{ label: "Home" }],
  },
  {
    key: "projects",
    path: "/projects",
    label: "Projects",
    icon: FolderKanban,
    element: <ProjectsPage />,
    breadcrumbs: [{ label: "Projects" }],
    // Nested routing example: details inside projects
    children: [
      {
        key: "project-details",
        path: "/projects/:id", // dynamic segment (project id)
        label: "Project Details",
        element: <div>Project Details Placeholder</div>, // replace with ProjectDetailsPage
        // Dynamic breadcrumbs: placeholder resolved later (e.g., project name)
        breadcrumbs: [
          { label: "Projects", href: "/projects" },
          { label: ":projectName" },
        ],
      },
    ],
  },
]
