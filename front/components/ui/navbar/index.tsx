import Link from "next/link";
import {cn} from "@/lib/utils";
import React from "react";
import {DarkModeToggle} from "@/components/ui/dark-mode";
import NavUser from "@/components/ui/navbar/user";
import User from "@/entities/user";


export default function NavBar({user}: { user: User }) {
    return (
        <header className="border-b">
            <div className="container flex h-14 items-center">
                <nav
                    className={cn("flex items-center space-x-4 lg:space-x-6")}
                >
                    <Link
                        href="/home"
                        className="text-sm font-medium transition-colors hover:text-primary"
                    >
                        Home
                    </Link>

                    <Link
                        href="/manage-courses"
                        className="text-sm font-medium transition-colors hover:text-primary"
                    >
                        Meus cursos
                    </Link>
                </nav>

                <div className="ml-auto flex items-center space-x-4">
                    <DarkModeToggle/>
                    <NavUser user={user}/>
                </div>
            </div>
        </header>
    )
}

