import NextAuth from "next-auth";

declare module "next-auth" {
  interface Session {
    id: string;
    email: string;
    name: string;
    token: string;
  }
}
