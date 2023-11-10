import { useState } from "react";

type TUseLogout = {
  loading: boolean;
  error: string | null;
  finished: boolean;
  logout: () => void;
};

const useLogout = (): TUseLogout => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [finished, setFinish] = useState<boolean>(false);

  const logout = async () => {
    setLoading(true);
    try {
      const response = await fetch(
        `${process.env.NEXT_PUBLIC_SERVER_HOST}/api/logout`,
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
        }
      );

      if (response.ok) {
        setFinish(true);
      } else {
        const errorData = await response.json();
        setError(errorData.message);
      }
    } catch (error) {
      setError("Ocorreu um erro ao sair.");
    }
    setLoading(false);
  };

  return { loading, finished, error, logout };
};

export default useLogout;
