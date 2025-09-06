// Sidebar navigation using shadcn/ui components
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupLabel,
  SidebarGroupContent,
  SidebarMenu,
  SidebarMenuItem,
  SidebarMenuButton,
  SidebarHeader,
  SidebarFooter,
  SidebarSeparator,
} from "@/components/ui/sidebar"
import logo from "@/assets/appicon.png"
import { ThemeToggle } from "@/components/ui/theme-toggle"
import { NavLink } from "react-router-dom"
import { useTheme } from "next-themes"

/**
 * SidebarNavLink
 *
 * - Wrapper around NavLink to integrate with SidebarMenuButton.
 * - Preserves active state styling via react-router.
 */
function SidebarNavLink({ to, icon: Icon, label }) {
  return (
    <NavLink to={to}>
      {({ isActive }) => (
        <SidebarMenuButton
          isActive={isActive}
          className="data-[active=true]:bg-primary data-[active=true]:text-primary-foreground"
        >
          <Icon className="size-4" />
          <span>{label}</span>
        </SidebarMenuButton>
      )}
    </NavLink>
  )
}

/**
 * AppSidebar
 *
 * - Main navigation sidebar for the application.
 * - Contains:
 *   • Header with brand logo + app name
 *   • Navigation group (routes)
 *   • Appearance group (theme toggle)
 *   • Footer with version information
 *
 * Props:
 * - version {string} → current app version from backend
 * - routes {Array} → list of navigation routes with key, label, icon, path
 */
export default function AppSidebar({ version, routes }) {
  const { theme, setTheme } = useTheme()

  return (
    <Sidebar>
      {/* Header: brand logo + app name */}
      <SidebarHeader>
        <div className="flex h-12 items-center gap-2 border-b px-3">
          <img src={logo} alt="ContractCheck logo" className="h-8 w-8 rounded-sm" />
          <span className="text-lg font-semibold">ContractCheck</span>
        </div>
      </SidebarHeader>

      <SidebarContent>
        {/* Navigation group: dynamic routes */}
        <SidebarGroup>
          <SidebarGroupLabel>Navigation</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              {routes.map(({ key, label, icon, path }) => (
                <SidebarMenuItem key={key}>
                  <SidebarNavLink to={path} icon={icon} label={label} />
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>

        {/* Appearance group: theme switcher */}
        <SidebarGroup>
          <SidebarGroupLabel>Appearance</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              <SidebarMenuItem>
                {/* Toggle theme like a menu item */}
                <SidebarMenuButton
                  onClick={() => setTheme(theme === "dark" ? "light" : "dark")}
                  className="flex w-full items-center gap-2"
                >
                  <ThemeToggle asChild />
                  <span>{theme === "dark" ? "Dark mode" : "Light mode"}</span>
                </SidebarMenuButton>
              </SidebarMenuItem>
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>

      {/* Footer: app version info */}
      <SidebarFooter>
        <SidebarSeparator className="my-2" />
        <div className="px-2 py-2 text-xs text-muted-foreground">
          v{version || "…"}
        </div>
      </SidebarFooter>
    </Sidebar>
  )
}
