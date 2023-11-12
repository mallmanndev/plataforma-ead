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
import {
  TCreateCourseItemData,
  createCourseItemSchema,
} from "@/contracts/course";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { AlertCircle } from "lucide-react";
import { useToast } from "../ui/use-toast";
import { useEffect } from "react";

type TCreateData = Omit<TCreateCourseItemData, "user_id" | "section_id">;
type TUpdateData = Omit<
  TCreateCourseItemData,
  "user_id" | "section_id" | "video_id"
>;

type ItemForm = {
  loading: boolean;
  error: string | null;
  defaultValues: TCreateData | TUpdateData;
  buttonText: string;
  onSubmit: (data: TCreateData | TUpdateData) => void;
};

export default function ItemForm({
  defaultValues,
  loading,
  error,
  buttonText,
  onSubmit,
}: ItemForm) {
  const { toast } = useToast();
  const form = useForm<TCreateData | TUpdateData>({
    resolver: zodResolver(
      createCourseItemSchema.omit({
        user_id: true,
        section_id: true,
        video_id: true,
      })
    ),
    defaultValues: defaultValues,
  });

  useEffect(() => {
    if (error)
      toast({
        variant: "destructive",
        title: "Não foi possível criar item",
        description: error,
      });
  }, [toast, error]);

  return (
    <div>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
          <FormField
            control={form.control}
            name="title"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Titulo*</FormLabel>
                <FormControl>
                  <Input type="text" placeholder="Titulo do video" {...field} />
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
                  <Textarea placeholder="Descrição do video" {...field} />
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
