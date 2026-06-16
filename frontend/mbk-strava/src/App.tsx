import { useState } from 'react'
import { BrowserRouter as Router, Routes, Route , Navigate , useLocation } from "react-router-dom";

import './App.css'
import Navbar from './components/navbar/nav';
import Login from './components/pages/login/login';
import Signup from './components/pages/signup/signup';
import Sessions from './components/pages/sessions/sessions';
import Statistics from './components/pages/statistics/statistics';
function App() {
 
  const showNavbar =
    location.pathname !== "/login" &&
    location.pathname !== "/signup";

  return (
    <>
      <Router>
        {showNavbar && <Navbar />}
        <div>
          <Routes>
            <Route path="/" element={<Navigate to="/login" />} />
            <Route path="/login" element={<Login/>} />
            <Route path="/signup" element={<Signup/>} />
            <Route path="/sessions" element={<Sessions />} />
            <Route path="/statistics" element={<Statistics/>} />
          </Routes>
        </div>
      </Router>
    </>
  )
}

export default App
