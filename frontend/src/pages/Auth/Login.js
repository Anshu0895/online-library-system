import React, { useState } from 'react';
import LoginForm from '../../Components/Auth/LoginForm';

const Login = ({ onLoginSuccess }) => {
  const [token, setToken] = useState('');
  const [role, setRole] = useState('');
  const[id,setid]=useState('');
  const handleLoginSuccess = (token, role, id) => {
    setToken(token);
    setRole(role);
    setid(id)
    // Store user ID in localStorage
    localStorage.setItem("token", token);
    localStorage.setItem("role", role);
    localStorage.setItem("id", id); 
  
    onLoginSuccess(token);
  };

  return (
    <div className="login-page">
      <LoginForm onLoginSuccess={handleLoginSuccess} />
    </div>
  );
};

export default Login;
