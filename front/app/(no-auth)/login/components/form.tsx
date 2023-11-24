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
import { useLogin } from "@/hooks/login";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { AlertCircle } from "lucide-react";
import { useEffect } from "react";
import { useRouter } from "next/navigation";

const required_error = "Este campo é obrigatório.";
const loginSchema = z.object({
  email: z.string({ required_error }).email({ message: "Email inválido!" }),
  password: z
    .string({ required_error })
    .min(8, { message: "A senha deve conter mais de 8 digitos!" }),
});

type TLoginSchema = z.infer<typeof loginSchema>;

export function LoginForm() {
  const { replace } = useRouter();
  const { loading, login, error, success } = useLogin();
  const form = useForm<TLoginSchema>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  useEffect(() => {
    if (success) replace(`/home`);
  }, [replace, success]);

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(login)} className="space-y-4">
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

        <Button type="submit" className="w-full" disabled={loading}>
          {loading && <Icons.spinner className="mr-2 h-4 w-4 animate-spin" />}
          Login
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
function setTokenCoockie(token: string) {
  throw new Error("Function not implemented.");
}
