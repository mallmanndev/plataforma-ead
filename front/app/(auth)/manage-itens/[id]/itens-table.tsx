"use client";

import {
  Table,
  TableBody,
  TableCaption,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { useToast } from "@/components/ui/use-toast";
import { useEffect } from "react";
import EmptyItens from "../components/empty-itens";
import useGetSection from "@/hooks/get-section";
import ItemTableRow from "./item-table-row";

export default function ItensTable({ sectionId }: { sectionId: string }) {
  const { toast } = useToast();
  const { loading, section, error, refetch } = useGetSection(sectionId);

  useEffect(() => {
    if (error) {
      toast({
        variant: "destructive",
        title: "Não foi possível buscar os itens",
        description: error,
      });
    }
  }, [toast, error]);

  if (loading) {
    return <p>Carregando itens...</p>;
  }

  if (!section || section.itens.length < 1) {
    return <EmptyItens sectionId={sectionId} />;
  }

  return (
    <Table>
      <TableCaption>Gerencie seus cursos</TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead>Titulo</TableHead>
          <TableHead>Descrição</TableHead>
          <TableHead>Video</TableHead>
          <TableHead className="w-[150px]">Criado em</TableHead>
          <TableHead className="w-[50px]"></TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {section.itens.map((item) => (
          <ItemTableRow key={item.id} item={item} refetch={refetch} />
        ))}
      </TableBody>
    </Table>
  );
}
