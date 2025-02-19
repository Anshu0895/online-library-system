import React, { useState } from 'react';
import LoginForm from '../../Components/Auth/LoginForm';

const Login = () => {
  const [token, setToken] = useState('');
  const [role, setRole] = useState('');

  const handleLoginSuccess = (token, role) => {

     
    setToken(token);
    setRole(role);
    // Save the token and role in local storage or update the application state
    localStorage.setItem('token', token);
    localStorage.setItem('role', role);
    // Redirect to a protected route or update the UI as needed

    
  };

  return (
    <div className="login-page">
      <LoginForm onLoginSuccess={handleLoginSuccess} />
    </div>
  );
};

export default Login;
