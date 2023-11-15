import jwt from "jsonwebtoken";
import User from "@/entities/user";
import { cookies } from "next/headers";

const validateToken = (): User | null => {
  try {
    const cookiesStore = cookies();
    const token = cookiesStore.get("token");
    if (!token) {
      return null;
    }

    return jwt.verify(
      token.value,
      process.env.JWT_SECRET as string
    ) as unknown as User;
  } catch (error) {
    return null;
  }
};

export default validateToken;
