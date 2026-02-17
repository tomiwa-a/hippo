import { motion } from 'framer-motion';

const useCases = [
  {
    title: "Legal & Compliance",
    description: "Index thousands of case files, depositions, and contracts. Search for concepts, not just keywords. All data stays on your air-gapped machine.",
    icon: (
      <svg className="w-8 h-8 text-rose-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M3 6l3 1m0 0l-3 9a5.002 5.002 0 006.001 0M6 7l3 9M6 7l6-2m6 2l3-1m-3 1l-3 9a5.002 5.002 0 006.001 0M18 7l3 9m-3-9l-6-2m0-2v2m0 16V5m0 16H9m3 0h3" />
      </svg>
    ),
  },
  {
    title: "Medical Research",
    description: "Instantly retrieve patient history from disparate EMR exports and PDF reports. Maintain absolute HIPAA compliance by never uploading PHI to the cloud.",
    icon: (
      <svg className="w-8 h-8 text-rose-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M19.428 15.428a2 2 0 00-1.022-.547l-2.384-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z" />
      </svg>
    ),
  },
  {
    title: "Engineering",
    description: "Search your entire local codebase semantically. 'Where is the auth retry logic?' yields the exact function, even if you forgot the variable names.",
    icon: (
      <svg className="w-8 h-8 text-rose-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4" />
      </svg>
    ),
  }
];

export function UseCases() {
  return (
    <section id="use-cases" className="py-24 bg-zinc-950 relative overflow-hidden">
      {/* Background Decor */}
      <div className="absolute top-0 inset-x-0 h-px bg-gradient-to-r from-transparent via-rose-500/20 to-transparent" />
      
      <div className="container px-6 mx-auto relative z-10">
        <div className="mb-16 md:text-center max-w-3xl mx-auto">
          <h2 className="text-3xl font-bold tracking-tight text-white sm:text-4xl mb-6">
            Powering the <span className="text-rose-500">Privileged World</span>.
          </h2>
          <p className="text-zinc-400 text-lg">
            Hippo is built for industries where data sovereignty is non-negotiable.
            Your secrets never leave your silicon.
          </p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
          {useCases.map((useCase, i) => (
            <motion.div
              key={i}
              initial={{ opacity: 0, y: 20 }}
              whileInView={{ opacity: 1, y: 0 }}
              viewport={{ once: true }}
              transition={{ delay: i * 0.1 }}
              className="relative p-8 group transition-colors"
            >
              {/* Unique 'Corner Brackets' Effect */}
              <div className="absolute top-0 left-0 w-8 h-8 border-t-2 border-l-2 border-zinc-800 group-hover:border-rose-500 transition-colors duration-300" />
              <div className="absolute top-0 right-0 w-8 h-8 border-t-2 border-r-2 border-zinc-800 group-hover:border-rose-500 transition-colors duration-300" />
              <div className="absolute bottom-0 left-0 w-8 h-8 border-b-2 border-l-2 border-zinc-800 group-hover:border-rose-500 transition-colors duration-300" />
              <div className="absolute bottom-0 right-0 w-8 h-8 border-b-2 border-r-2 border-zinc-800 group-hover:border-rose-500 transition-colors duration-300" />
              
              {/* Hover Background Hint */}
              <div className="absolute inset-4 bg-zinc-900/0 group-hover:bg-zinc-900/40 transition-colors duration-300 -z-10" />

              <div className="mb-6 relative">
                 <div className="p-3 w-fit transition-transform duration-300 group-hover:scale-110 group-hover:text-rose-400">
                    {useCase.icon}
                 </div>
              </div>
              
              <h3 className="text-xl font-bold text-white mb-3 group-hover:text-rose-500 transition-colors">
                {useCase.title}
              </h3>
              <p className="text-zinc-400 leading-relaxed group-hover:text-zinc-300 transition-colors">
                {useCase.description}
              </p>
            </motion.div>
          ))}
        </div>
      </div>
    </section>
  );
}
