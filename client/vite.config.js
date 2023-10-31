// vite.config.js
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import dns from 'dns'

dns.setDefaultResultOrder('verbatim')

export default defineConfig({
    server: {
        proxy: {
            '/api': 'http://localhost:9000'
        }
    },
    plugins: [react()]
})