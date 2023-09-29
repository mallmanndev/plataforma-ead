"use client";

import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {useForm} from "react-hook-form";
import {Input} from "@/components/ui/input";
import * as React from "react";
import {Icons} from "@/components/ui/icons";
import {Button} from "@/components/ui/button";
import {Textarea} from "@/components/ui/textarea";
import FileUpload from "@/components/file-upload";

export default function CreateCourseForm() {
    const [isLoading, setIsLoading] = React.useState<boolean>(false)
    const form = useForm()

    const onFileUpload = () => {
    }

    const onSubmit = () => {

    }

    return (
        <div>
            <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                    <FormField
                        control={form.control}
                        name="name"
                        render={({field}) => (
                            <FormItem>
                                <FormLabel>Nome</FormLabel>
                                <FormControl>
                                    <Input type="text" placeholder="Nome do curso" {...field} />
                                </FormControl>
                                <FormMessage/>
                            </FormItem>
                        )}
                    />

                    <FormField
                        control={form.control}
                        name="description"
                        render={({field}) => (
                            <FormItem>
                                <FormLabel>Descrição</FormLabel>
                                <FormControl>
                                    <Textarea placeholder="Descrição do curso" {...field} />
                                </FormControl>
                                <FormMessage/>
                            </FormItem>
                        )}
                    />


                    <input type="hidden" name="imageUrl"/>

                    <FileUpload onFileUpload={onFileUpload}/>

                    <Button type="submit" disabled={isLoading}>
                        {isLoading && (
                            <Icons.spinner className="mr-2 h-4 w-4 animate-spin"/>
                        )}
                        Criar curso
                    </Button>
                </form>
            </Form>
        </div>
    )
}