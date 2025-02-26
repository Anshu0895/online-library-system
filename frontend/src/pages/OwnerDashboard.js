
import React, { useState, useEffect } from "react";
import api from "../utils/api";
import "../Css/OwnerDashboard.css";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";


const OwnerDashboard = ({ token }) => {
  const [libraries, setLibraries] = useState([]);
  const [admins, setAdmins] = useState([]);

  const [showCreateModal, setShowCreateModal] = useState(false);
  const [showUpdateModal, setShowUpdateModal] = useState(false);
  const [showLibraryModal, setShowLibraryModal] = useState(false);

  const [selectedAdmin, setSelectedAdmin] = useState(null);
  const [adminData, setAdminData] = useState({ name: "", email: "", password: "", contact_number: "" });
  const [newLibrary, setNewLibrary] = useState({ name: "" });

  useEffect(() => {
    if (token) {
      fetchLibraries();
      fetchAdmins();
    } }, [token]);

  const fetchLibraries = async () => {
    try {
      const res = await api.get("/libraries", { headers: { Authorization: `${token}` } });
      setLibraries(res.data);
    } catch (err) {
      toast.error("Failed to fetch libraries");
    }
  };

  const fetchAdmins = async () => {
    try {
      const res = await api.get("/users", { headers: { Authorization: `${token}` } });
      setAdmins(res.data.filter(user => user.role === "LibraryAdmin"));
    } catch (err) {
      toast.error("Failed to fetch admins");
    }
  };

  const handleCreateAdmin = async (e) => {
    e.preventDefault();
    try {
      const res = await api.post("/users", adminData, { headers: { Authorization: `${token}` } });
      setAdmins([...admins, res.data]);
      toast.success("Admin created successfully");
      setShowCreateModal(false);
    } catch (err) {
      toast.error("Error creating admin");
    }
  };

  const handleUpdateAdmin = async (e) => {
    e.preventDefault();
    try {
      const res = await api.put(`/users/${selectedAdmin.id}`, selectedAdmin, { headers: { Authorization: `${token}` } });
      setAdmins(admins.map(admin => (admin.id === selectedAdmin.id ? res.data : admin)));
      toast.success("Admin updated successfully");
      setShowUpdateModal(false);
    } catch (err) {
      toast.error("Error updating admin");
    }
  };

  const handleDeleteAdmin = async (id) => {
    try {
      await api.delete(`/users/${id}`, { headers: { Authorization: `${token}` } });
      setAdmins(admins.filter(admin => admin.id !== id));
      toast.success("Admin deleted successfully");
    } catch (err) {
      toast.error("Error deleting admin");
    }
  };

  const handleCreateLibrary = async (e) => {
    e.preventDefault();
    try {
      await api.post("/libraries", newLibrary, { headers: { Authorization: `${token}` } });
      fetchLibraries();
      toast.success("Library created successfully");
      setShowLibraryModal(false);
    } catch (err) {
      toast.error("Error creating library");
    }
  };

  return (
   
     
      <div className="owner-dashboard">

        {/* Library List */}
        <div className="card">
          <h2>Owner Dashboard</h2>
          <h3>Libraries</h3>
          <button className="add-btn" onClick={() => setShowLibraryModal(true)}>+ Create Library</button>
          <ul>
            {libraries.map(lib => (
              <li key={lib.id}>{lib.name} (ID: {lib.id})</li>
            ))}
          </ul>
        </div>

        {/* Admin List */}
        <div className="card">
          <h3>Library Admins</h3>
          <button className="add-btn" onClick={() => setShowCreateModal(true)}>+ Add Admin</button>
          <ul>
            {admins.map(admin => (
              <li key={admin.id}>
                {admin.name} ({admin.email})
                <div className="action-buttons">
                  <button className="edit-btn" onClick={() => { setSelectedAdmin(admin); setShowUpdateModal(true); }}>Edit</button>
                  <button className="delete-btn" onClick={() => handleDeleteAdmin(admin.id)}>Delete</button>
                </div>
              </li>
            ))}
          </ul>
        </div>

        {/* Create Admin Modal */}
        {showCreateModal && (
          <div className="modal">
            <div className="modal-content">
              <h3>Create Admin</h3>
              <form onSubmit={handleCreateAdmin}>
                <input type="text" placeholder="Name" value={adminData.name} onChange={(e) => setAdminData({ ...adminData, name: e.target.value })} required />
                <input type="email" placeholder="Email" value={adminData.email} onChange={(e) => setAdminData({ ...adminData, email: e.target.value })} required />
                <input type="password" placeholder="Password" value={adminData.password} onChange={(e) => setAdminData({ ...adminData, password: e.target.value })} required />
                <input type="text" placeholder="Contact Number" value={adminData.contact_number} onChange={(e) => setAdminData({ ...adminData, contact_number: e.target.value })} required />
                <button type="submit">Create</button>
                <button className="cancel-btn" onClick={() => setShowCreateModal(false)}>Cancel</button>
              </form>
            </div>
          </div>
        )}

        {/* Update Admin Modal */}
        {showUpdateModal && (
          <div className="modal">
            <div className="modal-content">
              <h3>Update Admin</h3>
              <form onSubmit={handleUpdateAdmin}>
                <input type="text" placeholder="Name" value={selectedAdmin.name} onChange={(e) => setSelectedAdmin({ ...selectedAdmin, name: e.target.value })} required />
                <input type="email" placeholder="Email" value={selectedAdmin.email} onChange={(e) => setSelectedAdmin({ ...selectedAdmin, email: e.target.value })} required />
                <input type="text" placeholder="Contact Number" value={selectedAdmin.contact_number} onChange={(e) => setSelectedAdmin({ ...selectedAdmin, contact_number: e.target.value })} required />
                <button type="submit">Update</button>
                <button className="cancel-btn" onClick={() => setShowUpdateModal(false)}>Cancel</button>
              </form>
            </div>
          </div>
        )}

        {/* Create Library Modal */}
        {showLibraryModal && (
          <div className="modal">
            <div className="modal-content">
              <h3>Create Library</h3>
              <form onSubmit={handleCreateLibrary}>
                <input type="text" placeholder="Library Name" value={newLibrary.name} onChange={(e) => setNewLibrary({ name: e.target.value })} required />
                <button type="submit">Create</button>
                <button className="cancel-btn" onClick={() => setShowLibraryModal(false)}>Cancel</button>
              </form>
            </div>
          </div>
        )}
      </div>
  
  );
};

export default OwnerDashboard;