"use client";

import Header from "@/components/Header";
import ScrollToTop from "@/components/ScrollToTop";
import { Inter } from "next/font/google";
import "node_modules/react-modal-video/css/modal-video.css";
import "../styles/index.css";
import { Providers } from "./providers";

const inter = Inter({ subsets: ["latin"] });

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html suppressHydrationWarning lang="en">
      <head />
      <body
        className={`bg-black text-white dark:text-slate-200 ${inter.className}`}
      >
        <Providers>
          <Header />
          {children}
          <ScrollToTop />
        </Providers>
      </body>
    </html>
  );
}
