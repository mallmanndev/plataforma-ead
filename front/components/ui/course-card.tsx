import {Card, CardDescription, CardFooter, CardHeader, CardTitle} from "@/components/ui/card";
import {Button} from "@/components/ui/button";
import React from "react";

type TCourseCardProps = {
    course: {
        id: string,
        name: string,
        description: string,
        image: string
    }
}

export default function CourseCard({course}: TCourseCardProps) {
    return (
        <Card>
            <CardHeader>
                <CardTitle>{course.name}</CardTitle>
                <CardDescription>{course.description}</CardDescription>
            </CardHeader>
            <CardFooter>
                <Button className="w-full">ACESSAR</Button>
            </CardFooter>
        </Card>
    )
}