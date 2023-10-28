import NavBar from "@/components/ui/navbar";
import { redirect } from "next/navigation";
import validateToken from "@/lib/validate-token";

export default function Layout({ children }: { children: React.ReactNode }) {
  const user = validateToken();

  if (!user) return redirect("/login");

  return (
    <>
      <NavBar user={user} />
      <div className="container">{children}</div>
    </>
  );
}
