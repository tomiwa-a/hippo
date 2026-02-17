import { useState, useEffect } from 'react';
import { motion } from 'framer-motion';

const lines = [
  { text: "$ hippo query \"integrating auth module\"", delay: 0, type: 'command' },
  { text: "> Connecting to local brain... (Connected)", delay: 1000, type: 'system' },
  { text: "> Semantic search in progress... (14ms)", delay: 1500, type: 'system' },
  { text: "", delay: 1800, type: 'spacer' },
  { text: "[1] src/auth/providers.go (Score: 0.98)", delay: 2000, type: 'result' },
  { text: "    func SetupAuth(p *Provider) { ... }", delay: 2100, type: 'code' },
  { text: "", delay: 2300, type: 'spacer' },
  { text: "[2] docs/architecture/auth_flow.md (Score: 0.92)", delay: 2400, type: 'result' },
  { text: "    \"The authentication module uses OAuth2...\"", delay: 2500, type: 'code' },
  { text: "", delay: 2700, type: 'spacer' },
  { text: "> Query complete. 2 results found.", delay: 2900, type: 'success' },
];

export function TerminalDemo() {
  const [visibleLines, setVisibleLines] = useState<number>(0);

  useEffect(() => {
    // Reset animation loop
    const timeouts: ReturnType<typeof setTimeout>[] = [];
    
    const runAnimation = () => {
      setVisibleLines(0);
      let accumulatedDelay = 0;

      lines.forEach((line, index) => {
        accumulatedDelay = line.delay;
        const timeout = setTimeout(() => {
          setVisibleLines(index + 1);
        }, accumulatedDelay);
        timeouts.push(timeout);
      });

      // Restart loop
      const resetTimeout = setTimeout(runAnimation, accumulatedDelay + 4000);
      timeouts.push(resetTimeout);
    };

    runAnimation();

    return () => timeouts.forEach(clearTimeout);
  }, []);

  return (
    <section className="py-24 bg-zinc-950 flex flex-col items-center justify-center overflow-hidden">
      <div className="container px-6 mx-auto flex flex-col items-center">
        
        <div className="text-center max-w-2xl mb-12">
          <h2 className="text-3xl font-bold tracking-tight text-white sm:text-4xl mb-4">
            Terminal Intelligence.
          </h2>
          <p className="text-zinc-400 text-lg">
            Don't leave your shell. Query your entire digital life with natural language.
          </p>
        </div>

        {/* Terminal Window */}
        <motion.div 
          className="w-full max-w-3xl bg-[#1E1E1E] rounded-xl border border-white/10 shadow-2xl overflow-hidden font-mono text-sm md:text-base"
          initial={{ opacity: 0, y: 40 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          transition={{ duration: 0.8 }}
        >
          {/* Window Header */}
          <div className="bg-[#2D2D2D] px-4 py-2 flex items-center space-x-2 border-b border-black/20">
            <div className="w-3 h-3 rounded-full bg-[#FF5F56]"></div>
            <div className="w-3 h-3 rounded-full bg-[#FFBD2E]"></div>
            <div className="w-3 h-3 rounded-full bg-[#27C93F]"></div>
            <div className="flex-1 text-center text-zinc-400 text-xs">hippo-daemon — zsh</div>
          </div>

          {/* Terminal Content */}
          <div className="p-6 h-[400px] overflow-hidden flex flex-col space-y-2">
            {lines.slice(0, visibleLines).map((line, i) => (
              <motion.div 
                key={i}
                initial={{ opacity: 0, x: -10 }}
                animate={{ opacity: 1, x: 0 }}
                className={`${
                  line.type === 'command' ? 'text-rose-400 font-bold' :
                  line.type === 'system' ? 'text-zinc-500 italic' :
                  line.type === 'result' ? 'text-green-400 font-semibold' :
                  line.type === 'code' ? 'text-zinc-300 pl-4 border-l-2 border-zinc-700 ml-1' :
                  line.type === 'success' ? 'text-rose-400 mt-2' : ''
                }`}
              >
                {line.text}
              </motion.div>
            ))}
            
            {/* Blinking Cursor */}
            <motion.div 
              className="w-2.5 h-5 bg-zinc-500 inline-block"
              animate={{ opacity: [1, 0] }}
              transition={{ repeat: Infinity, duration: 0.8 }}
            />
          </div>
        </motion.div>

      </div>
    </section>
  );
}
