import React, { useState } from 'react';
import axios from 'axios';
import './LoginForm.css';
const LoginForm = ({ onLoginSuccess }) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');

    try {
      const response = await axios.post('/login', { email, password });
      const { token } = response.data;
      
      // Call the onLoginSuccess callback with the token
      onLoginSuccess(token);
    } catch (err) {
      setError('Invalid email or password');
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
      </form>
    </div>
    </div>
  );
};

export default LoginForm;
