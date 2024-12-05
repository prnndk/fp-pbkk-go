'use client';

import React, { useState } from 'react';
import { useRouter } from 'next/navigation';

const UserDashboard = () => {
    const router = useRouter();
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [tickets, setTickets] = useState([
        { id: 1, name: 'Concert A' },
        { id: 2, name: 'Concert B' },
        { id: 3, name: 'Concert C' },
    ]);

    const handleReturn = () => {
        router.push('/');
    };

    const handleDelete = () => {
        setIsModalOpen(true);
    };

    const closeModal = () => {
        setIsModalOpen(false);
    };

    const confirmDelete = () => {
        // Add delete logic here
        setIsModalOpen(false);
    };

    const deleteTicket = (id) => {
        setTickets(tickets.filter(ticket => ticket.id !== id));
    };

    const user = {
        name: 'John Doe',
        email: 'john.doe@example.com',
    };

    return (
        <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
            <div style={{ border: '1px solid black', padding: '40px', width: '500px', backgroundColor: 'white', color: 'black' }}>
                <h1 style={{ textAlign: 'center' }}>User Dashboard</h1>
                <br />
                <br />
                <div>
                    <p><strong>Name:</strong> {user.name}</p>
                    <p><strong>Email:</strong> {user.email}</p>
                </div>
                <br />
                <div>
                    <h2>Event Tickets</h2>
                    <ul>
                        {tickets.map(ticket => (
                            <li key={ticket.id} style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                                {ticket.name}
                                <button onClick={() => deleteTicket(ticket.id)} style={{ color: 'red' }}>Delete</button>
                            </li>
                        ))}
                    </ul>
                </div>
                <br />
                <button onClick={handleReturn} style={{ color: 'blue' }}>Back to Home</button>
                <br />
                <button onClick={handleDelete} style={{ color: 'red' }}>Delete User</button>
            </div>
            {isModalOpen && (
                <div style={{
                    position: 'fixed', top: 0, left: 0, right: 0, bottom: 0, backgroundColor: 'rgba(0,0,0,0.5)',
                    display: 'flex', justifyContent: 'center', alignItems: 'center'
                }}>
                    <div style={{ backgroundColor: 'white', padding: '20px', borderRadius: '5px', width: '80%', maxWidth: '500px' }}>
                        <div style={{ padding: '16px', textAlign: 'center' }}>
                            <h1 style={{ color: 'black' }}>Delete Account</h1>
                            <p style={{ color: 'black' }}>Are you sure you want to delete your account?</p>
                            <br />
                            <div style={{ display: 'flex', justifyContent: 'space-between' }}>
                                <button onClick={closeModal} style={{ backgroundColor: '#ccc', color: 'black', padding: '14px 20px', border: 'none', cursor: 'pointer', width: '48%' }}>Cancel</button>
                                <button onClick={confirmDelete} style={{ backgroundColor: '#f44336', color: 'white', padding: '14px 20px', border: 'none', cursor: 'pointer', width: '48%' }}>Delete</button>
                            </div>
                        </div>
                    </div>
                </div>
            )}
        </div>
    );
};

export default UserDashboard;
