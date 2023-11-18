import { buttonVariants } from "@/components/ui/button";
import React from "react";
import { cn } from "@/lib/utils";
import Link from "next/link";
import { MonitorPlay } from "lucide-react";

export default function EmptItens({ sectionId }: { sectionId: string }) {
  return (
    <div className="flex h-[450px] shrink-0 items-center justify-center rounded-md border border-dashed">
      <div className="mx-auto flex max-w-[420px] flex-col items-center justify-center text-center">
        <MonitorPlay size={40} />
        <h3 className="mt-4 text-lg font-semibold">Nenhum item cadastrado</h3>
        <p className="mb-4 mt-2 text-sm text-muted-foreground">
          Comece a criar suas seções clicando no botão abaixo.
        </p>
        <Link
          href={`/create-item/${sectionId}`}
          className={cn(buttonVariants({ variant: "default" }))}
        >
          Criar item
        </Link>
      </div>
    </div>
  );
}
