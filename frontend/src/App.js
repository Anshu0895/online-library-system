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
import AdminDashboard from './pages/AdminDashboard';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import api from './utils/api';

const App = () => {
  const [token, setToken] = useState(localStorage.getItem('token'));
  const [user, setUser] = useState(null);
  useEffect(() => {
    const storedToken = localStorage.getItem('token');
    if (storedToken) {
      setToken(storedToken);
    }
  }, []);
  const fetchUserData = async (token) => {
    try {
      const response = await api.get('/user', {
        headers: {
          Authorization: `${token}`,
        },
      });
      setUser(response.data);
    } catch (err) {
      console.error("Failed to fetch user data", err);
    }
  };

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
          <Route path="/reader" element={<ReaderDashboard token={token} user={user}/>} /> 
          <Route path="/admin" element={<AdminDashboard token={token}/>} />
        
        </Routes>
        <ToastContainer />
      </Router>
      
  );
};

export default App;