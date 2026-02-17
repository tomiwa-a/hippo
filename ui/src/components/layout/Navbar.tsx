import { Button } from '../ui/Button';

export function Navbar() {
  return (
    <nav className="fixed top-0 left-0 right-0 z-50 transition-all border-b border-white/5 bg-zinc-950/80 backdrop-blur-md">
      <div className="max-w-7xl mx-auto px-6 h-16 flex items-center justify-between">
        {/* Logo Area */}
        <div className="flex items-center space-x-2">
          <div className="w-8 h-8 rounded-full bg-rose-500/10 flex items-center justify-center">
            <svg 
              viewBox="0 0 24 24" 
              className="w-5 h-5 text-rose-500"
              fill="none" 
              stroke="currentColor" 
              strokeWidth="2"
            >
              <path strokeLinecap="round" strokeLinejoin="round" d="M12 4v4m0 0l-2 2m2-2l2 2m-2-2v4m0 4v4m0 0l-2-2m2 2l2-2" />
            </svg>
          </div>
          <span className="text-lg font-bold tracking-tight text-white">
            Hippo
          </span>
        </div>

        {/* Navigation Links */}
        <div className="hidden md:flex items-center space-x-8">
          <a href="#features" className="text-sm font-medium text-zinc-400 hover:text-white transition-colors">
            Features
          </a>
          <a href="#secure" className="text-sm font-medium text-zinc-400 hover:text-white transition-colors">
            Secure Fields
          </a>
          <a href="#docs" className="text-sm font-medium text-zinc-400 hover:text-white transition-colors">
            Docs
          </a>
        </div>

        {/* CTA */}
        <div className="flex items-center space-x-4">

          <Button size="sm">
            Get Started
          </Button>
        </div>
      </div>
    </nav>
  );
}
