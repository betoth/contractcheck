/**
 * Barrel file for services.
 *
 * - Re-exports all service modules from a single entrypoint.
 * - Keeps imports clean and consistent across the app.
 *
 * Usage:
 *   import { appService, loggerService } from "@/shared/services"
 */
export { appService } from "./appService"
export { loggerService } from "./loggerService"
