"use client";

import React, { useEffect, useState } from "react";
import { useRouter, useParams } from "next/navigation";
import { buyTicket, fetchEventById } from "@/lib/api";
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

const formSchema = z.object({
  quantity: z.number().int().positive().min(1).max(10),
});

const EventSquare = () => {
  const { id } = useParams() as { id: string };
  const router = useRouter();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      quantity: 1,
    },
  });

  const [data, setData] = useState<any | null>(null);
  const [quantity, setQuantity] = useState<number>(1);

  const handleQuantityChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = parseInt(e.target.value) || 1; // Default to 1 if input is invalid
    setQuantity(value);
    form.setValue("quantity", value, { shouldValidate: true });
  };

  useEffect(() => {
    const userData = localStorage.getItem("user");
    if (userData) {
      const token = JSON.parse(userData).token;
      fetchEventById(token, id)
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

  const handleReturn = () => {
    router.push("/");
  };

  const handleBuyTicket = () => {
    const formData = form.getValues();

    buyTicket(
      id,
      formData.quantity,
      data.pricing * (quantity <= 10 ? quantity : 10),
    )
      .then(() => {
        alert("Ticket purchased successfully!");
        router.push("/");
      })
      .catch((error) => {
        console.error("Failed to purchase ticket:", error);
        alert("Failed to purchase ticket. Please try again.");
      });
  };

  return (
    <div className="flex h-screen items-center justify-center">
      <div className="w-96 border border-black bg-white p-10 text-black">
        <h2>{data.name}</h2>
        <br />
        <p>Tickets Available: {data.quota}</p>
        <p>
          Base Price:{" "}
          {new Intl.NumberFormat("id-ID", {
            style: "currency",
            currency: "IDR",
          }).format(data.pricing)}
        </p>
        <FormProvider {...form}>
          <form
            onSubmit={form.handleSubmit(handleBuyTicket)}
            className="space-y-8"
          >
            <FormField
              control={form.control}
              name="quantity"
              render={({ field, fieldState }) => (
                <FormItem>
                  <FormLabel>Quantity</FormLabel>
                  <FormControl>
                    <Input
                      type="number"
                      placeholder="Quantity"
                      {...field}
                      value={quantity}
                      onChange={(e) => {
                        handleQuantityChange(e);
                        field.onChange(e); // Ensure integration with react-hook-form
                      }}
                    />
                  </FormControl>
                  <FormMessage>{fieldState.error?.message}</FormMessage>
                </FormItem>
              )}
            />
          </form>
        </FormProvider>
        <p className="my-3">
          Total Price:{" "}
          {new Intl.NumberFormat("id-ID", {
            style: "currency",
            currency: "IDR",
          }).format(data.pricing * (quantity <= 10 ? quantity : 10))}
        </p>
        <div className="mt-4 flex flex-row space-x-4">
          <Button
            variant="outline"
            onClick={handleReturn}
            className="text-white"
          >
            Back To Home
          </Button>
          <Button
            type="submit"
            onClick={handleBuyTicket}
            className="bg-blue-500 text-white hover:bg-blue-700"
          >
            Beli Tiket
          </Button>
        </div>
      </div>
    </div>
  );
};

export default EventSquare;
