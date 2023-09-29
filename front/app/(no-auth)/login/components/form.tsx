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
const loginSchema = z.object({
    email: z.string({required_error}).email({message: "Email inválido!"}),
    password: z.string({required_error})
        .min(8, {message: "A senha deve conter mais de 8 digitos!"})
})

type TLoginSchema = z.infer<typeof loginSchema>

export function LoginForm({className, ...props}: UserAuthFormProps) {
    const [isLoading, setIsLoading] = React.useState<boolean>(false)
    const form = useForm<TLoginSchema>({
        resolver: zodResolver(loginSchema),
    })

    async function onSubmit(data: TLoginSchema) {
        console.log(data)
        setIsLoading(true)

        setTimeout(() => {
            setIsLoading(false)
        }, 3000)
    }

    return (
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
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
                    Login
                </Button>
            </form>
        </Form>
    )
}