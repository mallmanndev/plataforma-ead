'use client'

import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger
} from "@/components/ui/dropdown-menu";
import {Button} from "@/components/ui/button";
import {Avatar, AvatarFallback, AvatarImage} from "@/components/ui/avatar";
import React, {useEffect} from "react";
import User from "@/entities/user";
import getNameInitials from "@/lib/name-initials";
import useLogout from "@/hooks/logout";
import {useRouter} from "next/navigation";

export default function NavUser({user}: { user: User }) {
    const {push} = useRouter()
    const {logout, finished} = useLogout()

    useEffect(() => {
        if (finished) push('/login')
    }, [finished])

    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild id="user-avatar">
                <Button variant="ghost" className="relative h-8 w-8 rounded-full">
                    <Avatar className="h-8 w-8">
                        <AvatarImage src="/avatars/01.png" alt="@shadcn"/>
                        <AvatarFallback id="sigla-avatar">{getNameInitials(user.name)}</AvatarFallback>
                    </Avatar>
                </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent className="w-56" align="end" forceMount>
                <DropdownMenuLabel className="font-normal">
                    <div className="flex flex-col space-y-1">
                        <p id="user-name" className="text-sm font-medium leading-none">{user.name}</p>
                        <p id="user-email" className="text-xs leading-none text-muted-foreground">{user.email}</p>
                    </div>
                </DropdownMenuLabel>
                <DropdownMenuSeparator/>
                <DropdownMenuItem id="logout" onClick={logout}>Sair</DropdownMenuItem>
            </DropdownMenuContent>
        </DropdownMenu>
    )
}