"use client";

import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";
import { useRouter } from "next/navigation";

export default function Sections({ sections }: { sections: any[] }) {
  const router = useRouter();

  const handleItemClick = (item: any) => {
    router.push(`?item=${item.id}`);
  };

  return (
    <>
      {sections.map((section: any) => (
        <Accordion key={section.id} type="single" collapsible>
          <AccordionItem value="item-1">
            <AccordionTrigger>{section.name}</AccordionTrigger>
            <AccordionContent>
              <ul>
                {section.itens.map((item: any) => (
                  <li
                    key={item.id}
                    onClick={() => handleItemClick(item)}
                    className="px-4 py-2 cursor-pointer"
                  >
                    {item.title}
                  </li>
                ))}
              </ul>
            </AccordionContent>
          </AccordionItem>
        </Accordion>
      ))}
    </>
  );
}
