
import React, { useState, useEffect } from 'react';
import api from '../utils/api';
import '../Css/ReaderDashboard.css';
import { toast } from 'react-toastify';
const ReaderDashboard = ({ token }) => {
  const [books, setBooks] = useState([]);
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');
  const [showSearchPopup, setShowSearchPopup] = useState(false);

  // Search functionality states
  const [query, setQuery] = useState('');
  const [searchType, setSearchType] = useState('title');

  // Fetch all books on load
  useEffect(() => {
    fetchBooks();
  }, [token]);

  const fetchBooks = async () => {
    try {
      const response = await api.get('/books', {
        headers: { Authorization: `${token}` },
      });
      setBooks(response.data);
    } catch (err) {
      console.error("Error fetching books:", err);
      setError('Failed to fetch books');
    }
  };

  // Handle search function
  const handleSearch = async (e) => {
    e.preventDefault();
    setError('');
    setSuccess('');

    try {
      const response = await api.get('/books/search', {
        headers: { Authorization: `${token}` },
        params: { [searchType]: query },
      });
      setBooks(response.data); // Replace current books with search results
      setShowSearchPopup(false); // Close search popup after searching
    } catch (err) {
      console.error("Search error:", err);
      setError('Failed to search books');
      toast.error('Failed to search books', {
                    position: "top-center",
                    style: { backgroundColor: 'black', color: 'white' },
                  });
    }
  };

  const handleBorrowRequest = async (isbn) => {
    try {
      const readerId = localStorage.getItem("id"); // Get reader ID
      if (!readerId) {
        setError("User ID not found. Please log in again.");
        return;
      }
  
      await api.post(
        "/raise-request",
        { book_id: isbn, reader_id:Number(readerId) }, // Include reader ID
        { headers: { Authorization: `${token}` } }
      );
  
      setSuccess("Borrow request raised successfully");
    } catch (err) {
      console.error("Borrow request error:", err);
      setError("Failed to raise borrow request");
    }
  };
  

  return (
    <div className="reader-box">
      <div className="reader-dashboard-container">
        <h2>Reader Dashboard</h2>
        {error && <p className="error-message">{error}</p>}
        {success && <p className="success-message">{success}</p>}

        {/* Search Button */}
        <button className="search-button" onClick={() => setShowSearchPopup(true)}>Search Books</button>

        {/* Book Cards */}
        <div className="book-cards-container">
          {books.length === 0 ? (
            <p>No books found</p>
          ) : (
            books.map((book) => (
              <div key={book.isbn} className="book-card">
                <h3>{book.title}</h3>
                <p><strong>Author:</strong> {book.authors}</p>
                <p><strong>Publisher:</strong> {book.publisher}</p>
                <p className={`book-availability ${book.available_copies > 0 ? 'available' : 'not-available'}`}>
                  {book.available_copies > 0 ? "Available" : "Not Available"}
                </p>
                <button
                  className="borrow-button"
                  onClick={() => handleBorrowRequest(book.isbn)}
                  disabled={book.available_copies === 0}
                >
                  Borrow
                </button>
              </div>
            ))
          )}
        </div>

        {showSearchPopup && (
  <div className="search-overlay">
    <div className="search-popup">
      <div className="popup-content">
        <button className="close-button" onClick={() => setShowSearchPopup(false)}>X</button>
        <h3>Search Books</h3>
        <form onSubmit={handleSearch}>
          <div className="form-group">
            <label htmlFor="searchType">Search By:</label>
            <select id="searchType" value={searchType} onChange={(e) => setSearchType(e.target.value)}>
              <option value="title">Title</option>
              <option value="author">Author</option>
              <option value="publisher">Publisher</option>
            </select>
          </div>

          <div className="form-group">
            <label htmlFor="query">Enter {searchType}:</label>
            <input
              type="text"
              id="query"
              placeholder={`Search by ${searchType}`}
              value={query}
              onChange={(e) => setQuery(e.target.value)}
              required
            />
          </div>

          <button type="submit">Search</button>
        </form>
      </div>
    </div>
  </div>
)}

      </div>
    </div>
  );
};

export default ReaderDashboard;
