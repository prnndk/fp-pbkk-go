"use client";

import * as React from 'react';
import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import { useRouter } from 'next/navigation';

interface BasicCardProps {
  word: string;
  description: string;
  ticketCount: number;
  price: string;
}

export default function BasicCard({ word, description, ticketCount, price }: BasicCardProps) {
  const router = useRouter();

  const handleDetails = () => {
    router.push('/carddetails');
  };

  return (
      <Card sx={{ backgroundColor: 'rgb(29 36 48 / var(--tw-bg-opacity))' }} className='text-dark dark:text-white'>
        <CardContent>
          <Typography variant="h5" component="div">
            {word}
          </Typography>
          <Typography sx={{ mb: 1.5 }}>{price}</Typography>
          <Typography variant="body2">
            {description}
            <br />
            <br />
            {`Jumlah Tiket = ${ticketCount}`}
          </Typography>
        </CardContent>
        <CardActions>
          <Button onClick={handleDetails} size="small">Buy Ticket</Button>
        </CardActions>
      </Card>
  );
}
