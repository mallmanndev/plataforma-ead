import { useState } from "react";

type TUseDeleteCourse = {
  error?: string;
  loading: boolean;
  success: boolean;
  remove: (id: string) => Promise<void>;
};

export default function useDeleteCourse(): TUseDeleteCourse {
  const [loading, setLoading] = useState(false);
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState<string>();

  const remove = async (id: string) => {
    setLoading(true);
    const res = await fetch(`/api/courses/${id}`, {
      method: "DELETE",
      headers: { "Content-Type": "application/json" },
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
