
import React, { useState, useEffect } from 'react';
import api from '../utils/api';
import '../Css/AdminDashboard.css';
import { toast } from 'react-toastify';
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
      toast.success('Book added successflly', {
              position: "top-center",
              style: { backgroundColor: 'black', color: 'white' },
            });
      fetchBooks(); 
    } catch (err) {
      setError('Failed to add book');
      toast.error('Failed to add book', {
        position: "top-center",
        style: { backgroundColor: 'black', color: 'white' },
      });
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
      toast.success('Book deleted successflly', {
        position: "top-center",
        style: { backgroundColor: 'black', color: 'white' },
      });
      fetchBooks(); // Refresh books list
    } catch (err) {
      setError('Failed to delete book');
      toast.error('Failed to delete book', {
        position: "top-center",
        style: { backgroundColor: 'black', color: 'white' },
      });
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
      toast.success('Book updated successflly', {
        position: "top-center",
        style: { backgroundColor: 'black', color: 'white' },
      });
      fetchBooks(); // Refresh books list
      setUpdateBookFormVisible(false);
    } catch (err) {
      setError('Failed to update book');
      toast.error('Failed to update book', {
        position: "top-center",
        style: { backgroundColor: 'black', color: 'white' },
      });
    }
  };

  const handleEditClick = (book) => {
    setCurrentBook(book);
    setUpdateBookFormVisible(true);
  };

  const handleAcceptRequest = async (requestId) => {
    try {
      const adminId = localStorage.getItem("id"); // Fetch admin ID from local storage

      if (!adminId) {
        setError("Admin ID not found in local storage");
        return;
      }

      console.log("Sending ApproverId:", adminId);

      const response = await api.put(
        `/requests/${requestId}/approve`,
        { approver_id: Number(adminId) },
        {
          headers: {
            Authorization: `${token}`,
            "Content-Type": "application/json",
          },
        }
      );

      if (response.status === 200) {
        setRequests((prevRequests) =>
          prevRequests.filter((request) => request.req_id !== requestId)
        );
        fetchBooks();
        setSuccess("Request approved successfully");
        toast.success('Request approved successfully', {
          position: "top-center",
          style: { backgroundColor: 'black', color: 'white' },
        });
        
      } else {
        setError("Failed to approve request");
      }
    } catch (err) {
      setError("Failed to approve request");
      toast.error('Failed to approve request', {
        position: "top-center",
        style: { backgroundColor: 'black', color: 'white' },
      });
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
        toast.success('Request rejected successfully', {
          position: "top-center",
          style: { backgroundColor: 'black', color: 'white' },
        });
      } else {
        setError('Failed to reject request');
      }
    } catch (err) {
      setError('Failed to reject request');
      toast.error('Failed to reject request', {
        position: "top-center",
        style: { backgroundColor: 'black', color: 'white' },
      });
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
        <button className='add-book' onClick={() => setAddBookFormVisible(!isAddBookFormVisible)}>
          {isAddBookFormVisible ? 'Hide Add Book Form' : 'Show Add Book Form'}
        </button>
        {isAddBookFormVisible && (
          <div className='admin1-form'>
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
                onChange={(e) => setNewBook({ ...newBook, lib_id: parseInt(e.target.value, 10) })}
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
              <label htmlFor="total_copies">Total Copies:</label>
              <input
                type="number"
                id="total_copies"
                value={newBook.total_copies}
                onChange={(e) => setNewBook({ ...newBook, total_copies: parseInt(e.target.value, 10) })}
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
          </div>
        )}
        <button className='request-button' onClick={() => setRequestsVisible(!isRequestsVisible)}>
          {isRequestsVisible ? 'Hide Issue Requests' : 'Show Issue Requests'}
        </button>
        {isRequestsVisible && (
          <div className='requests-table'>
            <h3>Issue Requests</h3>
            <table>
              <thead>
                <tr>
                  <th>Request ID</th>
                  <th>Book ID</th>
                  <th>Reader ID</th>
                  <th>Status</th>
                  <th>Actions</th>
                  <th>Actions</th>
                </tr>
              </thead>
              <tbody>
                {requests.map((request) => (
                  <tr key={request.req_id}>
                    <td>{request.req_id}</td>
                    <td>{request.book_id}</td>
                    <td>{request.reader_id}</td>
                    <td>{request.request_type}</td>
                    <td><button className="accept-button" onClick={() => handleAcceptRequest(request.req_id)}>Accept</button></td>
                    <td><button className="reject-button" onClick={() => handleRejectRequest(request.req_id)}>Reject</button></td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
        <div className='all-books'>
          <h3>All Books</h3>
          <table>
            <thead>
              <tr>
                <th>ISBN</th>
                <th>Title</th>
                <th>Authors</th>
                <th>Publisher</th>
                <th>Total Copies</th>
                <th>Available Copies</th>
                <th>Actions</th>
                <th>Actions</th>
              </tr>
            </thead>
            <tbody>
              {books.map((book) => (
                <tr key={book.isbn}>
                  <td>{book.isbn}</td>
                  <td>{book.title}</td>
                  <td>{book.authors}</td>
                  <td>{book.publisher}</td>
                  <td>{book.total_copies}</td>
                  <td>{book.available_copies}</td>
                  <td><button className='accept' onClick={() => handleEditClick(book)}>Update</button></td>
                  <td><button className='delete' onClick={() => handleDeleteBook(book.isbn)}>Delete</button></td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>

        {isUpdateBookFormVisible && (

          <div className='update-form'>
          <form onSubmit={handleUpdateBook}>
            <div className='book-form'>
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
                <label htmlFor="available_copies">Available Copies:</label>
                <input
                  type="number"
                  id="available_copies"
                  value={currentBook.available_copies}
                  onChange={(e) => setCurrentBook({ ...currentBook, available_copies: parseInt(e.target.value, 10) })}
                  required
                />
              </div>
              <button type="submit">Update Book</button>
            </div>
          </form>
          </div>
        )}
      </div>
    </div>
  );
};

export default AdminDashboard;