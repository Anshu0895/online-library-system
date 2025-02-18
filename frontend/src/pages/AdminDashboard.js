import React, { useState, useEffect } from 'react';
import api from '../utils/api';
import '../Css/AdminDashboard.css';

const AdminDashboard = ({ token }) => {
  const [books, setBooks] = useState([]);
  const [requests, setRequests] = useState([]);
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');
  const [isAddBookFormVisible, setAddBookFormVisible] = useState(false);
  const [isRequestsVisible, setRequestsVisible] = useState(false);
  const [isUpdateBookFormVisible, setUpdateBookFormVisible] = useState(false);
  const [currentBook, setCurrentBook] = useState(null);
  const [newBook, setNewBook] = useState({
    isbn: '',
    title: '',
    authors: '',
    publisher: '',
    totalCopies: 0,
  });

  const fetchBooks = async () => {
    try {
      const response = await api.get('/book', {
        headers: {
          Authorization: `${token}`,
        },
      });
      setBooks(response.data);
    } catch (err) {
      setError('Failed to fetch books');
    }
  };

  const fetchRequests = async () => {
    try {
      const response = await api.get('/requests', {
        headers: {
          Authorization: `${token}`,
        },
      });
      setRequests(response.data);
    } catch (err) {
      setError('Failed to fetch requests');
    }
  };

  const handleAddBook = async (e) => {
    e.preventDefault();
    try {
      const response = await api.post('/books', newBook, {
        headers: {
          Authorization: `${token}`,
        },
      });
      setSuccess('Book added successfully');
      fetchBooks(); // Refresh books list
    } catch (err) {
      setError('Failed to add book');
    }
  };

  const handleDeleteBook = async (isbn) => {
    try {
      const response = await api.delete(`/books/${isbn}`, {
        headers: {
          Authorization: `${token}`,
        },
      });
      setSuccess('Book deleted successfully');
      fetchBooks(); // Refresh books list
    } catch (err) {
      setError('Failed to delete book');
    }
  };

  const handleUpdateBook = async (e) => {
    e.preventDefault();
    try {
      const response = await api.put(`/books/${currentBook.isbn}`, currentBook, {
        headers: {
          Authorization: `${token}`,
        },
      });
      setSuccess('Book updated successfully');
      fetchBooks(); // Refresh books list
      setUpdateBookFormVisible(false);
    } catch (err) {
      setError('Failed to update book');
    }
  };

  const handleEditClick = (book) => {
    setCurrentBook(book);
    setUpdateBookFormVisible(true);
  };

  const handleAcceptRequest = async (requestId) => {
    try {
      const response = await api.put(`/requests/${requestId}/approve`, {}, {
        headers: {
          Authorization: `${token}`,
        },
      });
      setSuccess('Request approved successfully');
      fetchRequests(); // Refresh requests list
    } catch (err) {
      setError('Failed to approve request');
    }
  };

  const handleRejectRequest = async (requestId) => {
    try {
      const response = await api.put(`/requests/${requestId}/reject`, {}, {
        headers: {
          Authorization: `${token}`,
        },
      });
      setSuccess('Request rejected successfully');
      fetchRequests(); // Refresh requests list
    } catch (err) {
      setError('Failed to reject request');
    }
  };

  useEffect(() => {
    fetchBooks();
    fetchRequests();
  }, []);

  return (
    <div className='admin-dashboard'>
      <h2>Admin Dashboard</h2>
      {error && <p className="error-message">{error}</p>}
      {success && <p className="success-message">{success}</p>}

      <button className='add-book' onClick={() => setAddBookFormVisible(!isAddBookFormVisible)}>
        {isAddBookFormVisible ? 'Hide Add Book Form' : 'Show Add Book Form'}
      </button>
      {isAddBookFormVisible && (
        <form onSubmit={handleAddBook}>
          <div className="form-group">
            <label htmlFor="isbn">ISBN:</label>
            <input
              type="text"
              id="isbn"
              value={newBook.isbn}
              onChange={(e) => setNewBook({ ...newBook, isbn: e.target.value })}
              required
            />
          </div>
          <div className="form-group">
            <label htmlFor="title">Title:</label>
            <input
              type="text"
              id="title"
              value={newBook.title}
              onChange={(e) => setNewBook({ ...newBook, title: e.target.value })}
              required
            />
          </div>
          <div className="form-group">
            <label htmlFor="authors">Authors:</label>
            <input
              type="text"
              id="authors"
              value={newBook.authors}
              onChange={(e) => setNewBook({ ...newBook, authors: e.target.value })}
              required
            />
          </div>
          <div className="form-group">
            <label htmlFor="publisher">Publisher:</label>
            <input
              type="text"
              id="publisher"
              value={newBook.publisher}
              onChange={(e) => setNewBook({ ...newBook, publisher: e.target.value })}
              required
            />
          </div>
          <div className="form-group">
            <label htmlFor="totalCopies">Total Copies:</label>
            <input
              type="number"
              id="totalCopies"
              value={newBook.totalCopies}
              onChange={(e) => setNewBook({ ...newBook, totalCopies: e.target.value })}
              required
            />
          </div>
          <button type="submit">Add Book</button>
        </form>
      )}
<div>----</div>
      <button className='request-button' onClick={() => setRequestsVisible(!isRequestsVisible)}>
        {isRequestsVisible ? 'Hide Issue Requests' : 'Show Issue Requests'}
      </button>
      {isRequestsVisible && (
        <div>
          <h3>Issue Requests</h3>
          <ul className="request-list">
            {requests.map((request) => (
              <li key={request.req_id} className="request-item">
                <span>Request ID: {request.req_id}</span>
                <span>Book ID: {request.book_id}</span>
                <span>Reader ID: {request.reader_id}</span>
                <button className="accept-button" onClick={() => handleAcceptRequest(request.req_id)}>Accept</button>
                <button className="reject-button" onClick={() => handleRejectRequest(request.req_id)}>Reject</button>
              </li>
            ))}
          </ul>
        </div>
      )}
<div className='all-books'>
  <h3>All Books</h3>
  <ul className="book-list">
    {books.map((book) => (
      <li key={book.isbn} className="book-item">
        <span>{book.title}</span>
        <span>{book.authors}</span>
        <span>{book.publisher}</span>
        <button className='accept' onClick={() => handleEditClick(book)}>Update</button>
        <button className='delete' onClick={() => handleDeleteBook(book.isbn)}>Delete</button>
      </li>
    ))}
  </ul>
</div>

{isUpdateBookFormVisible && (
  <form onSubmit={handleUpdateBook}>
    <div className="form-group">
      <label htmlFor="isbn">ISBN:</label>
      <input
        type="text"
        id="isbn"
        value={currentBook.isbn}
        onChange={(e) => setCurrentBook({ ...currentBook, isbn: e.target.value })}
        disabled
      />
    </div>
    <div className="form-group">
      <label htmlFor="title">Title:</label>
      <input
        type="text"
        id="title"
        value={currentBook.title}
        onChange={(e) => setCurrentBook({ ...currentBook, title: e.target.value })}
        required
      />
    </div>
    <div className="form-group">
      <label htmlFor="authors">Authors:</label>
      <input
        type="text"
        id="authors"
        value={currentBook.authors}
        onChange={(e) => setCurrentBook({ ...currentBook, authors: e.target.value })}
        required
      />
    </div>
    <div className="form-group">
      <label htmlFor="publisher">Publisher:</label>
      <input
        type="text"
        id="publisher"
        value={currentBook.publisher}
        onChange={(e) => setCurrentBook({ ...currentBook, publisher: e.target.value })}
        required
      />
    </div>
    <div className="form-group">
      <label htmlFor="totalCopies">Total Copies:</label>
      <input
        type="number"
        id="totalCopies"
        value={currentBook.totalCopies}
        onChange={(e) => setCurrentBook({ ...currentBook, totalCopies: e.target.value })}
        required
      />
    </div>
    <button type="submit">Update Book</button>
  </form>
)}

    </div>
  );
};

export default AdminDashboard;
