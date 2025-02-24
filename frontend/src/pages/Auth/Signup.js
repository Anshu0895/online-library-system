import React, { useState } from "react";
import { useNavigate } from "react-router-dom"; // Import useNavigate
import SignupForm from "../../Components/Auth/SignupForm";

const Signup = () => {
  const [token, setToken] = useState("");
  const navigate = useNavigate(); // Initialize navigate

  const handleSignupSuccess = (token) => {
    setToken(token);
    localStorage.setItem("token", token);
    navigate("/login"); // Redirect to login after successful signup
  };

  return (
    <div>
      <SignupForm onSignupSuccess={handleSignupSuccess} />
    </div>
  );
};

export default Signup;