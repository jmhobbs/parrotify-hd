import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  port: 3000,
  strictPort: true,
  plugins: [react()]
})
