import { Navbar } from './Navbar';

interface LayoutProps {
  children: React.ReactNode;
}

export function Layout({ children }: LayoutProps) {
  return (
    <div className="min-h-screen bg-zinc-950 text-zinc-200 font-sans antialiased selection:bg-rose-500/30 selection:text-rose-200">
      <Navbar />
      <main className="pt-16">
        {children}
      </main>
      
      {/* Subtle Background Noise/Mesh (Optional, keeping it clean for now) */}
      <div className="fixed inset-0 pointer-events-none z-[-1]">
        <div className="absolute inset-0 bg-[radial-gradient(circle_at_center,_var(--tw-gradient-stops))] from-zinc-900/20 via-zinc-950 to-zinc-950 opacity-50" />
      </div>
    </div>
  );
}
