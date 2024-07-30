import React, { useState, useEffect } from 'react';
import axios from 'axios';

const FlightStatus = ({ token }) => {
  const [flights, setFlights] = useState([]);

  useEffect(() => {
    const fetchFlights = async () => {
      const response = await axios.get('/api/flights', {
        headers: { Authorization: `Bearer ${token}` },
      });
      setFlights(response.data);
    };

    fetchFlights();
  }, [token]);

  return (
    <div>
      <h1>Flight Status</h1>
      <ul>
        {flights.map(flight => (
          <li key={flight.id}>{flight.status}</li>
        ))}
      </ul>
    </div>
  );
};

export default FlightStatus;
