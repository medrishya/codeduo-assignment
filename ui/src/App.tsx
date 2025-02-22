// src/App.tsx
import React from "react";
import { Button } from "antd";

const App: React.FC = () => {
  return (
    <div style={{ padding: 20 }}>
      <h1>Hello, Ant Design!</h1>
      <Button type="primary">Click Me</Button>
    </div>
  );
};

export default App;
