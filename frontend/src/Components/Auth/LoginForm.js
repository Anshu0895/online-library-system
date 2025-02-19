import React, { useState } from 'react';
import '../../Css/LoginForm.css';
import { Link } from "react-router-dom";
import api from '../../utils/api';
import { toast } from 'react-toastify';
import { useNavigate } from 'react-router-dom';

const LoginForm = ({ onLoginSuccess }) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');

    // Clear the old token before login
    localStorage.removeItem('token');
    localStorage.removeItem('role');

    try {
      const response = await api.post('/login', { email, password });
      console.log("Response data:", response.data); 
      const { token, role } = response.data;
      

      

      // Call the onLoginSuccess callback with the token and role
      onLoginSuccess(token, role);
      
      // Save token and role in local storage
      localStorage.setItem('token', token);
      localStorage.setItem('role', role);

      toast.success('Login successful!', {
        position: "top-center",
        style: { backgroundColor: 'black', color: 'white' },
      });

      // Navigate based on user role
      switch (role) {
        case 'Owner':
          navigate('/libraries');
          break;
        case 'Admin':
          navigate('/admin');
          break;
        case 'Reader':
          navigate('/reader');
          break;
        default:
          navigate('/');
          break;
      }

    } catch (err) {
      console.error("Login error:", err);
      setError('Invalid email or password');
      toast.error('Invalid email or password. Please try again.', {
        position: "top-center",
        style: { backgroundColor: 'black', color: 'white' },
      });
    }
  };

  return (
    <div className='login-box'> 
      <div className="login-form-container">
        <h2>Login</h2>
        {error && <p className="error-message">{error}</p>}
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="email">Email:</label>
            <input
              type="email"
              id="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />
          </div>
          <div className="form-group">
            <label htmlFor="password">Password:</label>
            <input
              type="password"
              id="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>
          <button type="submit">Login</button>
          <p>Don't have an account? <Link to="/signup">Signup here</Link></p>
        </form>
      </div>
    </div>
  );
};

export default LoginForm;
