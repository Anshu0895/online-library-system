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
    lib_id: '',
    title: '',
    authors: '',
    publisher: '',
    version: '',
    total_copies: '',
    available_copies: '',
  });
  const [reload, setReload] = useState(false);

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
      const response = await api.get('/pending-requests', {
        headers: {
          Authorization: `${token}`,
        },
      });
      setRequests(Array.isArray(response.data.requests) ? response.data.requests : []);
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
      if (response.status === 200) {
        setRequests((prevRequests) => prevRequests.filter((request) => request.req_id !== requestId));
        setSuccess('Request approved successfully');
      } else {
        setError('Failed to approve request');
      }
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
      if (response.status === 200) {
        setRequests((prevRequests) => prevRequests.filter((request) => request.req_id !== requestId));
        setSuccess('Request rejected successfully');
      } else {
        setError('Failed to reject request');
      }
    } catch (err) {
      setError('Failed to reject request');
    }
  };

  useEffect(() => {
    fetchBooks();
    fetchRequests();
  }, [reload]);
  return (
    <div className='bap'>
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
            <label htmlFor="lib_id">LIB_ID:</label>
            <input
              type="text"
              id="lib_id"
              value={newBook.lib_id}
              onChange={(e) => setNewBook({ ...newBook, lib_id: parseInt(e.target.value, 10)  })}
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
            <label htmlFor="version">Version:</label>
            <input
              type="text"
              id="version"
              value={newBook.version}
              onChange={(e) => setNewBook({ ...newBook, version: e.target.value })}
              required
            />
          </div>
          <div className="form-group">
            <label htmlFor=" total_copies">Total Copies:</label>
            <input
              type="number"
              id="total_copies"
              value={newBook.total_copies}
              onChange={(e) => setNewBook({ ...newBook,  total_copies: parseInt(e.target.value, 10) })}
              required
            />
          </div>
          <div className="form-group">
            <label htmlFor="available_copies">Available Copies:</label>
            <input
              type="number"
              id="available_copies"
              value={newBook.available_copies}
              onChange={(e) => setNewBook({ ...newBook, available_copies: parseInt(e.target.value, 10) })}
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
  <div className="book-list-header">
    <span>ISBN</span>
    <span>Title</span>
    <span>Authors</span>
    <span>Publisher</span>
    <span>Total Copies</span>
    <span>Available Copies</span>
    <span>Actions</span>
    <span>Actions</span>
  </div>
  <ul className="book-list">
    {books.map((book) => (
      <li key={book.isbn} className="book-item">
         <span>{book.isbn}</span>
        <span>{book.title}</span>
        <span>{book.authors}</span>
        <span>{book.publisher}</span>
        <span>{book.total_copies}</span>
        <span>{book.available_copies}</span>
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
      <label htmlFor="total_copies">Total Copies:</label>
      <input
        type="number"
        id="total_copies"
        value={currentBook.total_copies}
        onChange={(e) => setCurrentBook({ ...currentBook, total_copies: parseInt(e.target.value, 10) })}
        required
      />
      </div>
      <div className="form-group">
      <label htmlFor="available_copies:">Available Copies:</label>
      <input
        type="number"
        id="available_copies:"
        value={currentBook.available_copies}
        onChange={(e) => setCurrentBook({ ...currentBook, available_copies: parseInt(e.target.value, 10) })}
        required
      />
    </div>
    <button type="submit">Update Book</button>
  </form>
)}

    </div>
    </div>
  );
};

export default AdminDashboard;
