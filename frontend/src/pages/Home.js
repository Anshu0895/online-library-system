import React from "react";
import "../Css/HomePage.css"; 
import right2 from "../assests/right2.png"; 
import rightback from "../assests/rightback.png";
import { Link } from "react-router-dom";

const HomePage = () => {
  return (
    <div className="container">
      <div className="left-side">
        <h1>Welcome to BookStack Library</h1>
        <p>
        BookStack Library is your go-to destination for an extensive collection of books across various genres.
        Whether you're looking for fiction, non-fiction,
        or academic resources, we have something for everyone.
        Join us and embark on a literary adventure!
        </p>
        <div className="get-issued">
      <Link to="/login" className="get-button">Get issued</Link>
      </div>
      </div>
      
      <div className="right-side">
        <img src={right2} alt="Stack of Books 1" className="book-stack" />
        <img src={rightback} alt="Stack of Books 2" className="book-stack" />
      </div>
      
    </div>
  );
};

export default HomePage;
