import { env } from "process";
import UpdateSectionForm from "./update-section-form";

const getCourse = async (id: string) => {
  const res = await fetch(`${env.SERVER_HOST}/api/courses/${id}`, {
    cache: "no-cache",
  });

  if (!res.ok) {
    throw new Error("Failed to fetch course!");
  }

  return res.json();
};

export const metadata = {
  title: "Alterar Seção",
};

export default function CreateSection({ params }: { params: { id: string } }) {
  return (
    <>
      <h2 className="text-3xl font-bold tracking-tight mt-12">Criar Seção</h2>

      <div className="mt-12 flex-1">
        <UpdateSectionForm id={params.id} />
      </div>
    </>
  );
}
