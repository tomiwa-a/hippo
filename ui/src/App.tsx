import { Layout } from './components/layout/Layout';

function App() {
  return (
    <Layout>
      <div className="min-h-screen flex flex-col items-center justify-center space-y-4">
        <h1 className="text-4xl font-bold font-sans text-white">
          Hippo <span className="text-rose-500">Daemon</span>
        </h1>
        <p className="font-mono text-zinc-400">
          Local-first semantic brain.
        </p>
        <div className="p-4 bg-zinc-900 border border-white/10 rounded-lg">
          <code className="text-rose-400 font-mono">$ hippo query "init"</code>
        </div>
      </div>
    </Layout>
  )
}

export default App
