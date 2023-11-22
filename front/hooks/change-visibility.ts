import { getSession } from "next-auth/react";
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
      const session = await getSession();

      const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/courses/${id}/change-visibility`,
        {
          method: "PATCH",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${session?.token}`,
          },
          body: JSON.stringify({ visibility }),
        }
      );

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
