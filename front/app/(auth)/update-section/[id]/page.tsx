import { nextAuthOptions } from "@/app/api/auth/[...nextauth]/route";
import UpdateSectionForm from "./update-section-form";
import { getServerSession } from "next-auth";

const getSection = async (id: string) => {
  const session = await getServerSession(nextAuthOptions);

  const res = await fetch(`${process.env.SERVER_HOST}/sections/${id}`, {
    headers: { Authorization: `Bearer ${session?.token}` },
  });

  if (!res.ok) {
    throw new Error("Failed to fetch section!");
  }

  return res.json();
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
