// src/shared/components/molecules/ThemeMenuItem/ThemeMenuItem.jsx
import { SidebarMenuButton } from "@/components/ui/sidebar"
import { ThemeToggle } from "@/components/ui/theme-toggle"
import { useTheme } from "next-themes"

/**
 * ThemeMenuItem
 *
 * - Sidebar menu item that toggles between light/dark themes.
 * - Uses `next-themes` for state management.
 * - No logging (clean UI logic only).
 */
export default function ThemeMenuItem() {
  const { theme, setTheme } = useTheme()

  // Toggle between dark/light themes
  const handleToggle = () => {
    const newTheme = theme === "dark" ? "light" : "dark"
    setTheme(newTheme)
  }

  return (
    <SidebarMenuButton
      onClick={handleToggle}
      className="flex w-full items-center gap-2 data-[active=true]:bg-primary data-[active=true]:text-primary-foreground"
    >
      <ThemeToggle />
      <span>{theme === "dark" ? "Dark mode" : "Light mode"}</span>
    </SidebarMenuButton>
  )
}
