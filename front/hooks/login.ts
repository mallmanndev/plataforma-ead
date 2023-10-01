import {useState} from "react";

type User = {
    id: string
    name: string
    email: string
    phone: string
}

type TLoginData = {
    email: string
    password: string
}

type TUseLogin = {
    loading: boolean
    error: string | null
    user: User | null
    login: (data: TLoginData) => void
}

export const useLogin = (): TUseLogin => {
    const [loading, setLoading] = useState<boolean>(false)
    const [error, setError] = useState<string | null>(null)
    const [user, setUser] = useState<User | null>(null)

    const login = async (data: TLoginData) => {
        setLoading(true)
        try {
            const response = await fetch('http://localhost:3002/api/login', {
                method: 'POST',
                headers: {'Content-Type': 'application/json',},
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
            setError('Ocorreu um erro ao criar o usuário.');
        }
        setLoading(false)
    }

    return {loading, user, error, login}
}