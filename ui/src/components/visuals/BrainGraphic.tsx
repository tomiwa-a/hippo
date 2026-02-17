import { motion } from 'framer-motion';

export function BrainGraphic() {
  // Simple node positions for a "brain-like" graph
  const nodes = [
    { x: 400, y: 300, r: 6 }, // Center
    { x: 350, y: 250, r: 4 },
    { x: 450, y: 250, r: 4 },
    { x: 350, y: 350, r: 4 },
    { x: 450, y: 350, r: 4 },
    { x: 300, y: 300, r: 3 },
    { x: 500, y: 300, r: 3 },
    { x: 400, y: 200, r: 5 }, // Top
    { x: 400, y: 400, r: 5 }, // Bottom
    // Extra scattered nodes
    { x: 280, y: 240, r: 2 },
    { x: 520, y: 240, r: 2 },
    { x: 280, y: 360, r: 2 },
    { x: 520, y: 360, r: 2 },
  ];

  const connections = [
    [0, 1], [0, 2], [0, 3], [0, 4], // Center star
    [1, 7], [2, 7], // Top connections
    [3, 8], [4, 8], // Bottom connections
    [1, 5], [3, 5], // Left side
    [2, 6], [4, 6], // Right side
    [5, 9], [6, 10], // Outer left/right
    [5, 11], [6, 12],
  ];

  return (
    <div className="relative w-full h-[400px] md:h-[600px] flex items-center justify-center opacity-60 md:opacity-100">
      <svg viewBox="0 0 800 600" className="w-full h-full max-w-[800px]">
        {/* Glow Effects */}
        <defs>
          <filter id="glow" x="-50%" y="-50%" width="200%" height="200%">
            <feGaussianBlur stdDeviation="4" result="coloredBlur" />
            <feMerge>
              <feMergeNode in="coloredBlur" />
              <feMergeNode in="SourceGraphic" />
            </feMerge>
          </filter>
        </defs>

        {/* Connections */}
        {connections.map(([start, end], i) => (
          <motion.line
            key={`conn-${i}`}
            x1={nodes[start].x}
            y1={nodes[start].y}
            x2={nodes[end].x}
            y2={nodes[end].y}
            stroke="#F43F5E" // Rose-500
            strokeWidth="1"
            strokeOpacity="0.2"
            initial={{ pathLength: 0, opacity: 0 }}
            animate={{ pathLength: 1, opacity: 0.2 }}
            transition={{ duration: 1.5, delay: i * 0.05, ease: "easeInOut" }}
          />
        ))}

        {/* Nodes */}
        {nodes.map((node, i) => (
          <motion.circle
            key={`node-${i}`}
            cx={node.x}
            cy={node.y}
            r={node.r}
            fill={i === 0 ? "#F43F5E" : "#18181B"} // Center is Rose, others Zinc-900
            stroke="#F43F5E"
            strokeWidth="2"
            filter={i === 0 ? "url(#glow)" : ""}
            initial={{ scale: 0, opacity: 0 }}
            animate={{ scale: 1, opacity: 1 }}
            transition={{ duration: 0.5, delay: 1 + i * 0.05 }}
          />
        ))}

        {/* Pulsing signal on lines (simulating data transfer) */}
        <motion.circle
          r="3"
          fill="#fff"
          initial={{ opacity: 0 }}
          animate={{ 
            offsetDistance: ["0%", "100%"], 
            opacity: [0, 1, 0] 
          }}
          // Note: Framer Motion SVG path animation is simpler for path drawing, 
          // moving dots along lines usually requires CSS offset-path or custom logic.
          // For simplicity/robustness, we'll pulse the center node instead.
        />
        
        {/* Center Heartbeat */}
        <motion.circle
          cx={nodes[0].x}
          cy={nodes[0].y}
          r={nodes[0].r * 3}
          stroke="#F43F5E"
          strokeWidth="1"
          fill="none"
          initial={{ scale: 0.5, opacity: 0 }}
          animate={{ scale: 2.5, opacity: 0 }}
          transition={{ duration: 2, repeat: Infinity, ease: "easeOut" }}
        />
      </svg>
      
      {/* Background Gradient Mesh for Depth */}
      <div className="absolute inset-0 bg-rose-500/5 blur-[100px] rounded-full pointer-events-none" />
    </div>
  );
}
