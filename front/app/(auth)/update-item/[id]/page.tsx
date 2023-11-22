import UpdateItemForm from "./update-item-form";
import { Item } from "@/types/course";
import { getServerSession } from "next-auth/next";
import { nextAuthOptions } from "@/app/api/auth/[...nextauth]/route";

const getItem = async (id: string) => {
  const session = await getServerSession(nextAuthOptions);

  const res = await fetch(`${process.env.SERVER_HOST}/itens/${id}`, {
    headers: { Authorization: `Bearer ${session?.token}` },
  });

  if (!res.ok) {
    throw new Error("Failed to fetch item!");
  }

  return res.json();
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
