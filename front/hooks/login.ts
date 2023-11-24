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

    const response = await signIn("credentials", {
      redirect: false,
      email: data.email,
      password: data.password,
    });

    if (response?.error) {
      setError(response.error);
    } else {
      setSuccess(true);
    }

    setLoading(false);
  };

  return { loading, success, error, login };
};
