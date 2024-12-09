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
        <form>
          <h3>Payment Details</h3>
          <p>Ticket ID: <strong>5661fd3c-208a-48ab-8867-71166f51c0da</strong></p>
          <p>Metode Pembayaran: <strong>transfer</strong></p>
          <p>Bukti Bayar: https://storage-server.aryagading.com/893948930-fffgpgpeer3-33-2</p>
          <br />
          <button style={{ color: 'blue' }}>Bayar Tiket</button>
        </form>
        <br />
        <button onClick={handleReturn} style={{ color: 'blue' }}>Back to Home</button>
      </div>
    </div>
  );
};

export default EventSquare;
