import CreateItemForm from "./create-item-form";

export const metadata = {
  title: "Criar Item",
};

export default function CreateItem({ params }: { params: { id: string } }) {
  return (
    <>
      <h2 className="text-3xl font-bold tracking-tight mt-12">Criar Item</h2>

      <div className="mt-12 flex-1">
        <CreateItemForm sectionId={params.id} />
      </div>
    </>
  );
}
