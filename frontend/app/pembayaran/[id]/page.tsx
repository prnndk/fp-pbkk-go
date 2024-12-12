"use client";

import React, { useEffect, useState } from "react";
import { useRouter, useParams } from "next/navigation";
import axios from "axios";
import NotFound from "@/components/NotFound";
import { Button } from "@/components/ui/button";
import { z } from "zod";
import { FormProvider, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  FormField,
  FormItem,
  FormMessage,
  FormLabel,
  FormControl,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { fetchSingleTicket, postPembayaran, uploadSingleFile } from "@/lib/api";

const formSchema = z.object({
  paymentMethod: z.string().nonempty("Payment method is required"),
  file: z.instanceof(File).refine((file) => file.size > 0, {
    message: "File is required",
  }),
});

const Pembayaran = () => {
  const { id } = useParams() as { id: string };
  const router = useRouter();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      paymentMethod: "",
      file: undefined,
    },
  });

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

  const handleReturn = () => {
    router.push("/");
  };

  if (!data) {
    return <NotFound />;
  }

  const handleSubmit = async (formData: z.infer<typeof formSchema>) => {
    try {
      // Upload file
      const fileData = new FormData();
      fileData.append("file", formData.file);

      if (token) {
        uploadSingleFile(fileData, token).then(async (response) => {
          const path = response.data.path;
          const paymentResponse = await postPembayaran(
            id,
            formData.paymentMethod,
            token,
            path,
          );

          if (paymentResponse.status === 200) {
            alert("Payment successful");
            router.push("/");
          } else {
            alert("Payment failed");
          }
        });
      } else {
        throw new Error("Token is null");
      }

      // Submit payment details
    } catch (error) {
      console.error("Error submitting payment:", error);
      alert("An error occurred while submitting the payment");
    }
  };

  return (
    <div className="flex h-screen items-center justify-center">
      <div className="w-96 border border-black bg-white p-10 text-black">
        <h1 className="mb-4 text-2xl font-bold">
          Pembayaran Untuk {data.event.name}
        </h1>
        <p>Jumlah Tiket: {data.quantity}</p>
        <p>Total Harga: Rp {data.total_price.toLocaleString("id-ID")}</p>
        <FormProvider {...form}>
          <form
            onSubmit={form.handleSubmit(handleSubmit)}
            className="space-y-2"
          >
            <FormField
              control={form.control}
              name="paymentMethod"
              render={({ field, fieldState }) => (
                <FormItem>
                  <FormLabel>Metode Pembayaran</FormLabel>
                  <FormControl>
                    <Input
                      type="text"
                      placeholder="Payment Method"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage>{fieldState.error?.message}</FormMessage>
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="file"
              render={({ field, fieldState }) => (
                <FormItem>
                  <FormLabel>Bukti Bayar</FormLabel>
                  <FormControl>
                    <Input
                      type="file"
                      onChange={(e) => {
                        if (e.target.files) {
                          field.onChange(e.target.files[0]);
                        }
                      }}
                    />
                  </FormControl>
                  <FormMessage>{fieldState.error?.message}</FormMessage>
                </FormItem>
              )}
            />
            <Button
              type="submit"
              className="mb-4 w-full bg-green-400 text-white hover:bg-green-700"
              variant="default"
            >
              Bayar Tiket
            </Button>
          </form>
        </FormProvider>
        <Button onClick={handleReturn} className="my-2 w-full">
          Back to Home
        </Button>
      </div>
    </div>
  );
};

export default Pembayaran;
