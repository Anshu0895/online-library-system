import React, { useState } from 'react';
import { toast } from 'react-toastify'
import '../../Css/SignupForm.css';
import { Link } from "react-router-dom";
import api from '../../utils/api';

const SignupForm = ({ onSignupSuccess }) => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [contactNumber, setContactNumber] = useState('');
  const [role, setRole] = useState('');
  const [libID, setLibID] = useState('');
  const [error, setError] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');

    try {
      const response = await api.post('/signup', {
        name,
        email,
        password,
        contactNumber,
        role,
        libID,
      });
      const { token } = response.data;

      // Call the onSignupSuccess callback with the token
      onSignupSuccess(token);
      toast.success('Signup successful!',{
        position: "top-center",
        style: { backgroundColor: 'black', color: 'white' },
      });
    } catch (err) {
      setError('Failed to sign up');

      toast.error('Failed to sign up. Please try again.',{
        position: "top-center",
        style: { backgroundColor: 'black', color: 'white' },
      });
    }
  };

  return (
    <div className='signup-box'>
    <div className="signup-form-container">
      <h2>Sign-up</h2>
      {error && <p className="error-message">{error}</p>}
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="name">Name:</label>
          <input
            type="text"
            id="name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
        </div>
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
        <div className="form-group">
          <label htmlFor="contactNumber">Contact Number:</label>
          <input
            type="text"
            id="contactNumber"
            value={contactNumber}
            onChange={(e) => setContactNumber(e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label htmlFor="role">Role:</label>
          <select
            id="role"
            value={role}
            onChange={(e) => setRole(e.target.value)}
            required
          >
            <option value="">Select Role</option>
            <option value="Owner">Library Owner</option>
            <option value="Admin">Library Admin</option>
            <option value="Reader">Reader</option>
          </select>
        </div>
        <div className="form-group">
          <label htmlFor="libID">Library ID:</label>
          <input
            type="text"
            id="libID"
            value={libID}
            onChange={(e) => setLibID(e.target.value)}
            required
          />
        </div>
        <button type="submit">Signup</button>
        <p>Already have an account? <Link  to="/login">Login here</Link></p>
      </form>
    </div>
    </div>
  );
};

export default SignupForm;
