import { useState } from "react";

type TUseDeleteSection = {
  error?: string;
  loading: boolean;
  success: boolean;
  remove: (id: string) => Promise<void>;
};

export default function useDeleteSection(): TUseDeleteSection {
  const [loading, setLoading] = useState(false);
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState<string>();

  const remove = async (id: string) => {
    setLoading(true);
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_SERVER_HOST}/api/sections/${id}`,
      {
        method: "DELETE",
        headers: { "Content-Type": "application/json" },
      }
    );

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
