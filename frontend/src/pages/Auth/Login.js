import React, { useState } from 'react';
import LoginForm from '../../Components/Auth/LoginForm';

const Login = () => {
  const [token, setToken] = useState('');

  const handleLoginSuccess = (token) => {
    setToken(token);
    // Save the token in local storage or update the application state
    localStorage.setItem('token', token);
    // Redirect to a protected route or update the UI as needed
  };

  return (
    <div className="login-page">
      <LoginForm onLoginSuccess={handleLoginSuccess} />
    </div>
  );
};

export default Login;
