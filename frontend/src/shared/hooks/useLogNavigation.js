// src/shared/hooks/useLogNavigation.js
import { useEffect } from "react"
import { useLocation } from "react-router-dom"
import { loggerService } from "@/shared/services/loggerService"

/**
 * useLogNavigation
 *
 * Hook that automatically logs every route change.
 * - Captures pathname from react-router
 * - Sends structured log to loggerService
 *
 * Usage:
 *   Place once in App.jsx → useLogNavigation()
 */
export function useLogNavigation() {
  const location = useLocation()

  useEffect(() => {
    loggerService.debug("Navigation", { path: location.pathname })
  }, [location])
}
