"use client";
import Box from "@mui/material/Box";
import Card from "@/components/Card/card";
import Typography from "@mui/material/Typography";
import { useEffect, useState } from "react";
import { fetchEventData } from "@/lib/api";

const Hero = () => {
  const [eventData, setEventData] = useState<any[]>([]);

  useEffect(() => {
    const userData = localStorage.getItem("user");
    if (userData) {
      const token = JSON.parse(userData).token;
      fetchEventData(token)
        .then((data) => {
          console.log(data.data);
          setEventData(data.data);
        })
        .catch((err) => {
          console.error(err);
        });
    }
  }, []);

  return (
    <>
      <Typography variant="h4" align="center" gutterBottom sx={{ mt: 15 }}>
        Welcome To Ticket List!
      </Typography>
      <Box
        sx={{
          m: 10,
          display: "grid",
          gridTemplateColumns: "repeat(3, 1fr)",
          gap: 2,
          justifyContent: "center",
          alignItems: "center",
        }}
      >
        {eventData.map((event: any) => (
          <Card
            key={event.id}
            word={event.name}
            type={event.type.name}
            ticketCount={event.quota}
            price={event.pricing}
            id={event.id}
          />
        ))}
      </Box>
    </>
  );
};

export default Hero;
