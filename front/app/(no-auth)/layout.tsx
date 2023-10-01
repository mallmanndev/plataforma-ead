import {redirect} from 'next/navigation'
import validateToken from "@/lib/validate-token";

export default function Layout({children}: { children: React.ReactNode }) {
    const user = validateToken()

    if (user) return redirect('/home')

    return children
}