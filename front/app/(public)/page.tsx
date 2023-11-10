import validateToken from "@/lib/validate-token";
import { redirect } from "next/navigation";

export default function RootPage() {
  const user = validateToken();
  if (!user) return redirect("/login");
  return redirect("/home");
}
