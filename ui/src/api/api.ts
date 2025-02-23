const API_URL = "http://localhost:8080/tasks";

export const fetchTasks = async (page: number, limit: number) => {
  const response = await fetch(`${API_URL}?page=${page}&limit=${limit}`);
  return response.json();
};

export const createTask = async (task: { name: string; status: string }) => {
  const response = await fetch(API_URL, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(task),
  });
  return response.json();
};

export const updateTask = async (
  id: number,
  updates: Partial<{ name: string; status: string }>
) => {
  const response = await fetch(`${API_URL}/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(updates),
  });
  return response.json();
};

export const deleteTask = async (id: number) => {
  return fetch(`${API_URL}/${id}`, { method: "DELETE" });
};
