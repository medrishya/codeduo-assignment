// src/App.tsx
import React from "react";
import TaskHandler from "./components/TaskHandler";

const App: React.FC = () => {
  return (
    <div style={{ maxWidth: 600, margin: "0 auto" }}>
      <h1 style={{ marginLeft: "16px" }}>Task Manager</h1>
      <TaskHandler />
    </div>
  );
};

export default App;
