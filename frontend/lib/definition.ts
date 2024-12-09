export interface ApiResponse<T> {
    status: boolean;
    success: boolean;
    message: string;
    data: T;
    meta?: {
      page: number;
      per_page: number;
      max_page: number;
      count: number;
    };
  }

export interface Event {
    id: string;
    name: string;
    date: Date;
    pricing: number;
    isactive: boolean;
    quota: number;
    typeid: string;
    type: Type;
    timestamp: Timestamp;
}

// Define the Type interface as well
export interface Type {
    id: string;
    name: string; // Add properties of Type interface here
}

// Define the Timestamp interface if needed
export interface Timestamp {
    createdat: string;
	updatedat: string;
	deletedat: string;
}

export interface Pembayaran {
    id: string;
    ticket_id: string;
    metode_pembayaran: string;
    bukti_bayar: string;
    is_verified: boolean;
    timestamp: Timestamp;
}

export interface User {
    id: string;
    name: string;
    phone_number: string;
    email: string;
    password: string;
    role: string;
    timestamp: Timestamp;
}

export interface UserTicket {
    id: string;
    user_id: string;
    event_id: string;
    quantity: number;
    total_price: number;
}