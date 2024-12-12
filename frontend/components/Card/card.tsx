"use client";

import * as React from "react";
import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import { useRouter } from "next/navigation";
import { Button } from "../ui/button";

interface BasicCardProps {
  word: string;
  type: string;
  ticketCount: number;
  price: string;
  id: string;
}

export default function BasicCard({
  word,
  type,
  ticketCount,
  price,
  id,
}: BasicCardProps) {
  const router = useRouter();

  const handleDetails = () => {
    router.push(`/event/${id}`);
  };

  return (
    <Card variant="outlined" className="min-h-full">
      <CardContent>
        <Typography variant="h5" component="div">
          {word}
        </Typography>
        <Typography sx={{ mb: 1.5 }}>
          {new Intl.NumberFormat("id-ID", {
            style: "currency",
            currency: "IDR",
          }).format(Number(price))}
        </Typography>
        <Typography variant="body2">{type}</Typography>
        <Typography variant="body2">
          {`Jumlah Tiket Tersedia: ${ticketCount}`}
        </Typography>
      </CardContent>
      <CardActions>
        <Button
          onClick={handleDetails}
          variant="default"
          className="bg-blue-500 text-white hover:bg-blue-700"
        >
          Buy Ticket
        </Button>
      </CardActions>
    </Card>
  );
}
