import { getSession } from "next-auth/react";
import { useState } from "react";

type TUseDeleteItem = {
  error?: string;
  loading: boolean;
  success: boolean;
  remove: (id: string) => Promise<void>;
};

export default function useDeleteItem(): TUseDeleteItem {
  const [loading, setLoading] = useState(false);
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState<string>();

  const remove = async (id: string) => {
    setLoading(true);
    const session = await getSession();

    const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/itens/${id}`, {
      method: "DELETE",
      headers: { "Content-Type": "application/json", Authorization: `Bearer ${session?.token}`, },
    });

    if (res.ok) {
      setSuccess(true);
    } else {
      const errorData = await res.json();
      setError(errorData.message);
    }
    setLoading(false);
  };

  return { loading, error, success, remove };
}
