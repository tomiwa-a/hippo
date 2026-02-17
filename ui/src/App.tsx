import { Layout } from './components/layout/Layout';
import { Hero } from './components/sections/Hero';
import { Features } from './components/sections/Features';
import { TerminalDemo } from './components/sections/TerminalDemo';

function App() {
  return (
    <Layout>
      <Hero />
      <Features />
      <TerminalDemo />
    </Layout>
  )
}

export default App
