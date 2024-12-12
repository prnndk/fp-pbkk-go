import axios from "axios";
import { ApiResponse, Event, User, UserTicket } from "../lib/definition"; // Adjust the path to your interface file

const API_BASE_URL = "http://127.0.0.1:8888/api"; // Replace with your backend URL

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

// Fetch all events
export const fetchEvents = async (): Promise<ApiResponse<Event[]>> => {
  const response = await api.get<ApiResponse<Event[]>>("/events");
  return response.data;
};

// User login
export const loginUser = async (email, password) => {
  const formData = new FormData();
  formData.append("email", email);
  formData.append("password", password);

  const res = await fetch(`${API_BASE_URL}/user/login`, {
    method: "POST",
    body: formData,
  });

  if (!res.ok) {
    throw new Error("Failed to login");
  }

  const data = await res.json();
  return data; // This returns the response including the token and role
};

// User registration
export const registerUser = async (
  name: string,
  phone: string,
  email: string,
  password: string,
): Promise<ApiResponse<User>> => {
  const formData = new FormData();
  formData.append("email", email);
  formData.append("name", name);
  formData.append("phone_number", phone);
  formData.append("password", password);

  const response = await api.post<ApiResponse<User>>("/user", formData);
  return response.data;
};

// Fetch user data (Get Me)
export const fetchUser = async (token: string): Promise<ApiResponse<User>> => {
  try {
    const response = await api.get<ApiResponse<User>>("/user/me", {
      headers: {
        Authorization: `Bearer ${token}`, // Adding the Bearer token for authorization
      },
    });
    if (response.status !== 200) {
      throw new Error("Failed to fetch user data");
    }
    return response.data;
  } catch (error) {
    console.error("Error fetching user data:", error);
    throw new Error("Failed to fetch user data");
  }
};

// Delete user
export const deleteUser = async (token: string): Promise<void> => {
  try {
    const response = await api.delete("/user", {
      // Adjust the endpoint URL as needed
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    console.log("Delete user response status:", response.status); // Log response status
    if (response.status !== 200) {
      throw new Error("Failed to delete user");
    }
  } catch (error) {
    console.error("Error deleting user:", error);
    throw new Error("Failed to delete user");
  }
};

export const fetchEventData = async (
  token: string,
): Promise<ApiResponse<Event[]>> => {
  try {
    const response = await api.get<ApiResponse<Event[]>>("/event", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    if (response.status !== 200) {
      throw new Error("Failed to fetch event data");
    }
    return response.data;
  } catch (error) {
    console.error("Error fetching event data:", error);
    throw new Error("Failed to fetch event data");
  }
};

export const fetchUserEventData = async (
  token: string,
): Promise<ApiResponse<UserTicket[]>> => {
  try {
    const response = await api.get<ApiResponse<UserTicket[]>>("/ticket/me", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    if (response.status !== 200) {
      throw new Error("Failed to fetch event data");
    }
    return response.data;
  } catch (error) {
    console.error("Error fetching event data:", error);
    throw new Error("Failed to fetch event data");
  }
};

export const fetchEventById = async (token: string, id: string) => {
  try {
    const response = await api.get(`/event/${id}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    if (response.status !== 200) {
      throw new Error("Failed to fetch event data");
    }
    return response.data;
  } catch (error) {
    console.error("Error fetching event data:", error);
    throw new Error("Failed to fetch event data");
  }
};

export const buyTicket = async (
  event_id: string,
  quantity: number,
  total_price: number,
) => {
  try {
    const userData = localStorage.getItem("user");
    if (userData) {
      const token = JSON.parse(userData).token;

      const payload = {
        event_id,
        quantity: Number(quantity), // Send as an integer
        total_price, // Send as an integer
      };

      const response = await api.post("/ticket", payload, {
        headers: {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json", // Ensure JSON is sent
        },
      });

      return response.data;
    }
  } catch (error) {
    console.error("Error buying ticket:", error);
    throw new Error("Failed to buy ticket");
  }
};

export const fetchSingleTicket = async (token: string, id: string) => {
  try {
    const response = await api.get(`/ticket/${id}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    if (response.status !== 200) {
      throw new Error("Failed to fetch event data");
    }
    return response.data;
  } catch (error) {
    console.error("Error fetching event data:", error);
    throw new Error("Failed to fetch event data");
  }
};

export const uploadSingleFile = async (form: FormData, token: string) => {
  try {
    const uploadResponse = await api.post("/storage/upload", form, {
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "multipart/form-data",
      },
    });
    return uploadResponse.data;
  } catch {}
};

export const postPembayaran = async (
  ticket_id: string,
  paymentMethod: string,
  token: string,
  path: string,
) => {
  try {
    const paymentResponse = await api.post(
      "/pembayaran/create",
      {
        ticket_id: ticket_id,
        metode_pembayaran: paymentMethod,
        bukti_bayar: path,
      },
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      },
    );

    return paymentResponse;
  } catch (error) {
    console.error("Error posting payment:", error);
    throw new Error("Failed to post payment");
  }
  
};

export const deleteUserTicket = async (token: string, ticketId: string): Promise<void> => {
  try {
    const response = await api.delete(`/ticket/${ticketId}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    if (response.status !== 200) {
      throw new Error("Failed to delete user ticket");
    }
  } catch (error) {
    console.error("Error deleting user ticket:", error);
    throw new Error("Failed to delete user ticket");
  }
};
