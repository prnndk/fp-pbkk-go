"use client";

import React, { useEffect, useState } from "react";
import { useRouter, useParams } from "next/navigation";
import NotFound from "@/components/NotFound";

import { deleteUserTicket, fetchUserEventData } from "@/lib/api";
import { Button } from "@/components/ui/button";
import { Typography } from "@mui/material";
import Link from "next/link";

const Tickets = () => {
  const router = useRouter();

  const [data, setData] = useState<any | null>(null);
  const [token, setToken] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null); // Added state for error

  useEffect(() => {
    const userData = localStorage.getItem("user");
    if (userData) {
      setToken(JSON.parse(userData).token);
      const token = JSON.parse(userData).token;
      fetchUserEventData(token)
        .then((response) => {
          setData(response.data);
        })
        .catch((error) => {
          console.error("Failed to fetch event data:", error);
        });
    }
  }, []);

  if (!data) {
    return <NotFound />;
  }

  const deleteTicket = async (id: string) => {
    const userData = localStorage.getItem("user");
    if (userData) {
      const token = JSON.parse(userData).token;
      try {
        await deleteUserTicket(token, id);
        //reload windows using window
        router.push("/tickets");
      } catch (error) {
        console.error("Error deleting ticket:", error);
        setError("Failed to delete ticket");
      }
    }
  };
  return (
    <section className="flex h-screen items-center justify-center p-8">
      <div className="container mx-auto max-w-7xl">
        <h4 className="mb-6 text-2xl font-bold text-white">
          List User Tickets
        </h4>
        <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
          {data.map((item, index) => (
            <Link href={`/detail-ticket/${item.id}`} key={index}>
              <div
                key={index}
                className="flex flex-col gap-4 rounded-lg border border-gray-200 bg-white p-6 shadow-sm transition-all hover:shadow-md"
              >
                <h5 className="text-xl font-bold text-gray-900">
                  {item.event.name}
                </h5>
                <div className="space-y-2">
                  <Typography variant="body1" className="text-gray-600">
                    Jumlah Tiket: {item.quantity}
                  </Typography>
                  <Typography variant="body1" className="text-gray-600">
                    Total Harga: Rp {item.total_price.toLocaleString("id-ID")}
                  </Typography>
                </div>
                <Button
                  variant="destructive" // Changed from "contained" to "destructive"
                  className="bg-red-600 text-white hover:bg-red-800"
                  onClick={() => deleteTicket(item.id)}
                >
                  Delete Ticket
                </Button>
                {error && <Typography color="error">{error}</Typography>}
              </div>
            </Link>
          ))}
        </div>
      </div>
    </section>
  );
};

export default Tickets;
