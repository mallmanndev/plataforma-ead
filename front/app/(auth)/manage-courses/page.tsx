import { buttonVariants } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import Link from "next/link";
import CoursesTable from "./courses-table";
import { getServerSession } from "next-auth";
import { nextAuthOptions } from "@/app/api/auth/[...nextauth]/route";
import User from "@/entities/user";

export const metadata = {
  title: "Meus cursos",
};

export default async function MyCourses() {
  const user = (await getServerSession(nextAuthOptions)) as User;

  return (
    <>
      <div className="flex items-center justify-between mt-12">
        <h2 className="text-3xl font-bold tracking-tight">Meus cursos</h2>
        <Link
          href="/create-course"
          className={cn(buttonVariants({ variant: "default" }))}
        >
          Criar curso
        </Link>
      </div>

      <div className="mt-12">
        <CoursesTable userId={user.id} />
      </div>
    </>
  );
}
