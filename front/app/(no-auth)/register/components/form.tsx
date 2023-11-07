"use client";

import * as React from "react";

import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Icons } from "@/components/ui/icons";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useRouter } from "next/navigation";
import { useCreateUser } from "@/hooks/use-create-user";
import { useEffect } from "react";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { AlertCircle } from "lucide-react";

const required_error = "Este campo é obrigatório.";
const registerSchema = z
  .object({
    name: z.string({ required_error }).min(5, "Mínimo 5 caracteres."),
    phone: z.string({ required_error }).min(9, "Mínimo 9 caracteres."),
    email: z.string({ required_error }).email({ message: "Email inválido!" }),
    password: z
      .string({ required_error })
      .min(8, { message: "A senha deve conter mais de 8 digitos!" }),
    cpassword: z.string({ required_error }),
  })
  .refine((data) => data.password === data.cpassword, {
    message: "As senhas não conferem!",
    path: ["cpassword"],
  });

type TRegisterSchema = z.infer<typeof registerSchema>;

export function RegisterForm() {
  const { push } = useRouter();
  const { loading, error, user, create } = useCreateUser();
  const form = useForm<TRegisterSchema>({
    resolver: zodResolver(registerSchema),
  });

  useEffect(() => {
    if (user) push("/login");
  }, [user]);

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(create)} className="space-y-4">
        <FormField
          control={form.control}
          name="name"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Nome</FormLabel>
              <FormControl>
                <Input placeholder="Seu nome completo" {...field} />
              </FormControl>
              <FormMessage id="name-message" />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="phone"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Telefone</FormLabel>
              <FormControl>
                <Input type="text" placeholder="Seu telefone" {...field} />
              </FormControl>
              <FormMessage id="phone-message" />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="email"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Email</FormLabel>
              <FormControl>
                <Input
                  type="email"
                  placeholder="Seu email principal"
                  {...field}
                />
              </FormControl>
              <FormMessage id="email-message" />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="password"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Senha</FormLabel>
              <FormControl>
                <Input
                  type="password"
                  placeholder="Insira uma senha"
                  {...field}
                />
              </FormControl>
              <FormMessage id="password-message" />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="cpassword"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Confirmar senha</FormLabel>
              <FormControl>
                <Input
                  type="password"
                  placeholder="Confirme sua senha"
                  {...field}
                />
              </FormControl>
              <FormMessage id="cpassword-message" />
            </FormItem>
          )}
        />

        <Button type="submit" className="w-full" disabled={loading}>
          {loading && <Icons.spinner className="mr-2 h-4 w-4 animate-spin" />}
          Criar conta
        </Button>

        {error && (
          <Alert variant="destructive">
            <AlertCircle className="h-4 w-4" />
            <AlertDescription id="error-alert">{error}</AlertDescription>
          </Alert>
        )}
      </form>
    </Form>
  );
}
