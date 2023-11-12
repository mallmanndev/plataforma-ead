"use client";

import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { toast } from "@/components/ui/use-toast";
import useDeleteSection from "@/hooks/delete-section";
import { MonitorPlay, MoreHorizontal, Pencil, Trash } from "lucide-react";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

type TSectionOptionsProps = {
  id: string;
  onDelete: () => void;
};

export default function SectionOptions({ id, onDelete }: TSectionOptionsProps) {
  const { push } = useRouter();
  const { loading, error, success, remove } = useDeleteSection();

  useEffect(() => {
    if (error)
      toast({
        variant: "destructive",
        title: "Não foi possível excluir a seção",
        description: error,
      });
  }, [error]);

  useEffect(() => {
    if (success) {
      toast({ title: "Seção deletado com sucesso" });
      onDelete();
    }
  }, [onDelete, success]);

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button
          variant="ghost"
          className="flex h-8 w-8 p-0 data-[state=open]:bg-muted"
        >
          <MoreHorizontal />
          <span className="sr-only">Open menu</span>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-56">
        <DropdownMenuLabel>Opções</DropdownMenuLabel>
        <DropdownMenuSeparator />
        <DropdownMenuGroup>
          <DropdownMenuItem
            className="cursor-pointer"
            onClick={() => push(`/manage-itens/${id}`)}
          >
            <MonitorPlay className="mr-2 h-4 w-4" />
            <span>Conteúdo</span>
          </DropdownMenuItem>
          <DropdownMenuItem
            className="cursor-pointer"
            onClick={() => push(`/update-section/${id}`)}
          >
            <Pencil className="mr-2 h-4 w-4" />
            <span>Editar</span>
          </DropdownMenuItem>
          <DropdownMenuItem
            className="cursor-pointer"
            disabled={loading}
            onClick={() => remove(id)}
          >
            <Trash className="mr-2 h-4 w-4" />
            <span>Excluir</span>
          </DropdownMenuItem>
        </DropdownMenuGroup>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
