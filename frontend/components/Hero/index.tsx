import Link from "next/link";
import Box from '@mui/material/Box';
import Card from "@/components/Card/card";
import Typography from '@mui/material/Typography';

const Hero = () => {
  return (
    <>
    <Typography variant="h4" align="center" gutterBottom sx={{ mt: 15 }}>
      Welcome To Ticket List!
    </Typography>
    <Box sx={{ m: 10, display: 'grid', gridTemplateColumns: 'repeat(3, 1fr)', gap: 2, justifyContent: 'center', alignItems: 'center' }}>
     <Card 
       word="Event 1" 
       description="Lorem ipsum dolor sit amet, consectetur adipisicing elit. Mollitia enim veritatis est quas maiores? Sed, placeat quae eius fuga fugit animi dicta nostrum ratione cumque, consectetur beatae est, modi sapiente." 
       ticketCount={10} 
       price="$10"
     />
     <Card 
       word="Event 2" 
       description="Lorem ipsum dolor sit amet, consectetur adipisicing elit. Mollitia enim veritatis est quas maiores? Sed, placeat quae eius fuga fugit animi dicta nostrum ratione cumque, consectetur beatae est, modi sapiente." 
       ticketCount={20} 
       price="$20"
     />
     <Card 
       word="Event 3" 
       description="Lorem ipsum dolor sit amet, consectetur adipisicing elit. Mollitia enim veritatis est quas maiores? Sed, placeat quae eius fuga fugit animi dicta nostrum ratione cumque, consectetur beatae est, modi sapiente." 
       ticketCount={30} 
       price="$30"
     />
     <Card 
       word="Event 4" 
       description="Lorem ipsum dolor sit amet, consectetur adipisicing elit. Mollitia enim veritatis est quas maiores? Sed, placeat quae eius fuga fugit animi dicta nostrum ratione cumque, consectetur beatae est, modi sapiente." 
       ticketCount={40} 
       price="$40"
     />
    </Box>
    </>
  );
};

export default Hero;
