import UpdateCourseForm from "./update-course-form";

const getCourse = async (id: string) => {
  const res = await fetch(`${process.env.SERVER_HOST}/api/courses/${id}`, {
    cache: "no-cache",
  });

  if (!res.ok) {
    throw new Error("Failed to fetch course!");
  }

  return res.json();
};

export const metadata = {
  title: "Editar curso",
};

export default async function MyCourses({
  params,
}: {
  params: { id: string };
}) {
  const course = await getCourse(params.id);

  return (
    <>
      <h2 className="text-3xl font-bold tracking-tight mt-12">Editar Curso</h2>

      <div className="mt-12 flex-1">
        <UpdateCourseForm
          id={params.id}
          initialData={{ name: course.name, description: course.description }}
        />
      </div>
    </>
  );
}
