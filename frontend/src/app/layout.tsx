import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import './styles/style.css';


const inter = Inter({ subsets: ['vietnamese'] })

export const metadata: Metadata = {
  title: '勉強管理アプリ',
  description: '勉強管理アプリ',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="ja">
      <body className={inter.className}>{children}</body>
    </html>
  )
}
