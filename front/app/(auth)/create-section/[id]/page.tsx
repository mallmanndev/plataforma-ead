import SectionForm from "@/components/forms/section-form";
import CreateSectionForm from "./create-course-form";

export const metadata = {
  title: "Criar Seção",
};

export default function CreateSection({ params }: { params: { id: string } }) {
  return (
    <>
      <h2 className="text-3xl font-bold tracking-tight mt-12">Criar Seção</h2>

      <div className="mt-12 flex-1">
        <CreateSectionForm courseId={params.id} />
      </div>
    </>
  );
}
