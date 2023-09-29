import CreateCourseForm from "@/app/(auth)/create-course/components/create-course-form";

export default function MyCourses() {
    return (
        <>
            <h2 className="text-3xl font-bold tracking-tight mt-12">Criar curso</h2>

            <div className="mt-12 flex-1 lg:max-w-2xl">
                <CreateCourseForm/>
            </div>

        </>
    )
}