import { useEffect, useState } from "react"
import { appService } from "@/shared/services/appService"

/**
 * Custom hook to retrieve the current app version.
 *
 * - Encapsulates the call to `appService.getVersion()`
 * - Manages state internally using React's useState
 * - Implements a mounted flag to avoid setting state
 *   after the component has unmounted (memory leak prevention)
 *
 * Usage:
 *   const version = useAppVersion()
 *
 * Returns:
 *   string — current app version, empty string if not yet loaded
 */
export function useAppVersion() {
  const [version, setVersion] = useState("")

  useEffect(() => {
    let isMounted = true

    appService.getVersion().then((v) => {
      if (isMounted) setVersion(v)
    })

    return () => {
      // Cleanup: prevent state updates if component is unmounted
      isMounted = false
    }
  }, [])

  return version
}
