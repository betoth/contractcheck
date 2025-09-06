// Utility to merge class names (shadcn pattern).
// - `clsx`: handles conditional class merging
// - `tailwind-merge`: resolves Tailwind conflicts (e.g. "px-2" vs "px-4")
import { clsx } from "clsx"
import { twMerge } from "tailwind-merge"

/**
 * Merges class names safely, respecting Tailwind rules.
 * Example:
 *   cn("p-2", condition && "bg-red-500")
 */
export function cn(...inputs) {
  return twMerge(clsx(...inputs))
}
