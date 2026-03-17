import { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import { registerUser } from "../api/api";

function Register() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");
  const navigate = useNavigate();

  const handleRegister = async () => {
    const data = await registerUser(email, password);

    if (!data.error) {
      setMessage("Registration successful!");
      setTimeout(() => navigate("/login"), 1000);
    } else {
      setMessage(data.error);
    }
  };

  return (
    <div style={{ padding: "40px" }}>
      <h2>Register</h2>

      <input placeholder="Email" value={email} onChange={e => setEmail(e.target.value)} />
      <br /><br />

      <input type="password" placeholder="Password" value={password} onChange={e => setPassword(e.target.value)} />
      <br /><br />

      <button onClick={handleRegister}>Register</button>

      <p>{message}</p>

      <Link to="/login">Back to Login</Link>
    </div>
  );
}

export default Register;
