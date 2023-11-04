"use client";

import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import formatDate from "@/utils/formatDate";
import { useToast } from "@/components/ui/use-toast";
import { useEffect } from "react";
import EmptyItens from "../components/empty-itens";
import ItensOptions from "../components/itens-options";
import useGetSection from "@/hooks/get-section";

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
  }, [error]);

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
          <TableHead className="w-[150px]">Criado em</TableHead>
          <TableHead className="w-[50px]"></TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {section.itens.map((item) => (
          <TableRow key={item.id}>
            <TableCell className="font-medium">{item.title}</TableCell>
            <TableCell>{item.description}</TableCell>
            <TableCell>{formatDate(item.createdAt)}</TableCell>
            <TableCell>
              <ItensOptions id={item.id} onDelete={refetch} />
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
}
