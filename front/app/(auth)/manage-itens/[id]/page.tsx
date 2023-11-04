import { buttonVariants } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import Link from "next/link";
import ItensTable from "./itens-table";

export const metadata = {
  title: "Itens",
};

export default async function MyCourses({
  params,
}: {
  params: { id: string };
}) {
  return (
    <>
      <div className="flex items-center justify-between mt-12">
        <h2 className="text-3xl font-bold tracking-tight">Itens</h2>
        <Link
          href={`/create-item/${params.id}`}
          className={cn(buttonVariants({ variant: "default" }))}
        >
          Criar item
        </Link>
      </div>

      <div className="mt-12">
        <ItensTable sectionId={params.id} />
      </div>
    </>
  );
}
