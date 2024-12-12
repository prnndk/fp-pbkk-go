"use client";

import React, { useEffect, useState } from "react";
import { useRouter, useParams } from "next/navigation";
import NotFound from "@/components/NotFound";

import { fetchSingleTicket, deleteUserTicket } from "@/lib/api"; // Updated import
import { Typography } from "@mui/material";

const Tickets = () => {
  const router = useRouter();
  const { id } = useParams() as { id: string };

  const [data, setData] = useState<any | null>(null);
  const [token, setToken] = useState<string | null>(null);

  useEffect(() => {
    const userData = localStorage.getItem("user");
    if (userData) {
      setToken(JSON.parse(userData).token);
      const token = JSON.parse(userData).token;
      fetchSingleTicket(token, id)
        .then((response) => {
          setData(response.data);
        })
        .catch((error) => {
          console.error("Failed to fetch event data:", error);
        });
    }
  }, [id]);

  if (!data) {
    return <NotFound />;
  }

  return (
    <section className="flex h-screen items-center justify-center p-8">
      <div className="container mx-auto max-w-7xl">
        <h4 className="mb-6 text-2xl font-bold text-white">
          Data Pembayaran Ticket
        </h4>
        <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
          <div
            key={data.id}
            className="flex flex-col gap-4 rounded-lg border border-gray-200 bg-white p-6 shadow-sm transition-all hover:shadow-md"
          >
            <h5 className="text-xl font-bold text-gray-900">
              {data.event.name}
            </h5>
            <div className="space-y-2">
              <Typography
                variant="body1"
                className={`${
                  data.is_paid ? "text-green-600" : "text-gray-600"
                }`}
              >
                Status Pembayaran: {data.is_paid ? "Lunas" : "Belum Lunas"}
              </Typography>
              <Typography variant="body1" className="text-gray-600">
                Jumlah Tiket: {data.quantity}
              </Typography>
              <Typography variant="body1" className="text-gray-600">
                Total Harga: Rp {data.total_price.toLocaleString("id-ID")}
              </Typography>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default Tickets;
