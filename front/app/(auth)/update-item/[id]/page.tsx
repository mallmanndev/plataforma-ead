import validateToken from "@/lib/validate-token";
import CoursesServiceGrpc from "@/services/courses";
import UpdateItemForm from "./update-item-form";
import { Item } from "@/types/course";

const getItem = async (id: string) => {
  const user = validateToken();
  if (!user) {
    throw new Error("Failed to fetch section!");
  }

  const service = new CoursesServiceGrpc();
  const { error, response } = await service.GetItem({ id });

  if (error) {
    throw new Error("Failed to fetch section!");
  }

  return response;
};

export const metadata = {
  title: "Alterar Item",
};

export default async function UpdateItem({
  params,
}: {
  params: { id: string };
}) {
  const item = (await getItem(params.id)) as Item;

  return (
    <>
      <h2 className="text-3xl font-bold tracking-tight mt-12">Criar Item</h2>

      <div className="mt-12 flex-1">
        <UpdateItemForm
          id={item.id}
          initialData={{
            title: item.title,
            description: item.description,
          }}
        />
      </div>
    </>
  );
}
