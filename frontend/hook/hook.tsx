import { useEffect, useState } from "react";
import { fetchEvents } from "../lib/api"; // Adjust the path to your service file
import { ApiResponse, Event } from "../lib/definition";

export const useEvents = () => {
  const [events, setEvents] = useState<Event[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const loadEvents = async () => {
      try {
        setLoading(true);
        const response = await fetchEvents();
        if (response.status) {
          setEvents(response.data);
        } else {
          setError(response.message);
        }
      } catch (err) {
        setError("Failed to fetch events");
      } finally {
        setLoading(false);
      }
    };

    loadEvents();
  }, []);

  return { events, loading, error };
};
