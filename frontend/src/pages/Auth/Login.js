import React, { useContext } from 'react';
import { useNavigate } from 'react-router-dom';
import AuthContext from '../../contexts/AuthContext'; // Ensure this import is correct
import LoginForm from '../../Components/Auth/LoginForm';

const Login = () => {
  const { setAuthData } = useContext(AuthContext); // Check this line
  const navigate = useNavigate();

  const handleLoginSuccess = (token) => {
    // Store token in local storage or context
    localStorage.setItem('token', token);
    setAuthData({ token });

    // Redirect to the dashboard or home page
    navigate('/dashboard');
  };

  return (
    <div className="login-page">
      <LoginForm onLoginSuccess={handleLoginSuccess} />
    </div>
  );
};

export default Login;
