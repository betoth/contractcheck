"use client"

import * as React from "react"
import { Moon, Sun } from "lucide-react"
import { useTheme } from "next-themes"
import { cn } from "@/shared/lib/utils"

// ThemeToggle is now flexible: can render standalone OR inside a parent (using asChild)
export function ThemeToggle({ asChild = false, className, ...props }) {
  const { theme, setTheme } = useTheme()

  // Handle toggle
  const toggleTheme = () => {
    setTheme(theme === "dark" ? "light" : "dark")
  }

  // When `asChild` is true, it won't render its own <button>
  if (asChild) {
    return (
      <>
        {theme === "dark" ? (
          <Moon className={cn("h-4 w-4", className)} {...props} />
        ) : (
          <Sun className={cn("h-4 w-4", className)} {...props} />
        )}
      </>
    )
  }

  return (
    <button
      type="button"
      onClick={toggleTheme}
      className={cn(
        "inline-flex items-center justify-center rounded-md border border-input bg-background p-2 text-sm transition-colors hover:bg-accent hover:text-accent-foreground focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none disabled:opacity-50",
        className
      )}
      {...props}
    >
      {theme === "dark" ? (
        <Moon className="h-4 w-4" />
      ) : (
        <Sun className="h-4 w-4" />
      )}
      <span className="sr-only">Toggle theme</span>
    </button>
  )
}
