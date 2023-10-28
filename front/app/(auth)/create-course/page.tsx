import CreateCourseForm from "./create-course-form";

export const metadata = {
  title: "Criar curso",
};

export default function MyCourses() {
  return (
    <>
      <h2 className="text-3xl font-bold tracking-tight mt-12">Criar Curso</h2>

      <div className="mt-12 flex-1">
        <CreateCourseForm />
      </div>
    </>
  );
}
