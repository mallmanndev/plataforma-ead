import NavBar from "@/components/ui/navbar";

export default function Layout({children}: { children: React.ReactNode }) {
    return (
        <>
            <NavBar/>
            <div className="container">
                {children}
            </div>
        </>
    )
}