import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { getLinks, createLink, deleteLink } from "../api/api";

function Dashboard() {
  const [links, setLinks] = useState([]);
  const [title, setTitle] = useState("");
  const [url, setUrl] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    loadLinks();
  }, []);

  const loadLinks = async () => {
    const data = await getLinks();
    setLinks(data);
  };

  const handleCreate = async () => {
    await createLink(title, url);
    setTitle("");
    setUrl("");
    loadLinks();
  };

  const handleDelete = async (id) => {
    await deleteLink(id);
    loadLinks();
  };

  const logout = () => {
    localStorage.removeItem("token");
    navigate("/login");
  };

  return (
    <div style={{ padding: "40px" }}>
      <h2>Dashboard</h2>
      <button onClick={logout}>Logout</button>

      <h3>Add Link</h3>
      <input placeholder="Title" value={title} onChange={e => setTitle(e.target.value)} />
      <input placeholder="URL" value={url} onChange={e => setUrl(e.target.value)} />
      <button onClick={handleCreate}>Add</button>

      <h3>Your Links</h3>
      {links.map(link => (
        <div key={link.ID}>
          <strong>{link.Title}</strong> — {link.URL}
          <button onClick={() => handleDelete(link.ID)}>Delete</button>
        </div>
      ))}
    </div>
  );
}

export default Dashboard;
