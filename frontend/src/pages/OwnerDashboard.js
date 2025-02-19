import React, { useState ,useEffect} from 'react';
import api from '../utils/api';
import '../Css/OwnerDashboard.css';

const OwnerDashboard = ({ token }) => {
  const [libraries, setLibraries] = useState([]);
  const [name, setName] = useState('');
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  useEffect(() => {
    const fetchLibraries = async () => {
      try {
        const response = await api.get('/libraries', {
          headers: {
            Authorization: `${token}`, 
          },
        });
        setLibraries(response.data);
      } catch (err) {
        console.error(err);
      }
    };
    fetchLibraries();
  }, [token]);

  const handleCreateLibrary = async (e) => {
    e.preventDefault();
    setError('');
    setSuccess('');

    try {
      const response = await api.post('/libraries', { name }, {
        headers: {
          Authorization: `${token}`,
        },
      });
      setLibraries([...libraries, response.data]);
      setSuccess('Library created successfully');
    } catch (err) {
      if (err.response && err.response.status === 409) {
        setError(err.response.data.error);
      } else {
        setError(err.response.data.error);
      }
    }
  };

  return (
    <div className='owner-box'>
      <div className="dashboard-container">
        <h2>Owner Dashboard</h2>
        {error && <p className="error-message">{error}</p>}
        {success && <p className="success-message">{success}</p>}
        <form onSubmit={handleCreateLibrary}>
          <div className="form-group">
            <label htmlFor="name">Library Name:</label>
            <input
              type="text"
              id="name"
              value={name}
              onChange={(e) => setName(e.target.value)}
              required
            />
          </div>
          <button type="submit">Create Library</button>
        </form>
        <h3>Existing Libraries</h3>
        <ul>
          {libraries.map((lib) => (
            <li key={lib.id}>{lib.name}</li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default OwnerDashboard;