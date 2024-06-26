"use client";

import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { useForm } from "react-hook-form";
import { Input } from "@/components/ui/input";
import * as React from "react";
import { Icons } from "@/components/ui/icons";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import { zodResolver } from "@hookform/resolvers/zod";
import { createCourseSchema } from "@/contracts/course";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { AlertCircle } from "lucide-react";
import { useToast } from "../ui/use-toast";
import { useEffect } from "react";

type Params = {
  name: string;
  description: string;
  discord_url: string;
};

type CourseForm = {
  loading: boolean;
  error: string | null;
  defaultValues: Params;
  buttonText: string;
  onSubmit: (data: Params) => void;
};

export default function CourseForm({
  defaultValues,
  loading,
  error,
  buttonText,
  onSubmit,
}: CourseForm) {
  const { toast } = useToast();
  const form = useForm<Params>({
    resolver: zodResolver(createCourseSchema),
    defaultValues: defaultValues,
  });

  useEffect(() => {
    if (error) {
      toast({
        variant: "destructive",
        title: "Não foi possível excluir o curso",
        description: error,
      });
    }
  }, [toast, error]);

  return (
    <div>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Nome*</FormLabel>
                <FormControl>
                  <Input type="text" placeholder="Nome do curso" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="description"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Descrição*</FormLabel>
                <FormControl>
                  <Textarea placeholder="Descrição do curso" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="discord_url"
            render={({ field }) => (
              <FormItem>
                <FormLabel>URL do Discord</FormLabel>
                <FormControl>
                  <Input
                    type="text"
                    placeholder="Insira uma URL para duvidas ou comunidade"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          {error && (
            <Alert variant="destructive">
              <AlertCircle className="h-4 w-4" />
              <AlertTitle>Error</AlertTitle>
              <AlertDescription>{error}</AlertDescription>
            </Alert>
          )}

          <Button type="submit" disabled={loading}>
            {loading && <Icons.spinner className="mr-2 h-4 w-4 animate-spin" />}
            {buttonText}
          </Button>
        </form>
      </Form>
    </div>
  );
}
