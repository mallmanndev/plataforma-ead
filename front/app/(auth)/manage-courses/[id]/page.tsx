import { buttonVariants } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import Link from "next/link";
import validateToken from "@/lib/validate-token";
import User from "@/entities/user";
import SectionsTable from "./sections-table";

export const metadata = {
  title: "Seções",
};

export default async function MyCourses({
  params,
}: {
  params: { id: string };
}) {
  const user = validateToken() as User;

  return (
    <>
      <div className="flex items-center justify-between mt-12">
        <h2 className="text-3xl font-bold tracking-tight">Seções</h2>
        <Link
          href={`/manage-courses/${params.id}/create-section`}
          className={cn(buttonVariants({ variant: "default" }))}
        >
          Criar seção
        </Link>
      </div>

      <div className="mt-12">
        <SectionsTable userId={user.id} courseId={params.id} />
      </div>
    </>
  );
}
