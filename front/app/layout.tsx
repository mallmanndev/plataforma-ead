import './globals.css'
import type {Metadata} from 'next'
import {Inter} from 'next/font/google'

const inter = Inter({subsets: ['latin']})

import {ThemeProvider} from "@/components/ui/theme-provider"

export default function RootLayout({children}: { children: React.ReactNode }) {
    return (
        <>
            <html lang="en" suppressHydrationWarning>
            <head/>
            <body>
            <ThemeProvider
                attribute="class"
                defaultTheme="dark"
                enableSystem
                disableTransitionOnChange
            >
                {children}
            </ThemeProvider>
            </body>
            </html>
        </>
    )
}

