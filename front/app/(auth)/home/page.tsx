import CourseCard from "@/components/ui/course-card";
import CoursesServiceGrpc from "@/services/courses";

export const metadata = {
  title: "Home",
};

const getCourses = async () => {
  const service = new CoursesServiceGrpc();
  const { error, response } = await service.Get({
    id: "",
    user_id: "",
    visible: true,
  });

  if (error) {
    throw new Error("Failed to fetch course!");
  }

  return response;
};

export default async function Home() {
  const courses = await getCourses();

  return (
    <>
      <h2 className="text-3xl font-bold tracking-tight pt-4">Cursos</h2>

      <div className="items-start justify-center gap-6 rounded-lg pt-8 md:grid lg:grid-cols-2 xl:grid-cols-3">
        {courses?.map((course) => (
          <CourseCard key={course.id} course={course} />
        ))}
      </div>
    </>
  );
}
