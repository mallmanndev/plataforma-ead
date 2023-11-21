import { redirect } from "next/navigation";
import { getServerSession } from "next-auth";
import { nextAuthOptions } from "../api/auth/[...nextauth]/route";
import NavBar from "@/components/ui/navbar";
import User from "@/entities/user";

export default async function Layout({ children }: { children: React.ReactNode }) {
  const session = await getServerSession(nextAuthOptions)

  if (!session) return redirect("/login");

  return (
    <>
      <NavBar user={session as unknown as User} />
      <div className="container">{children}</div>
    </>
  );
}
