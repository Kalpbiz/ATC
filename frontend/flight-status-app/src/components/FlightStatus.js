import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './FlightStatus.css';

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
    <div className="flight-status-container">
      <h1>Flight Status</h1>
      <ul className="flight-list">
        {flights.map(flight => (
          <li key={flight.id} className="flight-item">
            <div className="flight-details">
              <span>Flight Number: {flight.flightNumber}</span>
              <span>Status: {flight.status}</span>
              <span>Gate: {flight.gate}</span>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default FlightStatus;
