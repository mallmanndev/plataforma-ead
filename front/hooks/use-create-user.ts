import { useState } from "react";

type User = {
  id: string;
  name: string;
  email: string;
  phone: string;
};

type TCreateUserData = {
  name: string;
  email: string;
  phone: string;
  password: string;
};

type TUseCreateUser = {
  loading: boolean;
  error: string | null;
  user: User | null;
  create: (data: TCreateUserData) => void;
};

export const useCreateUser = (): TUseCreateUser => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [user, setUser] = useState<User | null>(null);

  const create = async (data: TCreateUserData) => {
    setLoading(true);
    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/user`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
      });

      if (response.ok) {
        const newUser = await response.json();
        setUser(newUser);
      } else {
        const errorData = await response.json();
        setError(errorData.message);
      }
    } catch (error) {
      setError("Ocorreu um erro ao criar o usuário.");
    }
    setLoading(false);
  };

  return { loading, user, error, create };
};
