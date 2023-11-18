"use client";

import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useToast } from "@/components/ui/use-toast";
import useChangeVisibility from "@/hooks/change-visibility";
import { useEffect, useState } from "react";

type TVisibilitySelectProps = {
  id: string;
  isVisible: boolean;
};

export default function VisibilitySelect({
  id,
  isVisible,
}: TVisibilitySelectProps) {
  const { toast } = useToast();
  const [value, setValue] = useState(isVisible ? "public" : "private");
  const { changeVisibility, loading, error, success } = useChangeVisibility();

  const handleChange = (value: string) => {
    setValue(value);
    changeVisibility(id, value);
  };

  useEffect(() => {
    if (error)
      toast({
        variant: "destructive",
        title: "Não foi possível alterar a visibilidade do curso",
        description: error,
      });
  }, [error, toast]);

  useEffect(() => {
    if (success) toast({ title: "Visibilidade alterada com sucesso" });
  }, [success, toast]);

  return (
    <Select disabled={loading} onValueChange={handleChange} value={value}>
      <SelectTrigger className="h-8">
        <SelectValue placeholder="Visibilidade" />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          <SelectItem value="public">Público</SelectItem>
          <SelectItem value="private">Privado</SelectItem>
        </SelectGroup>
      </SelectContent>
    </Select>
  );
}
