import { motion } from 'framer-motion';
import { Button } from '../ui/Button';
import { BrainGraphic } from '../visuals/BrainGraphic';

export function Hero() {
  return (
    <section className="relative min-h-[90vh] flex flex-col items-center justify-center pt-20 overflow-hidden">
      <div className="container px-6 mx-auto grid grid-cols-1 lg:grid-cols-2 gap-12 items-center">
        
        {/* Text Content */}
        <motion.div 
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.6 }}
          className="text-center lg:text-left z-10"
        >
          <div className="inline-flex items-center space-x-2 bg-white/5 border border-white/10 rounded-full px-3 py-1 mb-6">
            <span className="flex w-2 h-2 rounded-full bg-rose-500 animate-pulse"></span>
            <span className="text-xs font-medium text-zinc-300">Phase 1: Alpha Available</span>
          </div>
          
          <h1 className="text-5xl md:text-7xl font-bold tracking-tight text-white mb-6 leading-[1.1]">
            The Memory of <br />
            <span className="text-transparent bg-clip-text bg-gradient-to-r from-rose-500 to-rose-300">
              Your Machine.
            </span>
          </h1>
          
          <p className="text-lg md:text-xl text-zinc-400 mb-8 max-w-lg mx-auto lg:mx-0 leading-relaxed">
            A local-first semantic brain for hackers, lawyers, and doctors. 
            Index your life without giving it to the cloud.
          </p>
          
          <div className="flex flex-col sm:flex-row items-center justify-center lg:justify-start space-y-4 sm:space-y-0 sm:space-x-4">
            <Button size="md" className="w-full sm:w-auto text-lg h-12 px-8">
              Download Daemon
            </Button>
            <Button variant="ghost" size="md" className="w-full sm:w-auto h-12 px-8 border border-white/10">
              Read the Docs
            </Button>
          </div>

          <div className="mt-8 flex items-center justify-center lg:justify-start space-x-6 text-sm text-zinc-500 font-mono">
            <span>MacOS 14+</span>
            <span>•</span>
            <span>Linux (Beta)</span>
            <span>•</span>
            <span>Windows (Soon)</span>
          </div>
        </motion.div>

        {/* Visual Content */}
        <motion.div 
          initial={{ opacity: 0, scale: 0.9 }}
          animate={{ opacity: 1, scale: 1 }}
          transition={{ duration: 0.8, delay: 0.2 }}
          className="relative lg:h-[700px] flex items-center justify-center"
        >
          <div className="absolute inset-0 bg-gradient-to-tr from-rose-500/10 to-transparent blur-3xl rounded-full" />
          <BrainGraphic />
        </motion.div>
      </div>
      
      {/* Scroll Indicator */}
      <motion.div 
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        transition={{ delay: 1, duration: 1 }}
        className="absolute bottom-10 left-1/2 -translate-x-1/2 flex flex-col items-center space-y-2"
      >
        <span className="text-xs text-zinc-600 font-mono uppercase tracking-widest">Scroll</span>
        <div className="w-[1px] h-12 bg-gradient-to-b from-zinc-800 to-transparent" />
      </motion.div>
    </section>
  );
}
