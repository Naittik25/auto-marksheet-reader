import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Navbar from './components/Navbar';

// Import Pages
import Home from './pages/Home';
import Scan from './pages/Scan';
import Results from './pages/Results';
import Export from './pages/Export'; // <--- NEW
import Extract from './pages/Extract'; // <--- NEW

import './styles/globals.css';

function App() {
  return (
    <Router>
      <div className="app-main">
        <Navbar />
        <div className="content-area">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/scan" element={<Scan />} />
            <Route path="/extract" element={<Extract />} /> {/* <--- NEW */}
            <Route path="/results" element={<Results />} />
            <Route path="/export" element={<Export />} />   {/* <--- NEW */}
          </Routes>
        </div>
      </div>
    </Router>
  );
}

export default App;