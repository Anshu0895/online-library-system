import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Login from './pages/Auth/Login';
import Signup from './pages/Auth/Signup';
import Home from './pages/Home'
import Navbar from './Components/Navbar'
import { useState,useEffect } from 'react';
import './App.css';
import OwnerDashboard from './pages/OwnerDashboard';
import ReaderDashboard from './pages/ReaderDashboard';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
const App = () => {
  const [token, setToken] = useState(localStorage.getItem('token'));

  useEffect(() => {
    const storedToken = localStorage.getItem('token');
    if (storedToken) {
      setToken(storedToken);
    }
  }, []);

  const handleLoginSuccess = (token) => {
    setToken(token);
    localStorage.setItem('token', token);
  };
  return (
    
    
      <Router>
      <Navbar/>
        <Routes>
      
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login onLoginSuccess={handleLoginSuccess} />} />
          <Route path="/signup" element={<Signup />} />
          <Route path="/libraries" element={<OwnerDashboard token={token} />} />
          <Route path="/books" element={<ReaderDashboard/>} /> 
        </Routes>
        <ToastContainer />
      </Router>
      
  );
};

export default App;