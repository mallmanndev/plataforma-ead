"use client"

import * as React from "react"

import {cn} from "@/lib/utils"
import {Input} from "@/components/ui/input";
import {Button, buttonVariants} from "@/components/ui/button";
import Link from "next/link";
import {Icons} from "@/components/ui/icons";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {useForm} from "react-hook-form";
import {z} from "zod";
import {zodResolver} from "@hookform/resolvers/zod";

interface UserAuthFormProps extends React.HTMLAttributes<HTMLDivElement> {
}

const required_error = "Este campo é obrigatório."
const registerSchema = z.object({
    name: z.string({required_error}).min(5, "Mínimo 5 caracteres."),
    phone: z.string({required_error}).min(9, "Mínimo 9 caracteres."),
    email: z.string({required_error}).email({message: "Email inválido!"}),
    password: z.string({required_error})
        .min(8, {message: "A senha deve conter mais de 8 digitos!"})
})

type TRegisterSchema = z.infer<typeof registerSchema>

export function RegisterForm({className, ...props}: UserAuthFormProps) {
    const [isLoading, setIsLoading] = React.useState<boolean>(false)
    const form = useForm<TRegisterSchema>({resolver: zodResolver(registerSchema)})

    const onSubmit = (data: TRegisterSchema) => {
        console.log(data)
    }

    return (
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
                <FormField
                    control={form.control}
                    name="name"
                    render={({field}) => (
                        <FormItem>
                            <FormLabel>Nome</FormLabel>
                            <FormControl>
                                <Input placeholder="Seu nome completo" {...field} />
                            </FormControl>
                            <FormMessage/>
                        </FormItem>
                    )}
                />

                <FormField
                    control={form.control}
                    name="phone"
                    render={({field}) => (
                        <FormItem>
                            <FormLabel>Telefone</FormLabel>
                            <FormControl>
                                <Input type="text" placeholder="Seu telefone" {...field} />
                            </FormControl>
                            <FormMessage/>
                        </FormItem>
                    )}
                />

                <FormField
                    control={form.control}
                    name="email"
                    render={({field}) => (
                        <FormItem>
                            <FormLabel>Email</FormLabel>
                            <FormControl>
                                <Input type="email" placeholder="Seu email principal" {...field} />
                            </FormControl>
                            <FormMessage/>
                        </FormItem>
                    )}
                />

                <FormField
                    control={form.control}
                    name="password"
                    render={({field}) => (
                        <FormItem>
                            <FormLabel>Senha</FormLabel>
                            <FormControl>
                                <Input type="password" placeholder="Insira uma senha" {...field} />
                            </FormControl>
                            <FormMessage/>
                        </FormItem>
                    )}
                />

                <Button type="submit" className="w-full" disabled={isLoading}>
                    {isLoading && (
                        <Icons.spinner className="mr-2 h-4 w-4 animate-spin"/>
                    )}
                    Criar conta
                </Button>
            </form>
        </Form>
    )
}