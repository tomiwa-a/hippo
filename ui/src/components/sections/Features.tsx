import { motion } from 'framer-motion';

const features = [
  {
    title: "Fully Air-Gapped",
    description: "Designed for offline-first environments. No data ever leaves your local machine.",
    icon: (
      <svg className="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728A9 9 0 015.636 5.636m12.728 12.728L5.636 5.636" />
      </svg>
    ),
    colSpan: "col-span-12 md:col-span-4",
    bg: "bg-zinc-900",
  },
  {
    title: "Zero-Knowledge Indexing",
    description: "Your semantic graph is encrypted at rest. Only you hold the keys to decrypt.",
    icon: (
      <svg className="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
      </svg>
    ),
    colSpan: "col-span-12 md:col-span-4",
    bg: "bg-zinc-900",
  },
  {
    title: "< 150ms Latency",
    description: "Sub-perceptual retrieval. Faster than typing your query.",
    icon: (
      <svg className="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
      </svg>
    ),
    colSpan: "col-span-12 md:col-span-4",
    bg: "bg-zinc-900",
  }
];

export function Features() {
  return (
    <section id="features" className="py-24 bg-zinc-950/50">
      <div className="container px-6 mx-auto">
        <div className="mb-16 text-center max-w-2xl mx-auto">
          <h2 className="text-3xl font-bold tracking-tight text-white sm:text-4xl mb-4">
            Built for the <span className="text-rose-500">Secure Perimeter</span>.
          </h2>
          <p className="text-zinc-400 text-lg">
            Hippo is designed for environments where data privacy is not optional.
          </p>
        </div>

        <div className="grid grid-cols-12 gap-6">
          {features.map((feature, i) => (
            <motion.div
              key={i}
              className={`${feature.colSpan} relative group p-8 rounded-2xl border border-white/5 bg-zinc-900 overflow-hidden hover:border-white/10 transition-colors`}
              initial={{ opacity: 0, y: 20 }}
              whileInView={{ opacity: 1, y: 0 }}
              viewport={{ once: true }}
              transition={{ delay: i * 0.1 }}
            >
              <div className="absolute inset-0 bg-gradient-to-br from-rose-500/5 via-transparent to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-500" />
              
              <div className="relative z-10">
                <div className="w-12 h-12 rounded-lg bg-rose-500/10 flex items-center justify-center text-rose-500 mb-6 group-hover:bg-rose-500/20 transition-colors">
                  {feature.icon}
                </div>
                
                <h3 className="text-xl font-bold text-white mb-2">
                  {feature.title}
                </h3>
                
                <p className="text-zinc-400 leading-relaxed">
                  {feature.description}
                </p>
              </div>
            </motion.div>
          ))}
          
          {/* Main Large Card - OS Level Search Demo Context */}
          <motion.div
            className="col-span-12 relative group p-8 md:p-12 rounded-2xl border border-white/5 bg-zinc-900 overflow-hidden mt-6"
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            transition={{ delay: 0.4 }}
          >
             <div className="absolute top-0 right-0 p-12 opacity-10 pointer-events-none">
                <svg className="w-64 h-64 text-rose-500" viewBox="0 0 24 24" fill="currentColor">
                   <path d="M4 6a2 2 0 012-2h12a2 2 0 012 2v7a2 2 0 01-2 2h-1.586l-1.707 1.707a1 1 0 01-1.414 0L9.586 17H6a2 2 0 01-2-2V6z" />
                </svg>
             </div>
             
             <div className="relative z-10 max-w-2xl">
                 <div className="inline-flex items-center space-x-2 bg-rose-500/10 border border-rose-500/20 rounded-full px-3 py-1 mb-6">
                    <span className="text-xs font-semibold text-rose-400">Universal SDK</span>
                 </div>
                 
                 <h3 className="text-3xl md:text-4xl font-bold text-white mb-6">
                    One Brain. Any App.
                 </h3>
                 
                 <p className="text-zinc-400 text-lg mb-8">
                    Hippo exposes a simple local HTTP API and MCP server. Connect your terminal, your IDE, or your custom internal tools to the same semantic graph.
                 </p>
                 
                 {/* Mini code snippet decoration */}
                 <div className="bg-zinc-950/50 rounded-lg p-4 font-mono text-sm border border-white/5 text-zinc-300">
                    <span className="text-rose-400">import</span> hippo <span className="text-rose-400">from</span> <span className="text-green-400">'@hippo/sdk'</span>;
                    <br /><br />
                    <span className="text-purple-400">const</span> results = <span className="text-rose-400">await</span> hippo.<span className="text-blue-400">query</span>(<span className="text-green-400">"q4 financial report"</span>);
                 </div>
             </div>
          </motion.div>
        </div>
      </div>
    </section>
  );
}
