"use client";

import React, { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { fetchUser, deleteUser, fetchUserEventData, deleteUserTicket } from "../../lib/api"; // Adjust the path to your api.ts file
import { UserTicket } from "@/lib/definition";
import { Button } from "@/components/ui/button";
import Link from "next/link";

const UserDashboard = () => {
  const router = useRouter();
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [tickets, setTickets] = useState<UserTicket[]>([]); // State to store fetched ticket data
  const [user, setUser] = useState<any | null>(null); // User state to store fetched user data
  const [error, setError] = useState<string | null>(null); // State for error handling
  const [loading, setLoading] = useState<boolean>(true); // State for loading indicator

  useEffect(() => {
    // Fetch user data when the component mounts
    const userData = localStorage.getItem("user"); // Assuming user data is stored in localStorage
    if (userData) {
      const token = JSON.parse(userData).token; // Extract token from user data
      fetchUser(token)
        .then((data) => {
          setUser(data.data); // Assuming user data is inside `data.data`
          setLoading(false); // Set loading to false after data is fetched
        })
        .catch((err) => {
          console.error(err);
          setError("Failed to fetch user data");
          setLoading(false); // Set loading to false even on error
        });

      fetchUserEventData(token).then((item) => {
        setTickets(item.data);
        setLoading(false);
      });
    } else {
      setError("No token found, please login");
      setLoading(false);
    }
  }, []);

  const handleReturn = () => {
    router.push("/");
  };

  const handleDelete = () => {
    setIsModalOpen(true);
  };

  const closeModal = () => {
    setIsModalOpen(false);
  };

  const confirmDelete = async () => {
    const userData = localStorage.getItem("user");
    if (userData) {
      const token = JSON.parse(userData).token;
      try {
        await deleteUser(token);
        console.log("User deleted successfully"); // Log success message
        localStorage.removeItem("user"); // Remove user data from localStorage
        router.push("/"); // Redirect to home page after deletion
      } catch (error) {
        console.error("Error in confirmDelete:", error); // Log error message
        setError("Failed to delete user");
      }
    }
    setIsModalOpen(false);
  };

  const deleteTicket = async (id: string) => {
    const userData = localStorage.getItem("user");
    if (userData) {
      const token = JSON.parse(userData).token;
      try {
        await deleteUserTicket(token, id);
        setTickets(tickets.filter((ticket) => ticket.id !== id));
      } catch (error) {
        console.error("Error deleting ticket:", error);
        setError("Failed to delete ticket");
      }
    }
  };

  // Display loading or error message if needed
  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error}</div>;
  }

  return (
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        height: "100vh",
      }}
    >
      <div
        style={{
          border: "1px solid black",
          padding: "40px",
          width: "500px",
          backgroundColor: "white",
          color: "black",
        }}
      >
        <h1 className="mb-6 text-center text-2xl font-bold">User Dashboard</h1>
        <div className="space-y-4">
          <div className="flex flex-col gap-2"></div>
          <p className="text-gray-800">
            <span className="font-semibold">Name:</span> {user?.name}
          </p>
          <p className="text-gray-800">
            <span className="font-semibold">Email:</span> {user?.email}
          </p>
        </div>

        <div className="mt-4 flex flex-row gap-4">
          <Button
            onClick={handleReturn}
            className="bg-blue-500 text-white transition-colors hover:bg-blue-600"
          >
            Back to Home
          </Button>
          <Link href="/tickets">
            <Button>Lihat daftar tiket</Button>
          </Link>
          <Button
            variant="destructive"
            className="bg-red-600 transition-colors hover:bg-red-700"
            onClick={handleDelete}
          >
            Delete this User
          </Button>
        </div>
        <br />
        <button onClick={handleReturn} style={{ color: "blue" }}>
          Back to Home
        </button>
        {/* <button onClick={handleDelete} style={{ color: "red" }}>
          Delete User
        </button> */}
      </div>
      {isModalOpen && (
        <div
          style={{
            position: "fixed",
            top: 0,
            left: 0,
            right: 0,
            bottom: 0,
            backgroundColor: "rgba(0,0,0,0.5)",
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
          }}
        >
          <div
            style={{
              backgroundColor: "white",
              padding: "20px",
              borderRadius: "5px",
              width: "80%",
              maxWidth: "500px",
            }}
          >
            <div style={{ padding: "16px", textAlign: "center" }}>
              <h1 style={{ color: "black" }}>Delete Account</h1>
              <p style={{ color: "black" }}>
                Are you sure you want to delete your account?
              </p>
              <br />
              <div style={{ display: "flex", justifyContent: "space-between" }}>
                <button
                  onClick={closeModal}
                  style={{
                    backgroundColor: "#ccc",
                    color: "black",
                    padding: "14px 20px",
                    border: "none",
                    cursor: "pointer",
                    width: "48%",
                  }}
                >
                  Cancel
                </button>
                <button
                  onClick={confirmDelete}
                  style={{
                    backgroundColor: "#f44336",
                    color: "white",
                    padding: "14px 20px",
                    border: "none",
                    cursor: "pointer",
                    width: "48%",
                  }}
                >
                  Delete
                </button>
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default UserDashboard;
