import React, { useState } from 'react';
import Login from './components/Login';
import FlightStatus from './components/FlightStatus';
import './App.css';

const App = () => {
  const [token, setToken] = useState(null);

  if (!token) {
    return <Login setToken={setToken} />;
  }

  return <FlightStatus token={token} />;
};

export default App;
