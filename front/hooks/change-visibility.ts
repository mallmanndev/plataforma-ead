import { useState } from "react";

type TUseChangeVisibility = {
  changeVisibility: (id: string, visibility: string) => void;
  loading: boolean;
  error: string | null;
  success: boolean;
};

export default function useChangeVisibility(): TUseChangeVisibility {
  const [loading, setLoading] = useState(false);
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const changeVisibility = (id: string, visibility: string) => {
    setLoading(true);
    setError(null);
    setSuccess(false);

    (async () => {
      const response = await fetch(`/api/courses/change-visibility/${id}`, {
        method: "PATCH",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ visibility }),
      });

      if (response.ok) {
        setLoading(false);
        setSuccess(true);
      } else {
        const errorData = await response.json();
        setError(errorData.message);
      }
    })();
  };

  return { changeVisibility, loading, error, success };
}
