import { defineConfig } from "vite"
import react from "@vitejs/plugin-react"
import path from "node:path"

export default defineConfig({
  plugins: [react()],
  server: { host: true, port: 5173, strictPort: true },
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
      "@wailsjs": path.resolve(__dirname, "./wailsjs"),
    },
  },
  build: { outDir: "dist", emptyOutDir: true },
})
