'use client';

import React from 'react';
import { useRouter } from 'next/navigation';

const EventSquare = () => {
  const router = useRouter();

  const handleReturn = () => {
    router.push('/');
  };

  return (
    <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
      <div style={{ border: '1px solid black', padding: '40px', width: '500px', backgroundColor: 'white', color: 'black' }}>
        <h2>Event 4</h2>
        <br />
        <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Mollitia enim veritatis est quas maiores? Sed, placeat quae eius fuga fugit animi dicta nostrum ratione cumque, consectetur beatae est, modi sapiente.</p>
        <br />
        <p>Tickets Available: 40</p>
        <p>Price: $40</p>
        <button style={{ color: 'blue' }}>Beli Tiket</button>
        <br />
        <br />
        <button onClick={handleReturn} style={{ color: 'blue' }}>Back to Home</button>
      </div>
    </div>
  );
};

export default EventSquare;
