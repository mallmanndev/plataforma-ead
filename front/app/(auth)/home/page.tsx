import CourseCard from "@/components/ui/course-card";

export default function Home() {

    const courses = [
        {
            id: "sfsdfsdfsdfsdfsdfsdfsdfsdfsdfsd",
            name: "Primeiros passos com Go | Plataforma EAD #01",
            description: "Neste primeiro episódio da série Plataforma EAD, iniciamos um microsserviço em GO cahamado service-core.",
            image: "/thumb.png",
        },
        {
            id: "sfsdfsdfsdfsdfsdfsdfsdfsdfsdfsd",
            name: "Primeiros passos com Go | Plataforma EAD #01",
            description: "Neste primeiro episódio da série Plataforma EAD, iniciamos um microsserviço em GO cahamado service-core.",
            image: "/thumb.png",
        },
        {
            id: "sfsdfsdfsdfsdfsdfsdfsdfsdfsdfsd",
            name: "Plataforma EAD com GO",
            description: "Neste primeiro episódio da série Plataforma EAD, iniciamos um microsserviço em GO cahamado service-core.",
            image: "/thumb.png",
        }
    ]

    return (
        <>
            <h2 className="text-3xl font-bold tracking-tight pt-4">Cursos</h2>

            <div
                className="items-start justify-center gap-6 rounded-lg pt-8 md:grid lg:grid-cols-2 xl:grid-cols-3">
                {courses.map(course => <CourseCard key={course.id} course={course}/>)}
            </div>
        </>
    )
}