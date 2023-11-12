import { TableCell, TableRow } from "@/components/ui/table";
import { Item } from "@/types/course";
import formatDate from "@/utils/formatDate";
import ItensOptions from "../components/itens-options";
import { AlertCircle, ArrowUpRight, CircleEllipsis } from "lucide-react";
import useGetVideo from "@/hooks/get-video";
import { useEffect } from "react";
import { toast } from "@/components/ui/use-toast";
import { Button } from "@/components/ui/button";

const FIVE_SECONDS_IN_MS = 5000;

export default function ItemTableRow({
  item,
  refetch,
}: {
  item: Item;
  refetch: () => void;
}) {
  const { error, video, refetch: refetchVideo } = useGetVideo(item.videoId);

  useEffect(() => {
    if (error) {
      toast({
        variant: "destructive",
        title: "Não foi possível buscar video.",
        description: error,
      });
    }
  }, [error]);

  useEffect(() => {
    if (video?.status === "pending") {
      const timer = setInterval(refetchVideo, FIVE_SECONDS_IN_MS);
      return () => clearTimeout(timer);
    }
  }, [refetchVideo, video]);

  return (
    <TableRow key={item.id}>
      <TableCell className="font-medium">{item.title}</TableCell>
      <TableCell>
        {item.description.slice(0, 150)}
        {item.description.length > 150 && "..."}
      </TableCell>
      <TableCell>
        {video?.status === "pending" && (
          <div className="flex items-center">
            <CircleEllipsis className="mr-2 h-4 w-4 text-muted-foreground" />
            <span>Processando</span>
          </div>
        )}
        {video?.status === "error" && (
          <div className="flex items-center text-red-600">
            <AlertCircle className="mr-2 h-4 w-4 text-muted-foreground text-red-600" />
            <span>Não foi possível processar</span>
          </div>
        )}
        {video?.status === "success" && (
          <Button variant="outline" className="py-0 h-8" asChild>
            <a href={`/courses/itens/${item.id}`}>
              Video <ArrowUpRight className="h-4 w-4" />
            </a>
          </Button>
        )}
      </TableCell>
      <TableCell>{formatDate(item.createdAt)}</TableCell>
      <TableCell>
        <ItensOptions id={item.id} onDelete={refetch} />
      </TableCell>
    </TableRow>
  );
}
