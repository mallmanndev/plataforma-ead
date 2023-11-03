import * as React from "react";
import Link from "next/link";
import {LoginForm} from "@/app/(no-auth)/login/components/form";
import LoginRegisterLayout from "@/components/layouts/LoginRegisterLayout";
import {cn} from "@/lib/utils";
import {buttonVariants} from "@/components/ui/button";

export const metadata = {
    title: "Login"
}

export default function Login() {
    return (
        <>
            <LoginRegisterLayout button={
                <Link
                    href="/register"
                    className={cn(
                        buttonVariants({variant: "outline"}),
                        "absolute right-4 top-4 md:right-8 md:top-8"
                    )}
                >
                    Cadastre-se
                </Link>
            }>
                <div className="mx-auto flex w-full flex-col justify-center space-y-6 sm:w-[350px]">
                    <div className="flex flex-col space-y-2 text-center">
                        <h1 className="text-2xl font-semibold tracking-tight">
                            Login
                        </h1>
                        <p className="text-sm text-muted-foreground">
                            Acesse sua conta com email e senha!
                        </p>
                    </div>
                    <div className={cn("grid gap-6")}>
                        <LoginForm/>
                        <div className="relative">
                            <div className="absolute inset-0 flex items-center">
                                <span className="w-full border-t"/>
                            </div>
                            <div className="relative flex justify-center text-xs uppercase">
                    <span className="bg-background px-2 text-muted-foreground">
                        Ou
                    </span>
                            </div>
                        </div>

                        <Link
                            href="/register"
                            className={cn(buttonVariants({variant: "outline"}))}
                        >
                            Cadastre-se
                        </Link>

                    </div>
                    <p className="px-8 text-center text-sm text-muted-foreground">
                        Ao continuar, você está de acordo com nossos{" "}
                        <Link
                            href="/terms"
                            className="underline underline-offset-4 hover:text-primary"
                        >
                            Termos de uso
                        </Link>
                        {" "}e{" "}
                        <Link
                            href="/privacy"
                            className="underline underline-offset-4 hover:text-primary"
                        >
                            Política de privacidade
                        </Link>
                        .
                    </p>
                </div>
            </LoginRegisterLayout>
        </>
    )
}