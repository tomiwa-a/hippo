import { Layout } from './components/layout/Layout';
import { Hero } from './components/sections/Hero';
import { Features } from './components/sections/Features';
import { UseCases } from './components/sections/UseCases';
import { TerminalDemo } from './components/sections/TerminalDemo';
import { Footer } from './components/layout/Footer';

function App() {
  return (
    <Layout>
      <Hero />
      <Features />
      <UseCases />
      <TerminalDemo />
      <Footer />
    </Layout>
  )
}

export default App
