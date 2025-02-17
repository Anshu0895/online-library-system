import React, { useState, useEffect } from 'react';
import axios from 'axios';
import '../Css/ReaderDashboard.css';

const ReaderDashboard = ({ token }) => {
  const [books, setBooks] = useState([]);
  const [searchQuery, setSearchQuery] = useState('');
  const [filteredBooks, setFilteredBooks] = useState([]);

  useEffect(() => {
    const fetchBooks = async () => {
      try {
        const response = await axios.get('/books', {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        setBooks(response.data);
      } catch (err) {
        console.error(err);
      }
    };
    fetchBooks();
  }, [token]);

  const handleSearch = (e) => {
    setSearchQuery(e.target.value);
    if (searchQuery) {
      setFilteredBooks(books.filter(book =>
        book.title.toLowerCase().includes(searchQuery.toLowerCase()) ||
        book.author.toLowerCase().includes(searchQuery.toLowerCase()) ||
        book.publisher.toLowerCase().includes(searchQuery.toLowerCase())
      ));
    } else {
      setFilteredBooks([]);
    }
  };

  const handleBorrowRequest = (bookId) => {
    // Handle borrow request logic here
    console.log(`Borrow request for book ID: ${bookId}`);
  };

  return (
    <div className='reader-box'>
    <div className="dashboard-container">
      <h2>Reader Dashboard</h2>
      <div className="search-bar">
        <input
          type="text"
          placeholder="Search by title, author, or publisher"
          value={searchQuery}
          onChange={handleSearch}
        />
      </div>
      <div className="book-list">
        {searchQuery && filteredBooks.length > 0 ? (
          filteredBooks.map((book) => (
            <div key={book.id} className="book-card">
              <h3>{book.title}</h3>
              <p>Author: {book.author}</p>
              <p>Publisher: {book.publisher}</p>
              <p>Status: {book.isAvailable ? 'Available' : 'Not Available'}</p>
              <button onClick={() => handleBorrowRequest(book.id)}>Borrow</button>
            </div>
          ))
        ) : (
          books.map((book) => (
            <div key={book.id} className="book-card">
              <h3>{book.title}</h3>
              <p>Author: {book.author}</p>
              <p>Publisher: {book.publisher}</p>
              <p>Status: {book.isAvailable ? 'Available' : 'Not Available'}</p>
              <button onClick={() => handleBorrowRequest(book.id)}>Borrow</button>
            </div>
          ))
        )}
      </div>
    </div>
    </div>
  );
};

export default ReaderDashboard;
