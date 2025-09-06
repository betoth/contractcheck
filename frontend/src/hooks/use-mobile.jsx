import * as React from "react"

const MOBILE_BREAKPOINT = 768

/**
 * useIsMobile
 *
 * Hook to detect if the viewport is considered "mobile".
 *
 * - Uses `window.matchMedia` to observe viewport changes.
 * - Returns `true` if the screen width is below MOBILE_BREAKPOINT.
 * - Simplified for desktop app context (Wails).
 *
 * @returns {boolean} true if viewport width < MOBILE_BREAKPOINT
 */
export function useIsMobile() {
  const [isMobile, setIsMobile] = React.useState(false)

  React.useEffect(() => {
    // Create a MediaQueryList for the given breakpoint
    const mql = window.matchMedia(`(max-width: ${MOBILE_BREAKPOINT - 1}px)`)

    // Update state on change
    const onChange = (e) => {
      setIsMobile(e.matches)
    }

    mql.addEventListener("change", onChange)
    setIsMobile(mql.matches) // Initial state update

    return () => mql.removeEventListener("change", onChange) // Cleanup listener
  }, [])

  return isMobile
}
