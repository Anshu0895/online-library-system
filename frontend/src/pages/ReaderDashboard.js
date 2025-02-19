import React, { useState, useEffect } from 'react';
import api from '../utils/api';
import '../Css/ReaderDashboard.css';

const ReaderDashboard = ({ token, user }) => {
  const [query, setQuery] = useState('');
  const [books, setBooks] = useState([]);
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  const handleSearch = async (e) => {
    e.preventDefault();
    setError('');
    setSuccess('');

    try {
      const response = await api.get('/books/search', {
        headers: {
          Authorization: `${token}`,
        },
        params: { title: query },
      });
      console.log("Search response:", response.data); // Debug log
      setBooks(response.data);
    } catch (err) {
      console.error("Search error:", err); // Debug log
      setError('Failed to search books');
    }
  };

  const handleBorrowRequest = async (isbn) => {
    
    try {
      // const parsedReaderID = parseInt(user.id, 10);
      // console.log("Sending ISBN and reader id:", isbn,parsedReaderID); // Debug log to ensure correct ISBN is sent
      const response = await api.post('/raise-request', { book_id: isbn}, {
        headers: {
          Authorization: `${token}`,
        },
      });
      console.log("Borrow request response:", response.data); // Debug log
      setSuccess('Borrow request raised successfully');
    } catch (err) {
      console.error("Borrow request error:", err); // Debug log
      setError('Failed to raise borrow request');
    }
  };

  return (
    <div className='reader-box'>
      <div className="dashboard-container">
        <h2>Reader Dashboard</h2>
        {error && <p className="error-message">{error}</p>}
        {success && <p className="success-message">{success}</p>}
        <form onSubmit={handleSearch}>
          <div className="form-group">
            <label htmlFor="query">Search Books:</label>
            <input
              type="text"
              id="query"
              value={query}
              onChange={(e) => setQuery(e.target.value)}
              required
            />
          </div>
          <button type="submit">Search</button>
        </form>
        <h3>Book Results</h3>
        <ul className="book-list"> 
          {books.map((book) => (
            <li key={book.isbn} className="book-item"> {/* Added class book-item */}
              <span className="book-title"><strong>Title:</strong> {book.title}</span>
              <span className="book-author"><strong>Author:</strong> {book.authors}</span>
              <span className="book-publisher"><strong>Publisher:</strong> {book.publisher}</span>
              <button className="borrow-button" onClick={() => handleBorrowRequest(book.isbn)}>Borrow</button>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default ReaderDashboard;
