import { useNavigate } from "react-router-dom"
import { Button } from "@/components/ui/button"

/**
 * NotFoundPage (404)
 *
 * - Centralized fallback page for unknown or invalid routes.
 * - Provides quick navigation options:
 *    • Go back to the previous page
 *    • Go home (root route "/")
 * - Keeps UX consistent across the app when users hit a non-existing path.
 */
export default function NotFoundPage() {
  const navigate = useNavigate()

  return (
    <div className="flex flex-1 flex-col items-center justify-center text-center space-y-4">
      {/* Page title (large 404 code) */}
      <h1 className="text-6xl font-bold text-primary">404</h1>

      {/* Short explanation message */}
      <p className="text-lg text-muted-foreground">
        Oops! The page you are looking for doesn’t exist.
      </p>

      {/* Navigation actions */}
      <div className="flex gap-2">
        <Button variant="default" onClick={() => navigate(-1)}>
          Go Back
        </Button>
        <Button variant="outline" onClick={() => navigate("/")}>
          Go Home
        </Button>
      </div>
    </div>
  )
}
