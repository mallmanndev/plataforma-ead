import { useState } from "react";
import { signIn } from "next-auth/react";

type User = {
  id: string;
  name: string;
  email: string;
  phone: string;
};

type TLoginData = {
  email: string;
  password: string;
};

type TUseLogin = {
  loading: boolean;
  error: string | null;
  success: boolean;
  login: (data: TLoginData) => void;
};

export const useLogin = (): TUseLogin => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<boolean>(false);

  const login = async (data: TLoginData) => {
    setLoading(true);
    try {
      const response = await signIn("credentials", {
        redirect: false,
        email: data.email,
        password: data.password,
      });
      console.log(response)

      if (response?.error) {
        setError(response.error);
        return;
      }

      setSuccess(true);
    } catch (error) {
      setError("Ocorreu um erro ao criar o usuário.");
    }
    setLoading(false);
  };

  return { loading, success, error, login };
};
