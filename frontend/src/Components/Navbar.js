import React from "react";
import {Link} from "react-router-dom";
import logo from "../assests/stack-of-books.png"
import "../Css/Navbar.css";

const Navbar = ({token,handleLogout}) => {
  return (
    <nav className="navbar">
      <div className="logo-container">
      <Link to="/">
        <img 
          src={logo}
          alt="Library Logo" 
          className="logo" 
          
        />
         </Link>
        <h1 className="title">BookStack</h1>
       
      </div>
      <div className="search-container">
        <input 
          type="text"
          placeholder="Search books..."
          className="search-input"
        />
      </div>
      <div>
      {token ? (
          <button onClick={handleLogout} className="nav-button">Logout</button>
        ) : (
          <>
            <Link to="/login" className="nav-button">Login</Link>
            <Link to="/signup" className="nav-button">Sign Up</Link>
          </>
        )}
      </div>
      
    </nav>
  );
};

export default Navbar;