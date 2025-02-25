// src/__tests__/LoginForm.test.js
import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import LoginForm from '../Components/Auth/LoginForm';
import { BrowserRouter} from 'react-router-dom';
import '@testing-library/jest-dom/extend-expect';

// Mocking environment variables
process.env.REACT_APP_API_BASE_URL = 'http://localhost:8080';

test('renders the login form with email and password fields', () => {
  render(<LoginForm onLoginSuccess={jest.fn()} />, { wrapper: BrowserRouter });
  const emailInput = screen.getByLabelText(/email/i);
  const passwordInput = screen.getByLabelText(/password/i);
  expect(emailInput).toBeInTheDocument();
  expect(passwordInput).toBeInTheDocument();
});


