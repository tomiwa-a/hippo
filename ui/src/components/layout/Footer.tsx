import { Button } from '../ui/Button';

export function Footer() {
  return (
    <footer className="bg-zinc-950 border-t border-white/5 py-12 md:py-16">
      <div className="container px-6 mx-auto">
        <div className="grid grid-cols-1 md:grid-cols-4 gap-8 mb-12">
          
          {/* Brand Column */}
          <div className="col-span-1 md:col-span-1">
            <div className="flex items-center space-x-2 mb-4">
              <div className="w-6 h-6 rounded-full bg-rose-500/10 flex items-center justify-center">
                <svg viewBox="0 0 24 24" className="w-4 h-4 text-rose-500" fill="none" stroke="currentColor" strokeWidth="2">
                  <path strokeLinecap="round" strokeLinejoin="round" d="M12 4v4m0 0l-2 2m2-2l2 2m-2-2v4m0 4v4m0 0l-2-2m2 2l2-2" />
                </svg>
              </div>
              <span className="text-lg font-bold tracking-tight text-white">Hippo</span>
            </div>
            <p className="text-zinc-500 text-sm mb-6">
              The semantic memory layer for your local environment. 
              <br />Open source and privacy-first.
            </p>
            <div className="flex space-x-4">
              {/* Social Icons (Placeholders) */}
              <a href="#" className="text-zinc-500 hover:text-white transition-colors">
                <span className="sr-only">GitHub</span>
                <svg className="w-5 h-5" fill="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                  <path fillRule="evenodd" d="M12 2C6.477 2 2 6.484 2 12.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0112 6.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.202 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.943.359.309.678.92.678 1.855 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0022 12.017C22 6.484 17.522 2 12 2z" clipRule="evenodd" />
                </svg>
              </a>
              <a href="#" className="text-zinc-500 hover:text-white transition-colors">
                <span className="sr-only">Discord</span>
                <svg className="w-5 h-5" fill="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                   <path d="M20.317 4.37a19.791 19.791 0 00-4.885-1.515.074.074 0 00-.079.037c-.21.375-.444.864-.608 1.25a18.27 18.27 0 00-5.487 0 12.64 12.64 0 00-.617-1.25.077.077 0 00-.079-.037A19.736 19.736 0 003.677 4.37a.07.07 0 00-.032.027C.533 9.046-.32 13.58.099 18.057a.082.082 0 00.031.057 19.9 19.9 0 005.993 3.03.078.078 0 00.084-.028 14.09 14.09 0 001.226-1.994.076.076 0 00-.041-.106 13.107 13.107 0 01-1.872-.892.077.077 0 01-.008-.128 10.2 10.2 0 00.372-.292.074.074 0 01.077-.01c3.928 1.793 8.18 1.793 12.062 0a.074.074 0 01.078.01c.118.098.246.198.373.292a.077.077 0 01-.006.127 12.299 12.299 0 01-1.873.892.077.077 0 00-.041.107c.36.698.772 1.362 1.225 1.993a.076.076 0 00.084.028 19.839 19.839 0 006.002-3.03.077.077 0 00.032-.054c.5-5.177-.838-9.674-3.549-13.66a.061.061 0 00-.031-.03zM8.02 15.33c-1.183 0-2.157-1.085-2.157-2.419 0-1.333.956-2.419 2.157-2.419 1.21 0 2.176 1.096 2.157 2.42 0 1.333-.956 2.418-2.157 2.418zm7.975 0c-1.183 0-2.157-1.085-2.157-2.419 0-1.333.955-2.419 2.157-2.419 1.21 0 2.176 1.096 2.157 2.42 0 1.333-.946 2.418-2.157 2.418z" />
                </svg>
              </a>
            </div>
          </div>

          {/* Links Column 1 */}
          <div className="col-span-1">
            <h4 className="font-semibold text-white mb-4">Product</h4>
            <ul className="space-y-3 text-sm text-zinc-400">
              <li><a href="#" className="hover:text-rose-400 transition-colors">Download</a></li>
              <li><a href="#" className="hover:text-rose-400 transition-colors">Documentation</a></li>
              <li><a href="#" className="hover:text-rose-400 transition-colors">Changelog</a></li>
              <li><a href="#" className="hover:text-rose-400 transition-colors">Roadmap</a></li>
            </ul>
          </div>

          {/* Links Column 2 */}
          <div className="col-span-1">
            <h4 className="font-semibold text-white mb-4">Community</h4>
            <ul className="space-y-3 text-sm text-zinc-400">
              <li><a href="#" className="hover:text-rose-400 transition-colors">GitHub</a></li>
              <li><a href="#" className="hover:text-rose-400 transition-colors">Discord</a></li>
              <li><a href="#" className="hover:text-rose-400 transition-colors">Twitter</a></li>
            </ul>
          </div>
          
           {/* CTA Column */}
           <div className="col-span-1 flex flex-col items-start">
             <h4 className="font-semibold text-white mb-4">Stay Updated</h4>
             <p className="text-zinc-500 text-sm mb-4">
                Subscribe to our newsletter for the latest updates on local-first AI.
             </p>
             <div className="flex w-full space-x-2">
                <input 
                  type="email" 
                  placeholder="email@example.com" 
                  className="bg-zinc-900 border border-white/10 rounded-md px-3 py-2 text-sm text-white focus:outline-none focus:ring-1 focus:ring-rose-500 w-full"
                />
                <Button size="sm">ok</Button>
             </div>
           </div>

        </div>
        
        <div className="border-t border-white/5 pt-8 flex flex-col md:flex-row items-center justify-between text-xs text-zinc-600">
          <p>© 2024 Hippo Local AI. All rights reserved.</p>
          <div className="flex space-x-6 mt-4 md:mt-0">
            <a href="#" className="hover:text-zinc-400">Privacy Policy</a>
            <a href="#" className="hover:text-zinc-400">Terms of Service</a>
          </div>
        </div>
      </div>
    </footer>
  );
}
