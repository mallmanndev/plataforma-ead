import { env } from "process";
import UpdateSectionForm from "./update-section-form";
import { cookies } from "next/headers";
import validateToken from "@/lib/validate-token";
import CoursesServiceGrpc from "@/services/courses";
import { Section } from "@/types/course";

const getSection = async (id: string) => {
  const user = validateToken();
  if (!user) {
    throw new Error("Failed to fetch section!");
  }

  const service = new CoursesServiceGrpc();
  const { error, response } = await service.GetSection({ id });

  if (error) {
    throw new Error("Failed to fetch section!");
  }

  return response as Section;
};

export const metadata = {
  title: "Alterar Seção",
};

export default async function CreateSection({
  params,
}: {
  params: { id: string };
}) {
  const section = await getSection(params.id);

  return (
    <>
      <h2 className="text-3xl font-bold tracking-tight mt-12">Criar Seção</h2>

      <div className="mt-12 flex-1">
        <UpdateSectionForm id={params.id} initialData={section} />
      </div>
    </>
  );
}
