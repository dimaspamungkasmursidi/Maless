import './globals.css';

export const metadata = {
  title: 'Maless App',
  description: 'Maless or "Malas Less" meaning "Less Laziness" is a to-do list app designed to help you stay productive by reducing laziness.',
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
