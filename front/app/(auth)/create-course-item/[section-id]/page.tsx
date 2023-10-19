import CreateItemForm from "@/app/(auth)/create-course-item/[section-id]/components/create-item-form";

export default function CreateCourseItem() {
    return (
        <>
            <h2 className="text-3xl font-bold tracking-tight mt-12">Criar aula</h2>

            <div className="mt-12 flex-1 lg:max-w-2xl">
                <CreateItemForm/>
            </div>
        </>
    )
}