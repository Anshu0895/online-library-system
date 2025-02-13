import React from 'react';
import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router-dom';
import Login from './pages/Auth/Login';
import Signup from './pages/Auth/Signup';
import { AuthProvider } from './contexts/AuthContext';
import './App.css';

const App = () => {
  return (
    <AuthProvider>
      <Router>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/signup" element={<Signup />} />
          <Route path="*" element={<Navigate to="/login" />} /> {/* Redirect all other paths to login */}
        </Routes>
      </Router>
    </AuthProvider>
  );
};

export default App;
