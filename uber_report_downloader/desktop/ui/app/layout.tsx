import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Uber Report Generator (by Carrum Tech Team)",
  description: "Created By Kapil Rohilla",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
