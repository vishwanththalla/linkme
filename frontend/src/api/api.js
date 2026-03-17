const API_URL = "http://localhost:8080";

export async function loginUser(email, password) {
  try {
    const response = await fetch(`${API_URL}/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, password }),
    });

    return await response.json();
  } catch (error) {
    return { error: "Server not reachable" };
  }
}

export async function registerUser(email, password) {
  try {
    const response = await fetch(`${API_URL}/register`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, password }),
    });

    return await response.json();
  } catch (error) {
    return { error: "Server not reachable" };
  }
}

/* -------- LINK APIs -------- */

export async function getLinks() {
  const response = await fetch(`${API_URL}/links`);
  return await response.json();
}

export async function createLink(url) {
  const response = await fetch(`${API_URL}/links`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ url }),
  });

  return await response.json();
}

export async function deleteLink(id) {
  const response = await fetch(`${API_URL}/links/${id}`, {
    method: "DELETE",
  });

  return await response.json();
}