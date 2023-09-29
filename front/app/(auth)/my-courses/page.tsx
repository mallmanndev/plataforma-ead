import {EmptyCourses} from "@/app/(auth)/my-courses/components/empty-courses";

export default function MyCourses() {
    return (
        <>
            <h2 className="text-3xl font-bold tracking-tight mt-12">Meus cursos</h2>

            <div className="mt-12">
                <EmptyCourses/>
            </div>

        </>
    )
}